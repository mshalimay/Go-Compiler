package parser

import (
	"fmt"
	"golite/types"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// Parser is an interface to expose to public desired methods/fields of parserWrapper
type Parser interface{
	Parse() 			*ast.Program
	GetProgramText() 	string
	GetErrors() 		[]*context.CompilerError
}


// parseWrapper is a wrapper around the antlr parser.
// It is a private struct, because contains methods and fields that should not be exposed.
// Parser interface is used to expose desired methods/fields of this struct.
type parserWrapper struct{
	*antlr.DefaultErrorListener 				// implements antlr.ErrorListener interface
	*BaseGoliteParserListener					// inherits from BaseGoliteParserListener
	antlrParser 	*GoliteParser
	errors 			[]*context.CompilerError
	lexer 			lexer.Lexer
	nodes 			map[string]interface{} 		// nodes to build the AST; holds structs from `golite/ast`
	UserTypes 		map[string]types.StructType
}

// instantiate a new parserWrapper
func NewParser(lexer lexer.Lexer) Parser{
	// create a pointer to a parserWrapper
	parser:= &parserWrapper{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{},
		nil, nil, lexer, make(map[string]interface{}), make(map[string]types.StructType)}
	
	// instantiate antler parser and override the default error listener
	antlrParser := NewGoliteParser(lexer.GetTokenStream())
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(parser)

	// fill parser fields and return
	parser.antlrParser = antlrParser
	return parser
}

//=============================================================================
// error listener methods
//==============================================================================
// Override the default error listener to add errors to parser
func (parser *parserWrapper) SyntaxError(recognizer antlr.Recognizer, 
	offendingSymbol interface{}, line, column int, msg string, 
	e antlr.RecognitionException) {
	
	parser.errors = append(parser.errors, &context.CompilerError{
		Line: line,
		Column: column, 
		Msg: msg, 
		Phase: context.PARSER,
	})

}

func (parser *parserWrapper) GetErrors() []*context.CompilerError{
	return parser.errors
}

//=============================================================================
// parseWrapper methods
//==============================================================================
// parse the lexed program and return an AST; if errors, return nil
func (parser *parserWrapper) Parse() *ast.Program {
	ctx := parser.antlrParser.Program()
	if context.HasErrors(parser.lexer.GetErrors()) || 
		context.HasErrors(parser.GetErrors()){
			return nil
		}
	
	// parser adheres to baseGoliteParserListener interface so can be passed to walker
	antlr.ParseTreeWalkerDefault.Walk(parser,ctx)
	_, _, key := GetTokenInfo(ctx)

	program := parser.nodes[key].(*ast.Program)

	program.UserTypes = parser.UserTypes

	return parser.nodes[key].(*ast.Program)
}


// return program text as seens by the antlr parser
func (parser *parserWrapper) GetProgramText() string{
	return parser.antlrParser.Program().GetText()
}

//==============================================================================
// Exit Listener methods
//==============================================================================
// below methods override the default exit methods in BaseGoliteParserListener for each node
// defined in the GoliteParser.g4 grammar file

// helper function to get line, column, and generate a key for a given context.
// Used to create unique keys in parser.nodes to build the AST
func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string){
	line, column:= ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()
	key := fmt.Sprintf("%d,%d", line, column)
	return line, column, key
}

func (parser parserWrapper) ExitProgram(ctx *ProgramContext){
	//production rules: 
	// program: types declarations functions

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get all elements in types, declarations, and functions
	typesCtxs := ctx.Types().AllTypeDeclaration()
	declsCtxs := ctx.Declarations().AllDeclaration()
	funcsCtxs := ctx.Functions().AllFunction()

	var types []*ast.TypeDecl
	var decls []*ast.Declaration
	var funcs []*ast.Function

	// build slice of `typeDeclaration` from types: typeDeclaration*;
	for _, typesCtx := range typesCtxs{
		_, _, typesKey := GetTokenInfo(typesCtx)
		types = append(types, parser.nodes[typesKey].(*ast.TypeDecl))
	}

	// build slice of `declaration` from declarations: declaration*;
	for _, declsCtx := range declsCtxs{
		_, _, declsKey := GetTokenInfo(declsCtx)
		decls = append(decls, parser.nodes[declsKey].(*ast.Declaration))
	}

	// build slice of `function` from functions: function*;
	for _, funcsCtx := range funcsCtxs{
		_, _, funcsKey := GetTokenInfo(funcsCtx)
		funcs = append(funcs, parser.nodes[funcsKey].(*ast.Function))
	}

	// create new `Program` token and pass up the tree
	parser.nodes[key] = ast.NewProgram(token.NewToken(line, col), types, decls, funcs)
}


