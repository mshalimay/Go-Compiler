// Code generated from GoliteLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package lexer

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
}

var GoliteLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golitelexerLexerInit() {
	staticData := &GoliteLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'type'", "'struct'", "'var'", "'int'", "'bool'", "'nil'", "'func'",
		"'true'", "'false'", "'delete'", "'new'", "'scan'", "'printf'", "'if'",
		"'else'", "'for'", "'return'", "'.'", "'{'", "'}'", "'('", "')'", "','",
		"';'", "'!'", "'/'", "'*'", "'-'", "'+'", "'='", "'<'", "'>'", "'<='",
		"'>='", "'=='", "'!='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "TYPE", "STRUCT", "VAR", "INT", "BOOL", "NIL", "FUNC", "TRUE", "FALSE",
		"DELETE", "NEW", "SCAN", "PRINTF", "IF", "ELSE", "FOR", "RETURN", "DOT",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "COMMA", "SEMICOLON", "NOT",
		"FSLASH", "ASTERISK", "MINUS", "PLUS", "EQUALS", "LT", "GT", "LTE",
		"GTE", "EQ", "NEQ", "AND", "OR", "ID", "NUMBER", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"TYPE", "STRUCT", "VAR", "INT", "BOOL", "NIL", "FUNC", "TRUE", "FALSE",
		"DELETE", "NEW", "SCAN", "PRINTF", "IF", "ELSE", "FOR", "RETURN", "DOT",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "COMMA", "SEMICOLON", "NOT",
		"FSLASH", "ASTERISK", "MINUS", "PLUS", "EQUALS", "LT", "GT", "LTE",
		"GTE", "EQ", "NEQ", "AND", "OR", "ID", "NUMBER", "STRING", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 263, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 5, 38, 225, 8, 38, 10,
		38, 12, 38, 228, 9, 38, 1, 39, 4, 39, 231, 8, 39, 11, 39, 12, 39, 232,
		1, 40, 1, 40, 5, 40, 237, 8, 40, 10, 40, 12, 40, 240, 9, 40, 1, 40, 1,
		40, 1, 41, 4, 41, 245, 8, 41, 11, 41, 12, 41, 246, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 5, 42, 255, 8, 42, 10, 42, 12, 42, 258, 9, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 2, 238, 256, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 1, 0, 4, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97,
		122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 267, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 1, 87,
		1, 0, 0, 0, 3, 92, 1, 0, 0, 0, 5, 99, 1, 0, 0, 0, 7, 103, 1, 0, 0, 0, 9,
		107, 1, 0, 0, 0, 11, 112, 1, 0, 0, 0, 13, 116, 1, 0, 0, 0, 15, 121, 1,
		0, 0, 0, 17, 126, 1, 0, 0, 0, 19, 132, 1, 0, 0, 0, 21, 139, 1, 0, 0, 0,
		23, 143, 1, 0, 0, 0, 25, 148, 1, 0, 0, 0, 27, 155, 1, 0, 0, 0, 29, 158,
		1, 0, 0, 0, 31, 163, 1, 0, 0, 0, 33, 167, 1, 0, 0, 0, 35, 174, 1, 0, 0,
		0, 37, 176, 1, 0, 0, 0, 39, 178, 1, 0, 0, 0, 41, 180, 1, 0, 0, 0, 43, 182,
		1, 0, 0, 0, 45, 184, 1, 0, 0, 0, 47, 186, 1, 0, 0, 0, 49, 188, 1, 0, 0,
		0, 51, 190, 1, 0, 0, 0, 53, 192, 1, 0, 0, 0, 55, 194, 1, 0, 0, 0, 57, 196,
		1, 0, 0, 0, 59, 198, 1, 0, 0, 0, 61, 200, 1, 0, 0, 0, 63, 202, 1, 0, 0,
		0, 65, 204, 1, 0, 0, 0, 67, 207, 1, 0, 0, 0, 69, 210, 1, 0, 0, 0, 71, 213,
		1, 0, 0, 0, 73, 216, 1, 0, 0, 0, 75, 219, 1, 0, 0, 0, 77, 222, 1, 0, 0,
		0, 79, 230, 1, 0, 0, 0, 81, 234, 1, 0, 0, 0, 83, 244, 1, 0, 0, 0, 85, 250,
		1, 0, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 121, 0, 0, 89, 90, 5, 112,
		0, 0, 90, 91, 5, 101, 0, 0, 91, 2, 1, 0, 0, 0, 92, 93, 5, 115, 0, 0, 93,
		94, 5, 116, 0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 117, 0, 0, 96, 97, 5,
		99, 0, 0, 97, 98, 5, 116, 0, 0, 98, 4, 1, 0, 0, 0, 99, 100, 5, 118, 0,
		0, 100, 101, 5, 97, 0, 0, 101, 102, 5, 114, 0, 0, 102, 6, 1, 0, 0, 0, 103,
		104, 5, 105, 0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 116, 0, 0, 106,
		8, 1, 0, 0, 0, 107, 108, 5, 98, 0, 0, 108, 109, 5, 111, 0, 0, 109, 110,
		5, 111, 0, 0, 110, 111, 5, 108, 0, 0, 111, 10, 1, 0, 0, 0, 112, 113, 5,
		110, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5, 108, 0, 0, 115, 12, 1,
		0, 0, 0, 116, 117, 5, 102, 0, 0, 117, 118, 5, 117, 0, 0, 118, 119, 5, 110,
		0, 0, 119, 120, 5, 99, 0, 0, 120, 14, 1, 0, 0, 0, 121, 122, 5, 116, 0,
		0, 122, 123, 5, 114, 0, 0, 123, 124, 5, 117, 0, 0, 124, 125, 5, 101, 0,
		0, 125, 16, 1, 0, 0, 0, 126, 127, 5, 102, 0, 0, 127, 128, 5, 97, 0, 0,
		128, 129, 5, 108, 0, 0, 129, 130, 5, 115, 0, 0, 130, 131, 5, 101, 0, 0,
		131, 18, 1, 0, 0, 0, 132, 133, 5, 100, 0, 0, 133, 134, 5, 101, 0, 0, 134,
		135, 5, 108, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 116, 0, 0, 137,
		138, 5, 101, 0, 0, 138, 20, 1, 0, 0, 0, 139, 140, 5, 110, 0, 0, 140, 141,
		5, 101, 0, 0, 141, 142, 5, 119, 0, 0, 142, 22, 1, 0, 0, 0, 143, 144, 5,
		115, 0, 0, 144, 145, 5, 99, 0, 0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 110,
		0, 0, 147, 24, 1, 0, 0, 0, 148, 149, 5, 112, 0, 0, 149, 150, 5, 114, 0,
		0, 150, 151, 5, 105, 0, 0, 151, 152, 5, 110, 0, 0, 152, 153, 5, 116, 0,
		0, 153, 154, 5, 102, 0, 0, 154, 26, 1, 0, 0, 0, 155, 156, 5, 105, 0, 0,
		156, 157, 5, 102, 0, 0, 157, 28, 1, 0, 0, 0, 158, 159, 5, 101, 0, 0, 159,
		160, 5, 108, 0, 0, 160, 161, 5, 115, 0, 0, 161, 162, 5, 101, 0, 0, 162,
		30, 1, 0, 0, 0, 163, 164, 5, 102, 0, 0, 164, 165, 5, 111, 0, 0, 165, 166,
		5, 114, 0, 0, 166, 32, 1, 0, 0, 0, 167, 168, 5, 114, 0, 0, 168, 169, 5,
		101, 0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 117, 0, 0, 171, 172, 5,
		114, 0, 0, 172, 173, 5, 110, 0, 0, 173, 34, 1, 0, 0, 0, 174, 175, 5, 46,
		0, 0, 175, 36, 1, 0, 0, 0, 176, 177, 5, 123, 0, 0, 177, 38, 1, 0, 0, 0,
		178, 179, 5, 125, 0, 0, 179, 40, 1, 0, 0, 0, 180, 181, 5, 40, 0, 0, 181,
		42, 1, 0, 0, 0, 182, 183, 5, 41, 0, 0, 183, 44, 1, 0, 0, 0, 184, 185, 5,
		44, 0, 0, 185, 46, 1, 0, 0, 0, 186, 187, 5, 59, 0, 0, 187, 48, 1, 0, 0,
		0, 188, 189, 5, 33, 0, 0, 189, 50, 1, 0, 0, 0, 190, 191, 5, 47, 0, 0, 191,
		52, 1, 0, 0, 0, 192, 193, 5, 42, 0, 0, 193, 54, 1, 0, 0, 0, 194, 195, 5,
		45, 0, 0, 195, 56, 1, 0, 0, 0, 196, 197, 5, 43, 0, 0, 197, 58, 1, 0, 0,
		0, 198, 199, 5, 61, 0, 0, 199, 60, 1, 0, 0, 0, 200, 201, 5, 60, 0, 0, 201,
		62, 1, 0, 0, 0, 202, 203, 5, 62, 0, 0, 203, 64, 1, 0, 0, 0, 204, 205, 5,
		60, 0, 0, 205, 206, 5, 61, 0, 0, 206, 66, 1, 0, 0, 0, 207, 208, 5, 62,
		0, 0, 208, 209, 5, 61, 0, 0, 209, 68, 1, 0, 0, 0, 210, 211, 5, 61, 0, 0,
		211, 212, 5, 61, 0, 0, 212, 70, 1, 0, 0, 0, 213, 214, 5, 33, 0, 0, 214,
		215, 5, 61, 0, 0, 215, 72, 1, 0, 0, 0, 216, 217, 5, 38, 0, 0, 217, 218,
		5, 38, 0, 0, 218, 74, 1, 0, 0, 0, 219, 220, 5, 124, 0, 0, 220, 221, 5,
		124, 0, 0, 221, 76, 1, 0, 0, 0, 222, 226, 7, 0, 0, 0, 223, 225, 7, 1, 0,
		0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 78, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 231, 7,
		2, 0, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 230, 1, 0, 0,
		0, 232, 233, 1, 0, 0, 0, 233, 80, 1, 0, 0, 0, 234, 238, 5, 34, 0, 0, 235,
		237, 9, 0, 0, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 239,
		1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 241, 242, 5, 34, 0, 0, 242, 82, 1, 0, 0, 0, 243, 245, 7, 3, 0, 0,
		244, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246,
		247, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 6, 41, 0, 0, 249, 84,
		1, 0, 0, 0, 250, 251, 5, 47, 0, 0, 251, 252, 5, 47, 0, 0, 252, 256, 1,
		0, 0, 0, 253, 255, 9, 0, 0, 0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0,
		0, 256, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 259, 1, 0, 0, 0, 258,
		256, 1, 0, 0, 0, 259, 260, 5, 10, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262,
		6, 42, 0, 0, 262, 86, 1, 0, 0, 0, 6, 0, 226, 232, 238, 246, 256, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &GoliteLexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoliteLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GoliteLexer.g4"

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerTYPE      = 1
	GoliteLexerSTRUCT    = 2
	GoliteLexerVAR       = 3
	GoliteLexerINT       = 4
	GoliteLexerBOOL      = 5
	GoliteLexerNIL       = 6
	GoliteLexerFUNC      = 7
	GoliteLexerTRUE      = 8
	GoliteLexerFALSE     = 9
	GoliteLexerDELETE    = 10
	GoliteLexerNEW       = 11
	GoliteLexerSCAN      = 12
	GoliteLexerPRINTF    = 13
	GoliteLexerIF        = 14
	GoliteLexerELSE      = 15
	GoliteLexerFOR       = 16
	GoliteLexerRETURN    = 17
	GoliteLexerDOT       = 18
	GoliteLexerLBRACE    = 19
	GoliteLexerRBRACE    = 20
	GoliteLexerLPAREN    = 21
	GoliteLexerRPAREN    = 22
	GoliteLexerCOMMA     = 23
	GoliteLexerSEMICOLON = 24
	GoliteLexerNOT       = 25
	GoliteLexerFSLASH    = 26
	GoliteLexerASTERISK  = 27
	GoliteLexerMINUS     = 28
	GoliteLexerPLUS      = 29
	GoliteLexerEQUALS    = 30
	GoliteLexerLT        = 31
	GoliteLexerGT        = 32
	GoliteLexerLTE       = 33
	GoliteLexerGTE       = 34
	GoliteLexerEQ        = 35
	GoliteLexerNEQ       = 36
	GoliteLexerAND       = 37
	GoliteLexerOR        = 38
	GoliteLexerID        = 39
	GoliteLexerNUMBER    = 40
	GoliteLexerSTRING    = 41
	GoliteLexerWS        = 42
	GoliteLexerCOMMENT   = 43
)
