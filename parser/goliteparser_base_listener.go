// Code generated from GoliteParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr4-go/antlr/v4"

// BaseGoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type BaseGoliteParserListener struct{}

var _ GoliteParserListener = &BaseGoliteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoliteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoliteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoliteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoliteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoliteParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoliteParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseGoliteParserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseGoliteParserListener) ExitTypes(ctx *TypesContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseGoliteParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseGoliteParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseGoliteParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseGoliteParserListener) ExitFields(ctx *FieldsContext) {}

// EnterFieldsPrime is called when production fieldsPrime is entered.
func (s *BaseGoliteParserListener) EnterFieldsPrime(ctx *FieldsPrimeContext) {}

// ExitFieldsPrime is called when production fieldsPrime is exited.
func (s *BaseGoliteParserListener) ExitFieldsPrime(ctx *FieldsPrimeContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseGoliteParserListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseGoliteParserListener) ExitDecl(ctx *DeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoliteParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoliteParserListener) ExitType(ctx *TypeContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseGoliteParserListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseGoliteParserListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGoliteParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGoliteParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterIds is called when production ids is entered.
func (s *BaseGoliteParserListener) EnterIds(ctx *IdsContext) {}

// ExitIds is called when production ids is exited.
func (s *BaseGoliteParserListener) ExitIds(ctx *IdsContext) {}

// EnterIdsPrime is called when production idsPrime is entered.
func (s *BaseGoliteParserListener) EnterIdsPrime(ctx *IdsPrimeContext) {}

// ExitIdsPrime is called when production idsPrime is exited.
func (s *BaseGoliteParserListener) ExitIdsPrime(ctx *IdsPrimeContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseGoliteParserListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseGoliteParserListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseGoliteParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseGoliteParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGoliteParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGoliteParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParametersPrime is called when production parametersPrime is entered.
func (s *BaseGoliteParserListener) EnterParametersPrime(ctx *ParametersPrimeContext) {}

// ExitParametersPrime is called when production parametersPrime is exited.
func (s *BaseGoliteParserListener) ExitParametersPrime(ctx *ParametersPrimeContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseGoliteParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseGoliteParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseGoliteParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseGoliteParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoliteParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoliteParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoliteParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoliteParserListener) ExitBlock(ctx *BlockContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseGoliteParserListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseGoliteParserListener) ExitDelete(ctx *DeleteContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoliteParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoliteParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseGoliteParserListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseGoliteParserListener) ExitPrint(ctx *PrintContext) {}

// EnterPrintPrime is called when production printPrime is entered.
func (s *BaseGoliteParserListener) EnterPrintPrime(ctx *PrintPrimeContext) {}

// ExitPrintPrime is called when production printPrime is exited.
func (s *BaseGoliteParserListener) ExitPrintPrime(ctx *PrintPrimeContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseGoliteParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseGoliteParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseGoliteParserListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseGoliteParserListener) ExitLoop(ctx *LoopContext) {}

// EnterRet is called when production ret is entered.
func (s *BaseGoliteParserListener) EnterRet(ctx *RetContext) {}

// ExitRet is called when production ret is exited.
func (s *BaseGoliteParserListener) ExitRet(ctx *RetContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseGoliteParserListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseGoliteParserListener) ExitInvocation(ctx *InvocationContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoliteParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoliteParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentsPrime is called when production argumentsPrime is entered.
func (s *BaseGoliteParserListener) EnterArgumentsPrime(ctx *ArgumentsPrimeContext) {}

// ExitArgumentsPrime is called when production argumentsPrime is exited.
func (s *BaseGoliteParserListener) ExitArgumentsPrime(ctx *ArgumentsPrimeContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseGoliteParserListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseGoliteParserListener) ExitLvalue(ctx *LvalueContext) {}

// EnterLvaluePrime is called when production lvaluePrime is entered.
func (s *BaseGoliteParserListener) EnterLvaluePrime(ctx *LvaluePrimeContext) {}

// ExitLvaluePrime is called when production lvaluePrime is exited.
func (s *BaseGoliteParserListener) ExitLvaluePrime(ctx *LvaluePrimeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoliteParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoliteParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionPrime is called when production expressionPrime is entered.
func (s *BaseGoliteParserListener) EnterExpressionPrime(ctx *ExpressionPrimeContext) {}

// ExitExpressionPrime is called when production expressionPrime is exited.
func (s *BaseGoliteParserListener) ExitExpressionPrime(ctx *ExpressionPrimeContext) {}

// EnterBoolTerm is called when production boolTerm is entered.
func (s *BaseGoliteParserListener) EnterBoolTerm(ctx *BoolTermContext) {}

// ExitBoolTerm is called when production boolTerm is exited.
func (s *BaseGoliteParserListener) ExitBoolTerm(ctx *BoolTermContext) {}

// EnterBooltermPrime is called when production booltermPrime is entered.
func (s *BaseGoliteParserListener) EnterBooltermPrime(ctx *BooltermPrimeContext) {}

// ExitBooltermPrime is called when production booltermPrime is exited.
func (s *BaseGoliteParserListener) ExitBooltermPrime(ctx *BooltermPrimeContext) {}

// EnterEqualTerm is called when production equalTerm is entered.
func (s *BaseGoliteParserListener) EnterEqualTerm(ctx *EqualTermContext) {}

// ExitEqualTerm is called when production equalTerm is exited.
func (s *BaseGoliteParserListener) ExitEqualTerm(ctx *EqualTermContext) {}

// EnterEqualTermPrime is called when production equalTermPrime is entered.
func (s *BaseGoliteParserListener) EnterEqualTermPrime(ctx *EqualTermPrimeContext) {}

// ExitEqualTermPrime is called when production equalTermPrime is exited.
func (s *BaseGoliteParserListener) ExitEqualTermPrime(ctx *EqualTermPrimeContext) {}

// EnterRelationTerm is called when production relationTerm is entered.
func (s *BaseGoliteParserListener) EnterRelationTerm(ctx *RelationTermContext) {}

// ExitRelationTerm is called when production relationTerm is exited.
func (s *BaseGoliteParserListener) ExitRelationTerm(ctx *RelationTermContext) {}

// EnterRelationTermPrime is called when production relationTermPrime is entered.
func (s *BaseGoliteParserListener) EnterRelationTermPrime(ctx *RelationTermPrimeContext) {}

// ExitRelationTermPrime is called when production relationTermPrime is exited.
func (s *BaseGoliteParserListener) ExitRelationTermPrime(ctx *RelationTermPrimeContext) {}

// EnterSimpleTerm is called when production simpleTerm is entered.
func (s *BaseGoliteParserListener) EnterSimpleTerm(ctx *SimpleTermContext) {}

// ExitSimpleTerm is called when production simpleTerm is exited.
func (s *BaseGoliteParserListener) ExitSimpleTerm(ctx *SimpleTermContext) {}

// EnterSimpleTermPrime is called when production simpleTermPrime is entered.
func (s *BaseGoliteParserListener) EnterSimpleTermPrime(ctx *SimpleTermPrimeContext) {}

// ExitSimpleTermPrime is called when production simpleTermPrime is exited.
func (s *BaseGoliteParserListener) ExitSimpleTermPrime(ctx *SimpleTermPrimeContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGoliteParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGoliteParserListener) ExitTerm(ctx *TermContext) {}

// EnterTermPrime is called when production termPrime is entered.
func (s *BaseGoliteParserListener) EnterTermPrime(ctx *TermPrimeContext) {}

// ExitTermPrime is called when production termPrime is exited.
func (s *BaseGoliteParserListener) ExitTermPrime(ctx *TermPrimeContext) {}

// EnterUnaryTerm is called when production unaryTerm is entered.
func (s *BaseGoliteParserListener) EnterUnaryTerm(ctx *UnaryTermContext) {}

// ExitUnaryTerm is called when production unaryTerm is exited.
func (s *BaseGoliteParserListener) ExitUnaryTerm(ctx *UnaryTermContext) {}

// EnterSelectorTerm is called when production selectorTerm is entered.
func (s *BaseGoliteParserListener) EnterSelectorTerm(ctx *SelectorTermContext) {}

// ExitSelectorTerm is called when production selectorTerm is exited.
func (s *BaseGoliteParserListener) ExitSelectorTerm(ctx *SelectorTermContext) {}

// EnterSelectorTermPrime is called when production selectorTermPrime is entered.
func (s *BaseGoliteParserListener) EnterSelectorTermPrime(ctx *SelectorTermPrimeContext) {}

// ExitSelectorTermPrime is called when production selectorTermPrime is exited.
func (s *BaseGoliteParserListener) ExitSelectorTermPrime(ctx *SelectorTermPrimeContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseGoliteParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseGoliteParserListener) ExitFactor(ctx *FactorContext) {}

// EnterFactorPrime1 is called when production factorPrime1 is entered.
func (s *BaseGoliteParserListener) EnterFactorPrime1(ctx *FactorPrime1Context) {}

// ExitFactorPrime1 is called when production factorPrime1 is exited.
func (s *BaseGoliteParserListener) ExitFactorPrime1(ctx *FactorPrime1Context) {}

// EnterFactorPrime2 is called when production factorPrime2 is entered.
func (s *BaseGoliteParserListener) EnterFactorPrime2(ctx *FactorPrime2Context) {}

// ExitFactorPrime2 is called when production factorPrime2 is exited.
func (s *BaseGoliteParserListener) ExitFactorPrime2(ctx *FactorPrime2Context) {}