func (parser parserWrapper) ExitTypeDeclaration(ctx *TypeDeclarationContext){
	// production rules:
	// typeDeclaration: TYPE ID STRUCT LBRACE fields RBRACE SEMICOLON;
	// fields: decl SEMICOLON fieldsPrime*;
	// fieldsPrime: decl SEMICOLON;
	// decl: ID type;

	// generate key for current node
	line, col, key := GetTokenInfo(ctx)

	// slices of IDs and types to build fields in typeDecl
	fields := make([]*ast.Decl, 0)

	// get the struct ID
	strucID := ctx.ID().GetText()

	// get id and type from `decl` in `fields` context
	l, c, _ := GetTokenInfo(ctx.Fields().Decl())
	id := ctx.Fields().Decl().ID().GetText()
	ty := ctx.Fields().Decl().Type_().GetText()
	fields = append(fields, ast.NewDecl(token.NewToken(l, c), id, ty))

	// get all ids and types `decl` inside `fieldsPrime*`, if any
	fpCtxs := ctx.Fields().AllFieldsPrime()
	for _, fpCtx := range fpCtxs{
		l, c, _ := GetTokenInfo(fpCtx)
		id := fpCtx.Decl().ID().GetText()
		ty := fpCtx.Decl().Type_().GetText()

		// fill fields map with id and type from `decl` in `fields` context
		fields = append(fields, ast.NewDecl(token.NewToken(l, c), id, ty))
	}

	// create typeDecl token and pass up the tree
	parser.nodes[key] = ast.NewTypeDecl(token.NewToken(line, col), strucID, fields)

	// create type and add to parser.Types (if not duplicate)
	if _, ok := parser.UserTypes[strucID]; !ok {
		parser.UserTypes[strucID] = &types.StrucType{Name: strucID}
	}
}


func (parser parserWrapper) ExitDeclaration(ctx *DeclarationContext){
	// production rules:
	// declaration: VAR ids type SEMICOLON;
	// ids: ID idsPrime* ;
	// idsPrime: COMMA ID;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// slice to of IDs to build declaration token
	ids := []string{}

	// get `ids` context from `declaration` context
	idsCtx := ctx.Ids()
	
	// append first ID in `ids` to slice
	ids = append(ids, idsCtx.ID().GetText())
	
	// append remaining IDs in `idsPrime*` to slice, if any
	idsPrimeCtxs := idsCtx.AllIdsPrime()
	for _, idsPrimeCtx := range idsPrimeCtxs{
		ids = append(ids, idsPrimeCtx.ID().GetText())
	}

	// get string from `type` token in `declaration` context
	tyStr := ctx.Type_().GetText()

	// create new `Declaration` token and pass up the tree
	parser.nodes[key] = ast.NewDeclaration(token.NewToken(line, col), ids, tyStr)
}


func (parser *parserWrapper) ExitFunction(ctx *FunctionContext){
	// production rules:
	// function: FUNC ID parameters (returnType)? LBRACE declarations statements RBRACE;
	// parameters: LPAREN (decl parametersPrime*)? RPAREN; parametersPrime: COMMA decl;
	// declarations: declaration*;  declaration: ID type;
	// statements: statement*; statement: block | assignment | print | delete | conditional | loop | ret | invocation;
	// returnType: type;
	
	
	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get `parameters` token from previous nodes
	_, _, paramsKey := GetTokenInfo(ctx.Parameters())
	params := parser.nodes[paramsKey].(*ast.Parameters)

	// build slice of `declaration` from declarations: declaration*;
	var decls []*ast.Declaration
	declCtxs := ctx.Declarations().AllDeclaration()
	for _, declCtx := range declCtxs{
		_, _, declKey := GetTokenInfo(declCtx)
		decls = append(decls, parser.nodes[declKey].(*ast.Declaration))
	}

	// build slice of `statement` from statements: statement*;
	var stmts []ast.Statement
	stmtCtxs := ctx.Statements().AllStatement() 
	for _, stmtCtx := range stmtCtxs{
		_, _, stmtKey := GetTokenInfo(stmtCtx)
		stmts = append(stmts, parser.nodes[stmtKey].(ast.Statement))
	}
	
	// get `returnType` inside `function` context, if any
	var returnType string
	if ctx.ReturnType() != nil{
		returnType = ctx.ReturnType().GetText()
	}

	// create new `Function` token and pass up the tree
	parser.nodes[key] = ast.NewFunction(token.NewToken(line, col), ctx.ID().GetText(), returnType, params, decls, stmts)

}

