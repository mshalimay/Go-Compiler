// Code generated from GoliteParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr4-go/antlr/v4"

// GoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type GoliteParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterFieldsPrime is called when entering the fieldsPrime production.
	EnterFieldsPrime(c *FieldsPrimeContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterIds is called when entering the ids production.
	EnterIds(c *IdsContext)

	// EnterIdsPrime is called when entering the idsPrime production.
	EnterIdsPrime(c *IdsPrimeContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParametersPrime is called when entering the parametersPrime production.
	EnterParametersPrime(c *ParametersPrimeContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPrintPrime is called when entering the printPrime production.
	EnterPrintPrime(c *PrintPrimeContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterRet is called when entering the ret production.
	EnterRet(c *RetContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentsPrime is called when entering the argumentsPrime production.
	EnterArgumentsPrime(c *ArgumentsPrimeContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterLvaluePrime is called when entering the lvaluePrime production.
	EnterLvaluePrime(c *LvaluePrimeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionPrime is called when entering the expressionPrime production.
	EnterExpressionPrime(c *ExpressionPrimeContext)

	// EnterBoolTerm is called when entering the boolTerm production.
	EnterBoolTerm(c *BoolTermContext)

	// EnterBooltermPrime is called when entering the booltermPrime production.
	EnterBooltermPrime(c *BooltermPrimeContext)

	// EnterEqualTerm is called when entering the equalTerm production.
	EnterEqualTerm(c *EqualTermContext)

	// EnterEqualTermPrime is called when entering the equalTermPrime production.
	EnterEqualTermPrime(c *EqualTermPrimeContext)

	// EnterRelationTerm is called when entering the relationTerm production.
	EnterRelationTerm(c *RelationTermContext)

	// EnterRelationTermPrime is called when entering the relationTermPrime production.
	EnterRelationTermPrime(c *RelationTermPrimeContext)

	// EnterSimpleTerm is called when entering the simpleTerm production.
	EnterSimpleTerm(c *SimpleTermContext)

	// EnterSimpleTermPrime is called when entering the simpleTermPrime production.
	EnterSimpleTermPrime(c *SimpleTermPrimeContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTermPrime is called when entering the termPrime production.
	EnterTermPrime(c *TermPrimeContext)

	// EnterUnaryTerm is called when entering the unaryTerm production.
	EnterUnaryTerm(c *UnaryTermContext)

	// EnterSelectorTerm is called when entering the selectorTerm production.
	EnterSelectorTerm(c *SelectorTermContext)

	// EnterSelectorTermPrime is called when entering the selectorTermPrime production.
	EnterSelectorTermPrime(c *SelectorTermPrimeContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFactorPrime1 is called when entering the factorPrime1 production.
	EnterFactorPrime1(c *FactorPrime1Context)

	// EnterFactorPrime2 is called when entering the factorPrime2 production.
	EnterFactorPrime2(c *FactorPrime2Context)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitFieldsPrime is called when exiting the fieldsPrime production.
	ExitFieldsPrime(c *FieldsPrimeContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitIds is called when exiting the ids production.
	ExitIds(c *IdsContext)

	// ExitIdsPrime is called when exiting the idsPrime production.
	ExitIdsPrime(c *IdsPrimeContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParametersPrime is called when exiting the parametersPrime production.
	ExitParametersPrime(c *ParametersPrimeContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPrintPrime is called when exiting the printPrime production.
	ExitPrintPrime(c *PrintPrimeContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitRet is called when exiting the ret production.
	ExitRet(c *RetContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentsPrime is called when exiting the argumentsPrime production.
	ExitArgumentsPrime(c *ArgumentsPrimeContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitLvaluePrime is called when exiting the lvaluePrime production.
	ExitLvaluePrime(c *LvaluePrimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionPrime is called when exiting the expressionPrime production.
	ExitExpressionPrime(c *ExpressionPrimeContext)

	// ExitBoolTerm is called when exiting the boolTerm production.
	ExitBoolTerm(c *BoolTermContext)

	// ExitBooltermPrime is called when exiting the booltermPrime production.
	ExitBooltermPrime(c *BooltermPrimeContext)

	// ExitEqualTerm is called when exiting the equalTerm production.
	ExitEqualTerm(c *EqualTermContext)

	// ExitEqualTermPrime is called when exiting the equalTermPrime production.
	ExitEqualTermPrime(c *EqualTermPrimeContext)

	// ExitRelationTerm is called when exiting the relationTerm production.
	ExitRelationTerm(c *RelationTermContext)

	// ExitRelationTermPrime is called when exiting the relationTermPrime production.
	ExitRelationTermPrime(c *RelationTermPrimeContext)

	// ExitSimpleTerm is called when exiting the simpleTerm production.
	ExitSimpleTerm(c *SimpleTermContext)

	// ExitSimpleTermPrime is called when exiting the simpleTermPrime production.
	ExitSimpleTermPrime(c *SimpleTermPrimeContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTermPrime is called when exiting the termPrime production.
	ExitTermPrime(c *TermPrimeContext)

	// ExitUnaryTerm is called when exiting the unaryTerm production.
	ExitUnaryTerm(c *UnaryTermContext)

	// ExitSelectorTerm is called when exiting the selectorTerm production.
	ExitSelectorTerm(c *SelectorTermContext)

	// ExitSelectorTermPrime is called when exiting the selectorTermPrime production.
	ExitSelectorTermPrime(c *SelectorTermPrimeContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFactorPrime1 is called when exiting the factorPrime1 production.
	ExitFactorPrime1(c *FactorPrime1Context)

	// ExitFactorPrime2 is called when exiting the factorPrime2 production.
	ExitFactorPrime2(c *FactorPrime2Context)
}
