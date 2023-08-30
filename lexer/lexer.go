package lexer

import (
	"github.com/antlr4-go/antlr/v4"
	"fmt"
	"golite/context"
)


// Lexer is an interface to expose to public desired methods/fields of lexerWrapper
type Lexer interface {
	GetTokenStream() 			*antlr.CommonTokenStream
	GetTokenRule(tokenID int) 	string
	GetErrors() 				[]*context.CompilerError
	FillTokenStream()
	PrintAllTokensTexts()
	PrintAllTokenRules()
	PrintAllTokens()
}

// lexerWrapper is a wrapper around the antlr lexer.
// it is a private struct, because contains methods and fields that should not be exposed
// the Lexer interface will be used to expose desired methods/fields of this struct
type lexerWrapper struct {
	*antlr.DefaultErrorListener 		// implements antlr.ErrorListener
	antlrLexer 		*GoliteLexer
	errors 			[]*context.CompilerError
	stream 			*antlr.CommonTokenStream
}


// NewLexer creates a new Lexer from the input source path contaning the golite program
func NewLexer(inputSourcePath string) Lexer{
	input, _ := antlr.NewFileStream(inputSourcePath)
	
	lexer := &lexerWrapper{antlr.NewDefaultErrorListener(), nil, nil,  nil}
	
	// instaitate antlr lexer and override default error listener
	antlrLexer := NewGoliteLexer(input)
	antlrLexer.RemoveErrorListeners()
	antlrLexer.AddErrorListener(lexer)

	// fill lexer fields
	lexer.antlrLexer = antlrLexer
	lexer.stream = antlr.NewCommonTokenStream(antlrLexer, 0)

	return lexer
}


// Rreturns the token stream of the lexer
func (lexer *lexerWrapper) GetTokenStream() *antlr.CommonTokenStream {
	return lexer.stream
}

// Fills the lexer with ANTLR token stream
func (lexer *lexerWrapper) FillTokenStream() {
	lexer.stream.Fill()
}

// GetTokenRule returns the rule name of the token with the given tokenID
// eg. tokenID = 1 -> "INT" , tokenID = 2 -> "ID"
func (lexer *lexerWrapper) GetTokenRule(tokenID int) string {
	// -1 is EOF in ANTLR
	if tokenID == -1 { 
		return "EOF"
	}
	return lexer.antlrLexer.RuleNames[tokenID-1]
}

//==============================================================================
// ErrorListener methods
//==============================================================================
// Override antlr.ErrorListener methods
func (lexer *lexerWrapper) SyntaxError(recognizer antlr.Recognizer, 
	offendingSymbol interface{}, line, column int, msg string, 
	e antlr.RecognitionException) {
	lexer.errors = append(lexer.errors, &context.CompilerError{
		Line: line,
		Column: column, 
		Msg: msg, 
		Phase: context.LEXER,
	})
}

func (lexer *lexerWrapper) GetErrors() []*context.CompilerError{
	return lexer.errors
}


//==============================================================================
// methods for printing tokens
//==============================================================================

// print all tokens identified by the lexer in their original text form
func (lexer *lexerWrapper) PrintAllTokensTexts() {
	for _, token := range lexer.GetTokenStream().GetAllTokens() {
        fmt.Println(token.GetText())
 	}
}

func (lexer *lexerWrapper) PrintAllTokenRules() {
	for _, token := range lexer.GetTokenStream().GetAllTokens() {
		println(lexer.GetTokenRule(token.GetTokenType()))
	}
}

// print all tokens identified by the lexer in the symbols defined in the grammar
// in the format: Token.SYMBOL_NAME(line [,value]) where [...] = optional
// eg.: Token.TYPE(4); Token.ID(4, "myVal"); Token.NUMBER(7, 9); Token.STRUCT(4); etc

func (lexer *lexerWrapper) PrintAllTokens() {
	antlrLexer := lexer.antlrLexer
	
	var tokenType int;
	var tokenTxt string;
	var tokenSymbol string;

	for _, token := range lexer.GetTokenStream().GetAllTokens() {
		
		tokenType = token.GetTokenType()
		
		if tokenType == antlr.TokenEOF {
			fmt.Printf("Token.%v(%v)\n", token.GetText(), token.GetLine())
			break
		}
		
		tokenSymbol = antlrLexer.SymbolicNames[tokenType]
		tokenTxt = token.GetText()

		switch tokenSymbol {
			case "ID":
				fmt.Printf("Token.%v(%v, \"%s\")\n", tokenSymbol, token.GetLine(), tokenTxt)
			case "NUMBER":
				fmt.Printf("Token.%v(%v, %s)\n", tokenSymbol, token.GetLine(), tokenTxt)
			case "STRING":
				fmt.Printf("Token.%v(%v, %s)\n", tokenSymbol, token.GetLine(), tokenTxt)
			default:
				fmt.Printf("Token.%v(%v)\n", tokenSymbol, token.GetLine())
		}
	}
}