func (parser *parserWrapper) ExitParameters(ctx *ParametersContext){
	// production rules:
	// parameters: LPAREN (decl parametersPrime*)? RPAREN;    
	// parametersPrime: COMMA decl;
	// decl: ID type;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// ids and types slices used to build `decl` components eg: `a int`, `b bool`
	// obs: having separate slices for ID and Types help in syntatic analysis
 	params := make([]*ast.Decl, 0)
	
	// check if (decl (COMMA decl)*)? is present and parse accordingly
	// obs: @memo `decl` not the same as `declaration`
	var id, ty string

	if declCtx:= ctx.Decl(); declCtx != nil{	
		l,c, _ := GetTokenInfo(declCtx)
		id = declCtx.ID().GetText()
		ty = declCtx.Type_().GetText()
		params = append(params, ast.NewDecl(token.NewToken(l,c), id, ty))
	
		// get `decl` contexts from `parametersPrime*` and fill ids and types slices
		paramCtxs := ctx.AllParametersPrime()
		for _, paramCtx := range paramCtxs{
			l, c, _ := GetTokenInfo(paramCtx)
			id = paramCtx.Decl().ID().GetText()
			ty = paramCtx.Decl().Type_().GetText()
			params = append(params, ast.NewDecl(token.NewToken(l,c), id, ty))
		}
	}
	// create new `Parameters` token and pass up the tree
	parser.nodes[key] = ast.NewParameters(token.NewToken(line, col), params)
}


func (parser *parserWrapper) ExitStatement(ctx *StatementContext){
	// production rules:
	// statement: block | assignment | print | delete | conditional | loop | ret | invocation;
	
	// get key for current node
	_, _, key := GetTokenInfo(ctx)

	// check which statement is present and pass up the tree accordingly
	if blockCtx := ctx.Block(); blockCtx != nil{
		// case: block
		_, _, blockKey := GetTokenInfo(blockCtx)
		parser.nodes[key] = parser.nodes[blockKey].(ast.Statement)
	
	} else if assignmentCtx := ctx.Assignment(); assignmentCtx != nil{
		// case: assignment
		_, _, assignmentKey := GetTokenInfo(assignmentCtx)
		parser.nodes[key] = parser.nodes[assignmentKey].(ast.Statement)
	
	} else if printCtx := ctx.Print_(); printCtx != nil{
		// case: print
		_, _, printKey := GetTokenInfo(printCtx)
		parser.nodes[key] = parser.nodes[printKey].(ast.Statement)
	
	} else if deleteCtx := ctx.Delete_(); deleteCtx != nil{
		// case: delete
		_, _, deleteKey := GetTokenInfo(deleteCtx)
		parser.nodes[key] = parser.nodes[deleteKey].(ast.Statement)
	
	} else if conditionalCtx := ctx.Conditional(); conditionalCtx != nil{
		// case: conditional
		_, _, conditionalKey := GetTokenInfo(conditionalCtx)
		parser.nodes[key] = parser.nodes[conditionalKey].(ast.Statement)
	
	} else if loopCtx := ctx.Loop(); loopCtx != nil{
		// case: loop
		_, _, loopKey := GetTokenInfo(loopCtx)
		parser.nodes[key] = parser.nodes[loopKey].(ast.Statement)
	
	} else if retCtx := ctx.Ret(); retCtx != nil{
		// case: ret
		_, _, retKey := GetTokenInfo(retCtx)
		parser.nodes[key] = parser.nodes[retKey].(ast.Statement)
	
	} else if invocationCtx := ctx.Invocation(); invocationCtx != nil{
		// case: invocation
		_, _, invocationKey := GetTokenInfo(invocationCtx)
		parser.nodes[key] = parser.nodes[invocationKey].(ast.Statement)
	}
}

func (parser *parserWrapper) ExitBlock(ctx *BlockContext){
	// production rules:
	// block: LBRACE statement* RBRACE;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// slice of statements inside block
	var statements []ast.Statement

	// create the slice of statements from `statement*`
	// get all elements in statement*
	statementsCtx := ctx.Statements().AllStatement()

	// get `statement` from `statement*` and append to statements slice, if any
	for _, statementCtx := range statementsCtx{
		_, _, statementKey := GetTokenInfo(statementCtx)
		statements = append(statements, parser.nodes[statementKey].(ast.Statement))
	}

	// create new `Block` token and pass up the tree
	parser.nodes[key] = ast.NewBlock(token.NewToken(line, col), statements)
}



func (parser *parserWrapper) ExitDelete(ctx *DeleteContext){
	// production rules:
	// delete: DELETE expression SEMICOLON;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get `expression` inside `delete` context
	_, _, exprKey := GetTokenInfo(ctx.Expression())

	// create new `Delete` token and pass up the tree
	parser.nodes[key] = ast.NewDelete(token.NewToken(line, col), parser.nodes[exprKey].(ast.Expression))

}

func (parser *parserWrapper) ExitAssignment(ctx *AssignmentContext){
	// production rules:
	// assignment: lvalue EQUALS (expression | SCAN) SEMICOLON;
	// lvalue: ID lvaluePrime*; lvaluePrime: DOT ID;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get `lvalue` ` inside `assignment` context
	_, _, lvalueKey := GetTokenInfo(ctx.Lvalue())
	lvalue := parser.nodes[lvalueKey].(ast.Expression)

	// get `expression` or `SCAN` and pass new `Assignment` token to the tree
	if exprCtx:=ctx.Expression(); exprCtx != nil{
		// case lvalue EQUALS expression SEMICOLON
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = ast.NewAssignment(token.NewToken(line, col), lvalue, expr, nil)
	
	} else if ctx.SCAN() != nil {
		// case lvalue EQUALS SCAN SEMICOLON
		//NOTE: using same key for scan and assignment tokens bcs dont have antlr.context for SCAN
		scan := ast.NewScan(token.NewToken(line, col))
		parser.nodes[key] = ast.NewAssignment(token.NewToken(line, col), lvalue, nil, scan)
	}
}

func (parser *parserWrapper) ExitPrint(ctx *PrintContext){
	// production rules:
	// print: PRINTF LPAREN STRING printPrime* RPAREN SEMICOLON; 
	// printPrime: COMMA expression;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get string inside `print` context
	str := ctx.STRING().GetText()

	// build slice of expressions from `printPrime*`, if any
	var expressions []ast.Expression

	// get all elements in printPrime*
	printContexts := ctx.AllPrintPrime()
	if len(printContexts) > 0 {
		for _, printContext := range printContexts{
			// get `expression` inside printPrime and append to expressions slice
			_, _, exprKey := GetTokenInfo(printContext.Expression())
			expressions = append(expressions, parser.nodes[exprKey].(ast.Expression))
		}
	}

	// create new `Print` token and pass to the tree
	parser.nodes[key] = ast.NewPrint(token.NewToken(line, col), str, expressions)
}
		


func (parser *parserWrapper) ExitConditional(ctx *ConditionalContext){
	// porduction rules:
	// conditional: IF LPAREN expression RPAREN block (ELSE block)?;
	
	// NOTE: change Block(0) and Block(1) if to allow more blocks (like `else if` statements)
	
	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get expression and block inside `conditional` context
	_, _, exprKey := GetTokenInfo(ctx.Expression())
	_, _, blockKey := GetTokenInfo(ctx.Block(0))

	// get `expression` and `block` from `conditional` context
	expression := parser.nodes[exprKey].(ast.Expression)
	block := parser.nodes[blockKey].(ast.Statement)

	// get `elseBlock` from `conditional` context, if any
	var elseBlock ast.Statement = nil
	if ctx.ELSE() != nil{
		_, _, elseBlockKey := GetTokenInfo(ctx.Block(1))
		elseBlock = parser.nodes[elseBlockKey].(ast.Statement)
	}

	// create new `Conditional` token and pass up the tree
	parser.nodes[key] = ast.NewConditional(token.NewToken(line, col), expression, block, elseBlock)
}

func (parser *parserWrapper) ExitLoop(ctx *LoopContext){
	// production rules:
	// loop: FOR LPAREN expression RPAREN block
	
	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get expression inside `loop` context
	_, _, exprKey := GetTokenInfo(ctx.Expression())
	_, _, blockKey := GetTokenInfo(ctx.Block())

	// get `expression` and `block` from `loop` context
	expression := parser.nodes[exprKey].(ast.Expression)
	block := parser.nodes[blockKey].(ast.Statement)

	// create `Loop` token and pass up the tree
	parser.nodes[key] = ast.NewLoop(token.NewToken(line, col), expression, block)
}


func (parser *parserWrapper) ExitRet(ctx *RetContext){
	// production rules:
	// ret: RETURN expression? SEMICOLON;
	
	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get expression inside `ret` context, if any
	var expr ast.Expression
	if ctx.Expression() != nil{
		_, _, exprKey := GetTokenInfo(ctx.Expression())
		expr = parser.nodes[exprKey].(ast.Expression)
	}

	// create `Ret` token and pass up the tree
	parser.nodes[key] = ast.NewRet(token.NewToken(line, col), expr, nil)
}


func (parser *parserWrapper) ExitInvocation(ctx *InvocationContext){
	// production rules:
	// invocation: ID arguments;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// get `arguments` token inside `invocation` context
	_, _, argsKey := GetTokenInfo(ctx.Arguments())
	args := parser.nodes[argsKey].(*ast.Arguments)

	// create new `Invocation` token and pass up the tree
	parser.nodes[key] = ast.NewInvocation(token.NewToken(line, col), ctx.ID().GetText(), args)
}


func (parser *parserWrapper) ExitArguments(ctx *ArgumentsContext){
	// production rules:
	// arguments: LPAREN (expression argumentsPrime*)? RPAREN; 
	// argumentsPrime: COMMA expression;

	// get key for current node
	line, col, key := GetTokenInfo(ctx)

	// slice of expression terms inside parentheses. e.g.: []{a,b,c} for (a, b, c)
	var expressions []ast.Expression

	// create the slice of expressions from `arguments` and `argumentsPrime*`
	if expression := ctx.Expression(); ctx.Expression() != nil{
		
		// append `expression` in `arguments` to slice
		_, _, exprKey := GetTokenInfo(expression)
		expressions = append(expressions, parser.nodes[exprKey].(ast.Expression))
		
		// get all elements in argumentsPrime*
		argsContexts := ctx.AllArgumentsPrime()

		// get remaining expressions from `argumentsPrime*` and append to expressions slice, if any
		for _, argsContext := range argsContexts{
			// get `expression` inside argumentsPrime and append to expressions slice
			_, _, exprKey := GetTokenInfo(argsContext.Expression())
			expressions = append(expressions, parser.nodes[exprKey].(ast.Expression))
		}
	}
	// create new `Arguments` token and pass up the tree
	parser.nodes[key] = ast.NewArguments(token.NewToken(line, col), expressions) 
}


func (parser *parserWrapper) ExitLvalue(ctx *LvalueContext){
	// production rules:
	// lvalue: ID (DOT ID)*

	// get key for current context
	line, col, key := GetTokenInfo(ctx)

	// get all elements in lvaluePrime*
	lvalueContexts := ctx.AllLvaluePrime()

	// if there are no elements in lvaluePrime*, lvalue is just an ID; else, composite ID
	if len(lvalueContexts) == 0 {
		// case: ID
		// create new `LValue` token holding a `ast.Variable` and pass up the tree
		newVariable := ast.NewVariable(ctx.ID().GetText(), token.NewToken(line, col), nil)
		parser.nodes[key] = ast.NewLValue(token.NewToken(line, col), newVariable, nil, nil)
	} else {
		// case: ID (DOT ID)*
		ids := []string{}
		
		// the first ID comes from `lvalue` context
		ids = append(ids, ctx.ID().GetText())
		
		// build the ".ID.ID" chain with elements from `lvaluePrime*`
		for _, lvalueContext := range lvalueContexts{
			ids = append(ids, lvalueContext.ID().GetText())
		}
		// create new `LValue` token holding a `ast.Field` and pass up the tree
		newField := ast.NewField(token.NewToken(line, col), ids, nil)
		parser.nodes[key] = ast.NewLValue(token.NewToken(line, col), nil, newField, nil)
	}
}


func (parser *parserWrapper) ExitExpression(ctx *ExpressionContext){
	// production rules:
	// expression: boolTerm expressionPrime*;          
	// expressionPrime: OR boolTerm;
	
	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.BoolTerm())
	
	// get all elements in termPrime*
	expTermContexts := ctx.AllExpressionPrime()

	// if no elements in `equalTermPrime*`, `equalTerm` is just result from `relationTerm` exit method;
	// else, create a nested binOp using `relationTerm` and elements in `equalTermPrime*`
	
	if len(expTermContexts) == 0{ // case: equalTerm
		// pass expression up the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: boolTerm: equalTerm (AND equalTerm)*
		
		// get `equalTerm` from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from `equalTerm` and elements in `boolTermTermPrime*`
		for _, expTermContext := range expTermContexts{
			
			// get `equalTerm` inside `boolTermTermPrime`
			l, c, termKey := GetTokenInfo(expTermContext.BoolTerm())
			rhs := parser.nodes[termKey].(ast.Expression)

			// get >, <, >=, <= operator inside `boolTermTermPrime`
			operator := ast.StrToOperator(expTermContext.OR().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// create new `Expression` token and pass up the tree
		parser.nodes[key] = lhs
	}
}


func (parser *parserWrapper) ExitBoolTerm(ctx *BoolTermContext){
	// production rules:
	// boolTerm: equalTerm booltermPrime*;             
	// booltermPrime: AND equalTerm;

	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.EqualTerm())
	
	// get all elements in termPrime*
	btTermContexts := ctx.AllBooltermPrime()

	// if no elements in `equalTermPrime*`, `equalTerm` is just result from `relationTerm` exit method;
	// else, create a nested binOp using `relationTerm` and elements in `equalTermPrime*`
	
	if len(btTermContexts) == 0{ // case: equalTerm
		// pass expression up the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: boolTerm: equalTerm (AND equalTerm)*
		
		// get `equalTerm` from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from `equalTerm` and elements in `boolTermTermPrime*`
		for _, btTermContext := range btTermContexts{
			
			// get `equalTerm` inside `boolTermTermPrime`
			l, c, termKey := GetTokenInfo(btTermContext.EqualTerm())
			rhs := parser.nodes[termKey].(ast.Expression)

			// get >, <, >=, <= operator inside `boolTermTermPrime`
			operator := ast.StrToOperator(btTermContext.AND().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// pass final expression up the tree
		parser.nodes[key] = lhs
	}
}

func (parser *parserWrapper) ExitEqualTerm(ctx *EqualTermContext){
	// production rules:
	// equalTerm: relationTerm equalTermPrime*;             
	// equalTermPrime: op=(EQ | NEQ) relt=relationTerm;
	
	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.RelationTerm())
	
	// get all elements in termPrime*
	eqTermContexts := ctx.AllEqualTermPrime()

	// if no elements in `equalTermPrime*`, `equalTerm` is just result from `relationTerm` exit method;
	// else, create a nested binOp using `relationTerm` and elements in `equalTermPrime*`
	
	if len(eqTermContexts) == 0{ // case: relationTerm
		// pass term up the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: relationTerm ((EQ | NEQ) relationTerm)*
		
		// get `simpleTerm` from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from `simpleTerm` and elements in `relationTermPrime*`
		for _, eqTermContext := range eqTermContexts{
			
			// get `relationTerm` inside `eqTermPrime`
			l, c, termKey := GetTokenInfo(eqTermContext.RelationTerm())
			rhs := parser.nodes[termKey].(ast.Expression)

			// get >, <, >=, <= operator inside `relationTermPrime`
			operator := ast.StrToOperator(eqTermContext.GetOp().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// pass final expression up the tree
		parser.nodes[key] = lhs
	}
}

func (parser *parserWrapper) ExitRelationTerm(ctx *RelationTermContext){
	// production rules:
	// relationTerm: simpleTerm relationTermPrime*;    
	// relationTermPrime: op=(GT | LT | LTE | GTE) smpt=simpleTerm;

	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.SimpleTerm())
	
	// get all elements in termPrime*
	rTermContexts := ctx.AllRelationTermPrime()

	// if no elements in `simpleTermPrime*`, `simpleTerm` is just result from `term` exit method;
	// else, create a nested binOp using `term` and elements in simpleTermPrime*
	
	if len(rTermContexts) == 0 { // case: simpleTerm
		// pass term up the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: simpleTerm ((GT | LT | LTE | GTE) simpleTerm)*
		
		// get `simpleTerm` from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from `simpleTerm` and elements in `relationTermPrime*`
		for _, rTermContext := range rTermContexts{
			
			// get `simpleTerm` inside `relationTermPrime`
			l, c, termKey := GetTokenInfo(rTermContext.SimpleTerm())
			rhs := parser.nodes[termKey].(ast.Expression)

			// get >, <, >=, <= operator inside `relationTermPrime`
			operator := ast.StrToOperator(rTermContext.GetOp().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// pass final expression up the tree
		parser.nodes[key] = lhs
	}
}

func (parser *parserWrapper) ExitSimpleTerm(ctx *SimpleTermContext){
	// production rules:
	// simpleTerm: term simpleTermPrime*;
	// simpleTermPrime: op=(PLUS | MINUS) term;
	
	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.Term())
	
	// get all elements in termPrime*
	sTermContexts := ctx.AllSimpleTermPrime()

	// if no elements in `simpleTermPrime*`, `simpleTerm` is just result from `term` exit method;
	// else, create a nested binOp using `term` and elements in simpleTermPrime*
	
	if len(sTermContexts) == 0{ // case: term
		// pass term up the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: simpleTerm: term ((PLUS | MINUS) term)*;
		
		// get `term` from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from `term` and elements in `simpleTermPrime*`
		for _, sTermContext := range sTermContexts{
			
			// get `term` inside `simpleTermPrime`
			l, c, termKey := GetTokenInfo(sTermContext.Term())
			rhs := parser.nodes[termKey].(ast.Expression)

			// get +, - operator inside `simpleTermPrime`
			operator := ast.StrToOperator(sTermContext.GetOp().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// pass final expression up the tree
		parser.nodes[key] = lhs
	}
}

func (parser *parserWrapper) ExitTerm(ctx *TermContext){
	// production rules:
	// term: unaryTerm ((ASTERISK | FSLASH) unaryTerm)*
	// termPrime : op=(ASTERISK | FSLASH) unat=unaryTerm;

	// get key for current context; dont need line column because not creating a new token in this Exit method
	_, _, key := GetTokenInfo(ctx)
	
	// get the 'key' for the child node token
	_, _, exprKey := GetTokenInfo(ctx.UnaryTerm())
	
	// get all elements in termPrime*
	termContexts := ctx.AllTermPrime()

	
	// if there are no elements in termPrime*, term is just a unaryTerm; else, nested binOp
	if len(termContexts) == 0{ // case: unaryTerm
		// pass unaryTerm upp the tree
		expr := parser.nodes[exprKey].(ast.Expression)
		parser.nodes[key] = expr
	
	} else { // case: unaryTerm ((ASTERISK | FSLASH) unaryTerm)*
		// get unaryTerm from down the tree
		lhs := parser.nodes[exprKey].(ast.Expression)
		
		// compose binOp expresion from unaryTerm and elements in termPrime*
		for _, termContext := range termContexts{
			// get unaryTerm inside termPrime
			l, c, termKey := GetTokenInfo(termContext.GetUnat())

			// get *, / operator inside termPrime
			rhs := parser.nodes[termKey].(ast.Expression)
			operator := ast.StrToOperator(termContext.GetOp().GetText())

			// compose binOp expression
			lhs = ast.NewBinOp(operator, lhs, rhs, token.NewToken(l,c), nil)
		}
		// pass final expression up the tree
		parser.nodes[key] = lhs
	}
}

func (parser *parserWrapper) ExitUnaryTerm(ctx *UnaryTermContext) {
	// production rules:
	// unaryTerm: (NOT | MINUS)? selectorTerm;
	
	// compute key for UnaryTermContext
	line, column, key := GetTokenInfo(ctx)


	// get expression = selectorTerm from `unaryTerm` context
	_, _, factorKey := GetTokenInfo(ctx.SelectorTerm())
	factor := parser.nodes[factorKey].(ast.Expression)
	

	// get operator from `unaryTerm` context, if any
	var operator ast.Operator = -1 // no operator
	if ctx.NOT() != nil{
		// case: !selectorTerm
		operator = ast.StrToOperator(ctx.NOT().GetText())		
	} else if ctx.MINUS() != nil{
		// case: -selectorTerm
		operator = ast.StrToOperator(ctx.MINUS().GetText())
	}

	// create new `UnaryTerm` token and pass up the tree
	parser.nodes[key] = ast.NewUnaryTerm(token.NewToken(line, column), operator, factor, nil)
}

func (parser *parserWrapper) ExitSelectorTerm(ctx *SelectorTermContext){
	// production rules:
	// selectorTerm: factor selectorTermPrime*;  
	// selectorTermPrime: DOT ID;

	_, _, key := GetTokenInfo(ctx)

	// get `factor` in the `selectorTerm` context
	line, column, factorKey := GetTokenInfo(ctx.Factor())
	factor := parser.nodes[factorKey].(ast.Expression)

	// if factor is not FactorPrime2 or is ID(arguments), pass up the factor
	if f, ok := factor.(*ast.FactorPrime2); !ok || ok && f.Arguments != nil{
		parser.nodes[key] = factor
		return
	} else {
		// If not of the cases above, factor is a variable (`x`) or a field (`x.y.z`)
	
		// get the IDs in `selectorTermPrime*`
		selTermPrimeCtxs := ctx.AllSelectorTermPrime()
		
		// if no IDs, factor is a variable
		if len(selTermPrimeCtxs) == 0{
			// factor is a variable
			// REVIEW column and line here should be from the selector ctx or factor ctx?
			parser.nodes[key] = ast.NewVariable(f.Identifier, token.NewToken(line, column), nil)
			return
		}

		// otherwise, factor is a field. Build it from the IDs in `selectorTermPrime*`
		var ids[] string
		ids = append(ids, f.Identifier)
		for _, selTermPrimeCtx := range selTermPrimeCtxs{
			ids = append(ids, selTermPrimeCtx.ID().GetText())
		}
		parser.nodes[key] = ast.NewField(token.NewToken(line, column), ids, nil)
	}		
}


func (parser *parserWrapper) ExitFactor(ctx *FactorContext) {
	// production rules:
	// factor: factorPrime1 | factorPrime2 | NUMBER | NEW ID | TRUE | FALSE | NIL;
	// factorPrime1: LPAREN expression RPAREN;
	// factorPrime2: ID (arguments)?;

	// get line, column and generate node key for current context
	line, column, key := GetTokenInfo(ctx)
	
	// check which factor is present and pass up the tree accordingly

	// NUMBER
	if intFactor := ctx.NUMBER(); intFactor != nil{
		intValue, _ := strconv.Atoi(intFactor.GetText())
		parser.nodes[key] = ast.NewIntLiteral(int64(intValue), token.NewToken(line, column))
	
	// TRUE
	}else if trueFactor := ctx.TRUE(); trueFactor != nil{
		parser.nodes[key] = ast.NewBoolLiteral(true, token.NewToken(line, column))
	
	// FALSE
	}else if falseFactor := ctx.FALSE(); falseFactor != nil{
		parser.nodes[key] = ast.NewBoolLiteral(false, token.NewToken(line, column))
	
	// NIL
	}else if nilFactor := ctx.NIL(); nilFactor != nil{
		parser.nodes[key] = ast.NewNilLiteral(token.NewToken(line, column))
	
	// NEW ID
	}else if newFactor := ctx.NEW(); newFactor != nil{
		if idFactor := ctx.ID(); idFactor != nil{
			// NOTE: uncomment below if want to pass a variable to New token, instead of just string
			//newVar := ast.NewVariable(idFactor.GetText(), token.NewToken(line, column))
			parser.nodes[key] = ast.NewNew(token.NewToken(line, column), idFactor.GetText(), nil)
		} else {
			panic("'new' not followed by 'ID'")
		}

	// factorPrime1: LPAREN expression RPAREN;
	} else if factPrime := ctx.FactorPrime1(); factPrime != nil{
		
		// get `expression` inside `factorPrime1` context
		_, _, exprKey := GetTokenInfo(factPrime.Expression())
		expr := parser.nodes[exprKey].(ast.Expression)
		
		// pass factor up the tree
		parser.nodes[key] = ast.NewFactorPrime1(token.NewToken(line, column), expr, nil)
			
	// factorPrime2: ID (arguments)?;
	} else if factPrime := ctx.FactorPrime2(); factPrime != nil{
		// case: ID arguments
		if argsCtx := factPrime.Arguments(); argsCtx != nil{
			
			// get `arguments` token inside `factorPrime2` context
			_, _, argsKey := GetTokenInfo(argsCtx)
			args := parser.nodes[argsKey].(*ast.Arguments)

			// create `FactorPrime2` with arguments and pass up the tree
			parser.nodes[key] = ast.NewFactorPrime2(token.NewToken(line, column), factPrime.ID().GetText(), args, nil)
		// case: ID
		}else{
			// create `FactorPrime2` without arguments and pass up the tree
			// in `SelectorTerm` we decide if it is a `variable` or a `field`
			parser.nodes[key] = ast.NewFactorPrime2(token.NewToken(line, column), factPrime.ID().GetText(), nil, nil)
		}

	// unrecognized factor
	} else{
		panic("unknown factor")
	}
}
		

	