// Code generated from GoliteParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GoliteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoliteParser struct {
	*antlr.BaseParser
}

var GoliteParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &GoliteParserParserStaticData
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
		"program", "types", "typeDeclaration", "fields", "fieldsPrime", "decl",
		"type", "declarations", "declaration", "ids", "idsPrime", "functions",
		"function", "parameters", "parametersPrime", "returnType", "statements",
		"statement", "block", "delete", "assignment", "print", "printPrime",
		"conditional", "loop", "ret", "invocation", "arguments", "argumentsPrime",
		"lvalue", "lvaluePrime", "expression", "expressionPrime", "boolTerm",
		"booltermPrime", "equalTerm", "equalTermPrime", "relationTerm", "relationTermPrime",
		"simpleTerm", "simpleTermPrime", "term", "termPrime", "unaryTerm", "selectorTerm",
		"selectorTermPrime", "factor", "factorPrime1", "factorPrime2",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 385, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 105, 8,
		1, 10, 1, 12, 1, 108, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 5, 3, 121, 8, 3, 10, 3, 12, 3, 124, 9, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 136, 8, 6, 1,
		7, 5, 7, 139, 8, 7, 10, 7, 12, 7, 142, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 5, 9, 151, 8, 9, 10, 9, 12, 9, 154, 9, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 5, 11, 160, 8, 11, 10, 11, 12, 11, 163, 9, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 169, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 5, 13, 179, 8, 13, 10, 13, 12, 13, 182, 9, 13, 3,
		13, 184, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		5, 16, 194, 8, 16, 10, 16, 12, 16, 197, 9, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 207, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 221,
		8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 229, 8, 21, 10,
		21, 12, 21, 232, 9, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 247, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 257, 8, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 268, 8,
		27, 10, 27, 12, 27, 271, 9, 27, 3, 27, 273, 8, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 5, 29, 282, 8, 29, 10, 29, 12, 29, 285, 9,
		29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 292, 8, 31, 10, 31, 12, 31,
		295, 9, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 302, 8, 33, 10, 33,
		12, 33, 305, 9, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 312, 8, 35,
		10, 35, 12, 35, 315, 9, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 322,
		8, 37, 10, 37, 12, 37, 325, 9, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 5,
		39, 332, 8, 39, 10, 39, 12, 39, 335, 9, 39, 1, 40, 1, 40, 1, 40, 1, 41,
		1, 41, 5, 41, 342, 8, 41, 10, 41, 12, 41, 345, 9, 41, 1, 42, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 355, 8, 43, 1, 44, 1, 44,
		5, 44, 359, 8, 44, 10, 44, 12, 44, 362, 9, 44, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 375, 8, 46,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 3, 48, 383, 8, 48, 1, 48, 0,
		0, 49, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 0, 4, 1, 0, 35, 36,
		1, 0, 31, 34, 1, 0, 28, 29, 1, 0, 26, 27, 376, 0, 98, 1, 0, 0, 0, 2, 106,
		1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0,
		10, 128, 1, 0, 0, 0, 12, 135, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 143,
		1, 0, 0, 0, 18, 148, 1, 0, 0, 0, 20, 155, 1, 0, 0, 0, 22, 161, 1, 0, 0,
		0, 24, 164, 1, 0, 0, 0, 26, 175, 1, 0, 0, 0, 28, 187, 1, 0, 0, 0, 30, 190,
		1, 0, 0, 0, 32, 195, 1, 0, 0, 0, 34, 206, 1, 0, 0, 0, 36, 208, 1, 0, 0,
		0, 38, 212, 1, 0, 0, 0, 40, 216, 1, 0, 0, 0, 42, 224, 1, 0, 0, 0, 44, 236,
		1, 0, 0, 0, 46, 239, 1, 0, 0, 0, 48, 248, 1, 0, 0, 0, 50, 254, 1, 0, 0,
		0, 52, 260, 1, 0, 0, 0, 54, 264, 1, 0, 0, 0, 56, 276, 1, 0, 0, 0, 58, 279,
		1, 0, 0, 0, 60, 286, 1, 0, 0, 0, 62, 289, 1, 0, 0, 0, 64, 296, 1, 0, 0,
		0, 66, 299, 1, 0, 0, 0, 68, 306, 1, 0, 0, 0, 70, 309, 1, 0, 0, 0, 72, 316,
		1, 0, 0, 0, 74, 319, 1, 0, 0, 0, 76, 326, 1, 0, 0, 0, 78, 329, 1, 0, 0,
		0, 80, 336, 1, 0, 0, 0, 82, 339, 1, 0, 0, 0, 84, 346, 1, 0, 0, 0, 86, 354,
		1, 0, 0, 0, 88, 356, 1, 0, 0, 0, 90, 363, 1, 0, 0, 0, 92, 374, 1, 0, 0,
		0, 94, 376, 1, 0, 0, 0, 96, 380, 1, 0, 0, 0, 98, 99, 3, 2, 1, 0, 99, 100,
		3, 14, 7, 0, 100, 101, 3, 22, 11, 0, 101, 102, 5, 0, 0, 1, 102, 1, 1, 0,
		0, 0, 103, 105, 3, 4, 2, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0,
		106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 3, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 110, 5, 1, 0, 0, 110, 111, 5, 39, 0, 0, 111, 112, 5, 2,
		0, 0, 112, 113, 5, 19, 0, 0, 113, 114, 3, 6, 3, 0, 114, 115, 5, 20, 0,
		0, 115, 116, 5, 24, 0, 0, 116, 5, 1, 0, 0, 0, 117, 118, 3, 10, 5, 0, 118,
		122, 5, 24, 0, 0, 119, 121, 3, 8, 4, 0, 120, 119, 1, 0, 0, 0, 121, 124,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 7, 1, 0, 0,
		0, 124, 122, 1, 0, 0, 0, 125, 126, 3, 10, 5, 0, 126, 127, 5, 24, 0, 0,
		127, 9, 1, 0, 0, 0, 128, 129, 5, 39, 0, 0, 129, 130, 3, 12, 6, 0, 130,
		11, 1, 0, 0, 0, 131, 136, 5, 4, 0, 0, 132, 136, 5, 5, 0, 0, 133, 134, 5,
		27, 0, 0, 134, 136, 5, 39, 0, 0, 135, 131, 1, 0, 0, 0, 135, 132, 1, 0,
		0, 0, 135, 133, 1, 0, 0, 0, 136, 13, 1, 0, 0, 0, 137, 139, 3, 16, 8, 0,
		138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 15, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5,
		3, 0, 0, 144, 145, 3, 18, 9, 0, 145, 146, 3, 12, 6, 0, 146, 147, 5, 24,
		0, 0, 147, 17, 1, 0, 0, 0, 148, 152, 5, 39, 0, 0, 149, 151, 3, 20, 10,
		0, 150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 19, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5,
		23, 0, 0, 156, 157, 5, 39, 0, 0, 157, 21, 1, 0, 0, 0, 158, 160, 3, 24,
		12, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 23, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165,
		5, 7, 0, 0, 165, 166, 5, 39, 0, 0, 166, 168, 3, 26, 13, 0, 167, 169, 3,
		30, 15, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 171, 5, 19, 0, 0, 171, 172, 3, 14, 7, 0, 172, 173, 3, 32, 16,
		0, 173, 174, 5, 20, 0, 0, 174, 25, 1, 0, 0, 0, 175, 183, 5, 21, 0, 0, 176,
		180, 3, 10, 5, 0, 177, 179, 3, 28, 14, 0, 178, 177, 1, 0, 0, 0, 179, 182,
		1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 184, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 183, 176, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 186, 5, 22, 0, 0, 186, 27, 1, 0, 0, 0, 187,
		188, 5, 23, 0, 0, 188, 189, 3, 10, 5, 0, 189, 29, 1, 0, 0, 0, 190, 191,
		3, 12, 6, 0, 191, 31, 1, 0, 0, 0, 192, 194, 3, 34, 17, 0, 193, 192, 1,
		0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0,
		0, 196, 33, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 207, 3, 36, 18, 0, 199,
		207, 3, 40, 20, 0, 200, 207, 3, 42, 21, 0, 201, 207, 3, 38, 19, 0, 202,
		207, 3, 46, 23, 0, 203, 207, 3, 48, 24, 0, 204, 207, 3, 50, 25, 0, 205,
		207, 3, 52, 26, 0, 206, 198, 1, 0, 0, 0, 206, 199, 1, 0, 0, 0, 206, 200,
		1, 0, 0, 0, 206, 201, 1, 0, 0, 0, 206, 202, 1, 0, 0, 0, 206, 203, 1, 0,
		0, 0, 206, 204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 35, 1, 0, 0, 0,
		208, 209, 5, 19, 0, 0, 209, 210, 3, 32, 16, 0, 210, 211, 5, 20, 0, 0, 211,
		37, 1, 0, 0, 0, 212, 213, 5, 10, 0, 0, 213, 214, 3, 62, 31, 0, 214, 215,
		5, 24, 0, 0, 215, 39, 1, 0, 0, 0, 216, 217, 3, 58, 29, 0, 217, 220, 5,
		30, 0, 0, 218, 221, 3, 62, 31, 0, 219, 221, 5, 12, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 24, 0, 0,
		223, 41, 1, 0, 0, 0, 224, 225, 5, 13, 0, 0, 225, 226, 5, 21, 0, 0, 226,
		230, 5, 41, 0, 0, 227, 229, 3, 44, 22, 0, 228, 227, 1, 0, 0, 0, 229, 232,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 5, 22, 0, 0, 234, 235, 5, 24, 0,
		0, 235, 43, 1, 0, 0, 0, 236, 237, 5, 23, 0, 0, 237, 238, 3, 62, 31, 0,
		238, 45, 1, 0, 0, 0, 239, 240, 5, 14, 0, 0, 240, 241, 5, 21, 0, 0, 241,
		242, 3, 62, 31, 0, 242, 243, 5, 22, 0, 0, 243, 246, 3, 36, 18, 0, 244,
		245, 5, 15, 0, 0, 245, 247, 3, 36, 18, 0, 246, 244, 1, 0, 0, 0, 246, 247,
		1, 0, 0, 0, 247, 47, 1, 0, 0, 0, 248, 249, 5, 16, 0, 0, 249, 250, 5, 21,
		0, 0, 250, 251, 3, 62, 31, 0, 251, 252, 5, 22, 0, 0, 252, 253, 3, 36, 18,
		0, 253, 49, 1, 0, 0, 0, 254, 256, 5, 17, 0, 0, 255, 257, 3, 62, 31, 0,
		256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258,
		259, 5, 24, 0, 0, 259, 51, 1, 0, 0, 0, 260, 261, 5, 39, 0, 0, 261, 262,
		3, 54, 27, 0, 262, 263, 5, 24, 0, 0, 263, 53, 1, 0, 0, 0, 264, 272, 5,
		21, 0, 0, 265, 269, 3, 62, 31, 0, 266, 268, 3, 56, 28, 0, 267, 266, 1,
		0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0,
		0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 265, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 5, 22, 0, 0, 275, 55,
		1, 0, 0, 0, 276, 277, 5, 23, 0, 0, 277, 278, 3, 62, 31, 0, 278, 57, 1,
		0, 0, 0, 279, 283, 5, 39, 0, 0, 280, 282, 3, 60, 30, 0, 281, 280, 1, 0,
		0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0,
		284, 59, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 287, 5, 18, 0, 0, 287,
		288, 5, 39, 0, 0, 288, 61, 1, 0, 0, 0, 289, 293, 3, 66, 33, 0, 290, 292,
		3, 64, 32, 0, 291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1,
		0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 63, 1, 0, 0, 0, 295, 293, 1, 0, 0,
		0, 296, 297, 5, 38, 0, 0, 297, 298, 3, 66, 33, 0, 298, 65, 1, 0, 0, 0,
		299, 303, 3, 70, 35, 0, 300, 302, 3, 68, 34, 0, 301, 300, 1, 0, 0, 0, 302,
		305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 67, 1,
		0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 307, 5, 37, 0, 0, 307, 308, 3, 70,
		35, 0, 308, 69, 1, 0, 0, 0, 309, 313, 3, 74, 37, 0, 310, 312, 3, 72, 36,
		0, 311, 310, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313,
		314, 1, 0, 0, 0, 314, 71, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 7,
		0, 0, 0, 317, 318, 3, 74, 37, 0, 318, 73, 1, 0, 0, 0, 319, 323, 3, 78,
		39, 0, 320, 322, 3, 76, 38, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0,
		0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 75, 1, 0, 0, 0, 325,
		323, 1, 0, 0, 0, 326, 327, 7, 1, 0, 0, 327, 328, 3, 78, 39, 0, 328, 77,
		1, 0, 0, 0, 329, 333, 3, 82, 41, 0, 330, 332, 3, 80, 40, 0, 331, 330, 1,
		0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0,
		0, 334, 79, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337, 7, 2, 0, 0, 337,
		338, 3, 82, 41, 0, 338, 81, 1, 0, 0, 0, 339, 343, 3, 86, 43, 0, 340, 342,
		3, 84, 42, 0, 341, 340, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1,
		0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 83, 1, 0, 0, 0, 345, 343, 1, 0, 0,
		0, 346, 347, 7, 3, 0, 0, 347, 348, 3, 86, 43, 0, 348, 85, 1, 0, 0, 0, 349,
		350, 5, 25, 0, 0, 350, 355, 3, 88, 44, 0, 351, 352, 5, 28, 0, 0, 352, 355,
		3, 88, 44, 0, 353, 355, 3, 88, 44, 0, 354, 349, 1, 0, 0, 0, 354, 351, 1,
		0, 0, 0, 354, 353, 1, 0, 0, 0, 355, 87, 1, 0, 0, 0, 356, 360, 3, 92, 46,
		0, 357, 359, 3, 90, 45, 0, 358, 357, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0,
		360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 89, 1, 0, 0, 0, 362, 360,
		1, 0, 0, 0, 363, 364, 5, 18, 0, 0, 364, 365, 5, 39, 0, 0, 365, 91, 1, 0,
		0, 0, 366, 375, 3, 94, 47, 0, 367, 375, 3, 96, 48, 0, 368, 375, 5, 40,
		0, 0, 369, 370, 5, 11, 0, 0, 370, 375, 5, 39, 0, 0, 371, 375, 5, 8, 0,
		0, 372, 375, 5, 9, 0, 0, 373, 375, 5, 6, 0, 0, 374, 366, 1, 0, 0, 0, 374,
		367, 1, 0, 0, 0, 374, 368, 1, 0, 0, 0, 374, 369, 1, 0, 0, 0, 374, 371,
		1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 373, 1, 0, 0, 0, 375, 93, 1, 0,
		0, 0, 376, 377, 5, 21, 0, 0, 377, 378, 3, 62, 31, 0, 378, 379, 5, 22, 0,
		0, 379, 95, 1, 0, 0, 0, 380, 382, 5, 39, 0, 0, 381, 383, 3, 54, 27, 0,
		382, 381, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 97, 1, 0, 0, 0, 28, 106,
		122, 135, 140, 152, 161, 168, 180, 183, 195, 206, 220, 230, 246, 256, 269,
		272, 283, 293, 303, 313, 323, 333, 343, 354, 360, 374, 382,
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

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &GoliteParserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoliteParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoliteParser.g4"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF       = antlr.TokenEOF
	GoliteParserTYPE      = 1
	GoliteParserSTRUCT    = 2
	GoliteParserVAR       = 3
	GoliteParserINT       = 4
	GoliteParserBOOL      = 5
	GoliteParserNIL       = 6
	GoliteParserFUNC      = 7
	GoliteParserTRUE      = 8
	GoliteParserFALSE     = 9
	GoliteParserDELETE    = 10
	GoliteParserNEW       = 11
	GoliteParserSCAN      = 12
	GoliteParserPRINTF    = 13
	GoliteParserIF        = 14
	GoliteParserELSE      = 15
	GoliteParserFOR       = 16
	GoliteParserRETURN    = 17
	GoliteParserDOT       = 18
	GoliteParserLBRACE    = 19
	GoliteParserRBRACE    = 20
	GoliteParserLPAREN    = 21
	GoliteParserRPAREN    = 22
	GoliteParserCOMMA     = 23
	GoliteParserSEMICOLON = 24
	GoliteParserNOT       = 25
	GoliteParserFSLASH    = 26
	GoliteParserASTERISK  = 27
	GoliteParserMINUS     = 28
	GoliteParserPLUS      = 29
	GoliteParserEQUALS    = 30
	GoliteParserLT        = 31
	GoliteParserGT        = 32
	GoliteParserLTE       = 33
	GoliteParserGTE       = 34
	GoliteParserEQ        = 35
	GoliteParserNEQ       = 36
	GoliteParserAND       = 37
	GoliteParserOR        = 38
	GoliteParserID        = 39
	GoliteParserNUMBER    = 40
	GoliteParserSTRING    = 41
	GoliteParserWS        = 42
	GoliteParserCOMMENT   = 43
)

// GoliteParser rules.
const (
	GoliteParserRULE_program           = 0
	GoliteParserRULE_types             = 1
	GoliteParserRULE_typeDeclaration   = 2
	GoliteParserRULE_fields            = 3
	GoliteParserRULE_fieldsPrime       = 4
	GoliteParserRULE_decl              = 5
	GoliteParserRULE_type              = 6
	GoliteParserRULE_declarations      = 7
	GoliteParserRULE_declaration       = 8
	GoliteParserRULE_ids               = 9
	GoliteParserRULE_idsPrime          = 10
	GoliteParserRULE_functions         = 11
	GoliteParserRULE_function          = 12
	GoliteParserRULE_parameters        = 13
	GoliteParserRULE_parametersPrime   = 14
	GoliteParserRULE_returnType        = 15
	GoliteParserRULE_statements        = 16
	GoliteParserRULE_statement         = 17
	GoliteParserRULE_block             = 18
	GoliteParserRULE_delete            = 19
	GoliteParserRULE_assignment        = 20
	GoliteParserRULE_print             = 21
	GoliteParserRULE_printPrime        = 22
	GoliteParserRULE_conditional       = 23
	GoliteParserRULE_loop              = 24
	GoliteParserRULE_ret               = 25
	GoliteParserRULE_invocation        = 26
	GoliteParserRULE_arguments         = 27
	GoliteParserRULE_argumentsPrime    = 28
	GoliteParserRULE_lvalue            = 29
	GoliteParserRULE_lvaluePrime       = 30
	GoliteParserRULE_expression        = 31
	GoliteParserRULE_expressionPrime   = 32
	GoliteParserRULE_boolTerm          = 33
	GoliteParserRULE_booltermPrime     = 34
	GoliteParserRULE_equalTerm         = 35
	GoliteParserRULE_equalTermPrime    = 36
	GoliteParserRULE_relationTerm      = 37
	GoliteParserRULE_relationTermPrime = 38
	GoliteParserRULE_simpleTerm        = 39
	GoliteParserRULE_simpleTermPrime   = 40
	GoliteParserRULE_term              = 41
	GoliteParserRULE_termPrime         = 42
	GoliteParserRULE_unaryTerm         = 43
	GoliteParserRULE_selectorTerm      = 44
	GoliteParserRULE_selectorTermPrime = 45
	GoliteParserRULE_factor            = 46
	GoliteParserRULE_factorPrime1      = 47
	GoliteParserRULE_factorPrime2      = 48
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Types() ITypesContext
	Declarations() IDeclarationsContext
	Functions() IFunctionsContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Types()
	}
	{
		p.SetState(99)
		p.Declarations()
	}
	{
		p.SetState(100)
		p.Functions()
	}
	{
		p.SetState(101)
		p.Match(GoliteParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeDeclaration() []ITypeDeclarationContext
	TypeDeclaration(i int) ITypeDeclarationContext

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypesContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *GoliteParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_types)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(103)
			p.TypeDeclaration()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Fields() IFieldsContext
	RBRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
	return p
}

func InitEmptyTypeDeclarationContext(p *TypeDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypeDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeDeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypeDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *TypeDeclarationContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypeDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *GoliteParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_typeDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(GoliteParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(GoliteParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Fields()
	}
	{
		p.SetState(114)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl() IDeclContext
	SEMICOLON() antlr.TerminalNode
	AllFieldsPrime() []IFieldsPrimeContext
	FieldsPrime(i int) IFieldsPrimeContext

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func InitEmptyFieldsContext(p *FieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsContext) AllFieldsPrime() []IFieldsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IFieldsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsPrimeContext); ok {
			tst[i] = t.(IFieldsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) FieldsPrime(i int) IFieldsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsPrimeContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Decl()
	}
	{
		p.SetState(118)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserID {
		{
			p.SetState(119)
			p.FieldsPrime()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsPrimeContext is an interface to support dynamic dispatch.
type IFieldsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl() IDeclContext
	SEMICOLON() antlr.TerminalNode

	// IsFieldsPrimeContext differentiates from other interfaces.
	IsFieldsPrimeContext()
}

type FieldsPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsPrimeContext() *FieldsPrimeContext {
	var p = new(FieldsPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fieldsPrime
	return p
}

func InitEmptyFieldsPrimeContext(p *FieldsPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fieldsPrime
}

func (*FieldsPrimeContext) IsFieldsPrimeContext() {}

func NewFieldsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsPrimeContext {
	var p = new(FieldsPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fieldsPrime

	return p
}

func (s *FieldsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsPrimeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFieldsPrime(s)
	}
}

func (s *FieldsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFieldsPrime(s)
	}
}

func (p *GoliteParser) FieldsPrime() (localctx IFieldsPrimeContext) {
	localctx = NewFieldsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_fieldsPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Decl()
	}
	{
		p.SetState(126)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_() ITypeContext

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	ASTERISK() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL, 0)
}

func (s *TypeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_type)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(GoliteParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(GoliteParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserASTERISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(GoliteParserASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(GoliteParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
	return p
}

func InitEmptyDeclarationsContext(p *DeclarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (p *GoliteParser) Declarations() (localctx IDeclarationsContext) {
	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(137)
			p.Declaration()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Ids() IIdsContext
	Type_() ITypeContext
	SEMICOLON() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) Ids() IIdsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsContext)
}

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(GoliteParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Ids()
	}
	{
		p.SetState(145)
		p.Type_()
	}
	{
		p.SetState(146)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdsContext is an interface to support dynamic dispatch.
type IIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllIdsPrime() []IIdsPrimeContext
	IdsPrime(i int) IIdsPrimeContext

	// IsIdsContext differentiates from other interfaces.
	IsIdsContext()
}

type IdsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdsContext() *IdsContext {
	var p = new(IdsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
	return p
}

func InitEmptyIdsContext(p *IdsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
}

func (*IdsContext) IsIdsContext() {}

func NewIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsContext {
	var p = new(IdsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ids

	return p
}

func (s *IdsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *IdsContext) AllIdsPrime() []IIdsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IIdsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdsPrimeContext); ok {
			tst[i] = t.(IIdsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *IdsContext) IdsPrime(i int) IIdsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsPrimeContext)
}

func (s *IdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIds(s)
	}
}

func (s *IdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIds(s)
	}
}

func (p *GoliteParser) Ids() (localctx IIdsContext) {
	localctx = NewIdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_ids)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(149)
			p.IdsPrime()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdsPrimeContext is an interface to support dynamic dispatch.
type IIdsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsIdsPrimeContext differentiates from other interfaces.
	IsIdsPrimeContext()
}

type IdsPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdsPrimeContext() *IdsPrimeContext {
	var p = new(IdsPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_idsPrime
	return p
}

func InitEmptyIdsPrimeContext(p *IdsPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_idsPrime
}

func (*IdsPrimeContext) IsIdsPrimeContext() {}

func NewIdsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsPrimeContext {
	var p = new(IdsPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_idsPrime

	return p
}

func (s *IdsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *IdsPrimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *IdsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIdsPrime(s)
	}
}

func (s *IdsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIdsPrime(s)
	}
}

func (p *GoliteParser) IdsPrime() (localctx IIdsPrimeContext) {
	localctx = NewIdsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_idsPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(GoliteParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
	return p
}

func InitEmptyFunctionsContext(p *FunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionsContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *GoliteParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_functions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNC {
		{
			p.SetState(158)
			p.Function()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	Parameters() IParametersContext
	LBRACE() antlr.TerminalNode
	Declarations() IDeclarationsContext
	Statements() IStatementsContext
	RBRACE() antlr.TerminalNode
	ReturnType() IReturnTypeContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *FunctionContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *FunctionContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *FunctionContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(GoliteParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Parameters()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134217776) != 0 {
		{
			p.SetState(167)
			p.ReturnType()
		}

	}
	{
		p.SetState(170)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Declarations()
	}
	{
		p.SetState(172)
		p.Statements()
	}
	{
		p.SetState(173)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Decl() IDeclContext
	AllParametersPrime() []IParametersPrimeContext
	ParametersPrime(i int) IParametersPrimeContext

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ParametersContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersContext) AllParametersPrime() []IParametersPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametersPrimeContext); ok {
			len++
		}
	}

	tst := make([]IParametersPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametersPrimeContext); ok {
			tst[i] = t.(IParametersPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) ParametersPrime(i int) IParametersPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersPrimeContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GoliteParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserID {
		{
			p.SetState(176)
			p.Decl()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(177)
				p.ParametersPrime()
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(185)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersPrimeContext is an interface to support dynamic dispatch.
type IParametersPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Decl() IDeclContext

	// IsParametersPrimeContext differentiates from other interfaces.
	IsParametersPrimeContext()
}

type ParametersPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersPrimeContext() *ParametersPrimeContext {
	var p = new(ParametersPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parametersPrime
	return p
}

func InitEmptyParametersPrimeContext(p *ParametersPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parametersPrime
}

func (*ParametersPrimeContext) IsParametersPrimeContext() {}

func NewParametersPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersPrimeContext {
	var p = new(ParametersPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parametersPrime

	return p
}

func (s *ParametersPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ParametersPrimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParametersPrime(s)
	}
}

func (s *ParametersPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParametersPrime(s)
	}
}

func (p *GoliteParser) ParametersPrime() (localctx IParametersPrimeContext) {
	localctx = NewParametersPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_parametersPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(GoliteParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Decl()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
	return p
}

func InitEmptyReturnTypeContext(p *ReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *GoliteParser) ReturnType() (localctx IReturnTypeContext) {
	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_returnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GoliteParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549756560384) != 0 {
		{
			p.SetState(192)
			p.Statement()
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	Assignment() IAssignmentContext
	Print_() IPrintContext
	Delete_() IDeleteContext
	Conditional() IConditionalContext
	Loop() ILoopContext
	Ret() IRetContext
	Invocation() IInvocationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Ret() IRetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_statement)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Print_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.Delete_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(203)
			p.Loop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(204)
			p.Ret()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(205)
			p.Invocation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RBRACE() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Statements()
	}
	{
		p.SetState(210)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func InitEmptyDeleteContext(p *DeleteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_delete)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(GoliteParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Expression()
	}
	{
		p.SetState(214)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lvalue() ILvalueContext
	EQUALS() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext
	SCAN() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQUALS, 0)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Lvalue()
	}
	{
		p.SetState(217)
		p.Match(GoliteParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserNIL, GoliteParserTRUE, GoliteParserFALSE, GoliteParserNEW, GoliteParserLPAREN, GoliteParserNOT, GoliteParserMINUS, GoliteParserID, GoliteParserNUMBER:
		{
			p.SetState(218)
			p.Expression()
		}

	case GoliteParserSCAN:
		{
			p.SetState(219)
			p.Match(GoliteParserSCAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(222)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINTF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllPrintPrime() []IPrintPrimeContext
	PrintPrime(i int) IPrintPrimeContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINTF, 0)
}

func (s *PrintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *PrintContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING, 0)
}

func (s *PrintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) AllPrintPrime() []IPrintPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintPrimeContext); ok {
			len++
		}
	}

	tst := make([]IPrintPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintPrimeContext); ok {
			tst[i] = t.(IPrintPrimeContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) PrintPrime(i int) IPrintPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintPrimeContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(GoliteParserPRINTF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Match(GoliteParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(227)
			p.PrintPrime()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintPrimeContext is an interface to support dynamic dispatch.
type IPrintPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrintPrimeContext differentiates from other interfaces.
	IsPrintPrimeContext()
}

type PrintPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintPrimeContext() *PrintPrimeContext {
	var p = new(PrintPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_printPrime
	return p
}

func InitEmptyPrintPrimeContext(p *PrintPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_printPrime
}

func (*PrintPrimeContext) IsPrintPrimeContext() {}

func NewPrintPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintPrimeContext {
	var p = new(PrintPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_printPrime

	return p
}

func (s *PrintPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *PrintPrimeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrintPrime(s)
	}
}

func (s *PrintPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrintPrime(s)
	}
}

func (p *GoliteParser) PrintPrime() (localctx IPrintPrimeContext) {
	localctx = NewPrintPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_printPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(GoliteParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func InitEmptyConditionalContext(p *ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ConditionalContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_conditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(GoliteParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Expression()
	}
	{
		p.SetState(242)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Block()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(244)
			p.Match(GoliteParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(GoliteParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Expression()
	}
	{
		p.SetState(251)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetContext is an interface to support dynamic dispatch.
type IRetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsRetContext differentiates from other interfaces.
	IsRetContext()
}

type RetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetContext() *RetContext {
	var p = new(RetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ret
	return p
}

func InitEmptyRetContext(p *RetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ret
}

func (*RetContext) IsRetContext() {}

func NewRetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetContext {
	var p = new(RetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ret

	return p
}

func (s *RetContext) GetParser() antlr.Parser { return s.parser }

func (s *RetContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRETURN, 0)
}

func (s *RetContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *RetContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRet(s)
	}
}

func (s *RetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRet(s)
	}
}

func (p *GoliteParser) Ret() (localctx IRetContext) {
	localctx = NewRetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_ret)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(GoliteParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1649571531584) != 0 {
		{
			p.SetState(255)
			p.Expression()
		}

	}
	{
		p.SetState(258)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Arguments() IArgumentsContext
	SEMICOLON() antlr.TerminalNode

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func InitEmptyInvocationContext(p *InvocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_invocation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Arguments()
	}
	{
		p.SetState(262)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	AllArgumentsPrime() []IArgumentsPrimeContext
	ArgumentsPrime(i int) IArgumentsPrimeContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ArgumentsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllArgumentsPrime() []IArgumentsPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentsPrimeContext); ok {
			len++
		}
	}

	tst := make([]IArgumentsPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentsPrimeContext); ok {
			tst[i] = t.(IArgumentsPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) ArgumentsPrime(i int) IArgumentsPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsPrimeContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GoliteParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1649571531584) != 0 {
		{
			p.SetState(265)
			p.Expression()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(266)
				p.ArgumentsPrime()
			}

			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(274)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsPrimeContext is an interface to support dynamic dispatch.
type IArgumentsPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Expression() IExpressionContext

	// IsArgumentsPrimeContext differentiates from other interfaces.
	IsArgumentsPrimeContext()
}

type ArgumentsPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsPrimeContext() *ArgumentsPrimeContext {
	var p = new(ArgumentsPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_argumentsPrime
	return p
}

func InitEmptyArgumentsPrimeContext(p *ArgumentsPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_argumentsPrime
}

func (*ArgumentsPrimeContext) IsArgumentsPrimeContext() {}

func NewArgumentsPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsPrimeContext {
	var p = new(ArgumentsPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_argumentsPrime

	return p
}

func (s *ArgumentsPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsPrimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ArgumentsPrimeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgumentsPrime(s)
	}
}

func (s *ArgumentsPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgumentsPrime(s)
	}
}

func (p *GoliteParser) ArgumentsPrime() (localctx IArgumentsPrimeContext) {
	localctx = NewArgumentsPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_argumentsPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(GoliteParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllLvaluePrime() []ILvaluePrimeContext
	LvaluePrime(i int) ILvaluePrimeContext

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *LvalueContext) AllLvaluePrime() []ILvaluePrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILvaluePrimeContext); ok {
			len++
		}
	}

	tst := make([]ILvaluePrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILvaluePrimeContext); ok {
			tst[i] = t.(ILvaluePrimeContext)
			i++
		}
	}

	return tst
}

func (s *LvalueContext) LvaluePrime(i int) ILvaluePrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvaluePrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvaluePrimeContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *GoliteParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_lvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(280)
			p.LvaluePrime()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvaluePrimeContext is an interface to support dynamic dispatch.
type ILvaluePrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsLvaluePrimeContext differentiates from other interfaces.
	IsLvaluePrimeContext()
}

type LvaluePrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvaluePrimeContext() *LvaluePrimeContext {
	var p = new(LvaluePrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvaluePrime
	return p
}

func InitEmptyLvaluePrimeContext(p *LvaluePrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvaluePrime
}

func (*LvaluePrimeContext) IsLvaluePrimeContext() {}

func NewLvaluePrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvaluePrimeContext {
	var p = new(LvaluePrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lvaluePrime

	return p
}

func (s *LvaluePrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *LvaluePrimeContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, 0)
}

func (s *LvaluePrimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *LvaluePrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvaluePrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvaluePrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLvaluePrime(s)
	}
}

func (s *LvaluePrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLvaluePrime(s)
	}
}

func (p *GoliteParser) LvaluePrime() (localctx ILvaluePrimeContext) {
	localctx = NewLvaluePrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_lvaluePrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(GoliteParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BoolTerm() IBoolTermContext
	AllExpressionPrime() []IExpressionPrimeContext
	ExpressionPrime(i int) IExpressionPrimeContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) BoolTerm() IBoolTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermContext)
}

func (s *ExpressionContext) AllExpressionPrime() []IExpressionPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionPrimeContext); ok {
			len++
		}
	}

	tst := make([]IExpressionPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionPrimeContext); ok {
			tst[i] = t.(IExpressionPrimeContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) ExpressionPrime(i int) IExpressionPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionPrimeContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.BoolTerm()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(290)
			p.ExpressionPrime()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionPrimeContext is an interface to support dynamic dispatch.
type IExpressionPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode
	BoolTerm() IBoolTermContext

	// IsExpressionPrimeContext differentiates from other interfaces.
	IsExpressionPrimeContext()
}

type ExpressionPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionPrimeContext() *ExpressionPrimeContext {
	var p = new(ExpressionPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expressionPrime
	return p
}

func InitEmptyExpressionPrimeContext(p *ExpressionPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expressionPrime
}

func (*ExpressionPrimeContext) IsExpressionPrimeContext() {}

func NewExpressionPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionPrimeContext {
	var p = new(ExpressionPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expressionPrime

	return p
}

func (s *ExpressionPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionPrimeContext) OR() antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, 0)
}

func (s *ExpressionPrimeContext) BoolTerm() IBoolTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermContext)
}

func (s *ExpressionPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpressionPrime(s)
	}
}

func (s *ExpressionPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpressionPrime(s)
	}
}

func (p *GoliteParser) ExpressionPrime() (localctx IExpressionPrimeContext) {
	localctx = NewExpressionPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_expressionPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(GoliteParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.BoolTerm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolTermContext is an interface to support dynamic dispatch.
type IBoolTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTerm() IEqualTermContext
	AllBooltermPrime() []IBooltermPrimeContext
	BooltermPrime(i int) IBooltermPrimeContext

	// IsBoolTermContext differentiates from other interfaces.
	IsBoolTermContext()
}

type BoolTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTermContext() *BoolTermContext {
	var p = new(BoolTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTerm
	return p
}

func InitEmptyBoolTermContext(p *BoolTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTerm
}

func (*BoolTermContext) IsBoolTermContext() {}

func NewBoolTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTermContext {
	var p = new(BoolTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolTerm

	return p
}

func (s *BoolTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTermContext) EqualTerm() IEqualTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermContext)
}

func (s *BoolTermContext) AllBooltermPrime() []IBooltermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBooltermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IBooltermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBooltermPrimeContext); ok {
			tst[i] = t.(IBooltermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *BoolTermContext) BooltermPrime(i int) IBooltermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooltermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooltermPrimeContext)
}

func (s *BoolTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolTerm(s)
	}
}

func (s *BoolTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolTerm(s)
	}
}

func (p *GoliteParser) BoolTerm() (localctx IBoolTermContext) {
	localctx = NewBoolTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_boolTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.EqualTerm()
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserAND {
		{
			p.SetState(300)
			p.BooltermPrime()
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooltermPrimeContext is an interface to support dynamic dispatch.
type IBooltermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	EqualTerm() IEqualTermContext

	// IsBooltermPrimeContext differentiates from other interfaces.
	IsBooltermPrimeContext()
}

type BooltermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooltermPrimeContext() *BooltermPrimeContext {
	var p = new(BooltermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_booltermPrime
	return p
}

func InitEmptyBooltermPrimeContext(p *BooltermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_booltermPrime
}

func (*BooltermPrimeContext) IsBooltermPrimeContext() {}

func NewBooltermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooltermPrimeContext {
	var p = new(BooltermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_booltermPrime

	return p
}

func (s *BooltermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *BooltermPrimeContext) AND() antlr.TerminalNode {
	return s.GetToken(GoliteParserAND, 0)
}

func (s *BooltermPrimeContext) EqualTerm() IEqualTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermContext)
}

func (s *BooltermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooltermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooltermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBooltermPrime(s)
	}
}

func (s *BooltermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBooltermPrime(s)
	}
}

func (p *GoliteParser) BooltermPrime() (localctx IBooltermPrimeContext) {
	localctx = NewBooltermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_booltermPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(GoliteParserAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.EqualTerm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualTermContext is an interface to support dynamic dispatch.
type IEqualTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationTerm() IRelationTermContext
	AllEqualTermPrime() []IEqualTermPrimeContext
	EqualTermPrime(i int) IEqualTermPrimeContext

	// IsEqualTermContext differentiates from other interfaces.
	IsEqualTermContext()
}

type EqualTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualTermContext() *EqualTermContext {
	var p = new(EqualTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTerm
	return p
}

func InitEmptyEqualTermContext(p *EqualTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTerm
}

func (*EqualTermContext) IsEqualTermContext() {}

func NewEqualTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermContext {
	var p = new(EqualTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTerm

	return p
}

func (s *EqualTermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermContext) RelationTerm() IRelationTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationTermContext)
}

func (s *EqualTermContext) AllEqualTermPrime() []IEqualTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IEqualTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualTermPrimeContext); ok {
			tst[i] = t.(IEqualTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *EqualTermContext) EqualTermPrime(i int) IEqualTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermPrimeContext)
}

func (s *EqualTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTerm(s)
	}
}

func (s *EqualTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTerm(s)
	}
}

func (p *GoliteParser) EqualTerm() (localctx IEqualTermContext) {
	localctx = NewEqualTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoliteParserRULE_equalTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.RelationTerm()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserEQ || _la == GoliteParserNEQ {
		{
			p.SetState(310)
			p.EqualTermPrime()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualTermPrimeContext is an interface to support dynamic dispatch.
type IEqualTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRelt returns the relt rule contexts.
	GetRelt() IRelationTermContext

	// SetRelt sets the relt rule contexts.
	SetRelt(IRelationTermContext)

	// Getter signatures
	RelationTerm() IRelationTermContext
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode

	// IsEqualTermPrimeContext differentiates from other interfaces.
	IsEqualTermPrimeContext()
}

type EqualTermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	relt   IRelationTermContext
}

func NewEmptyEqualTermPrimeContext() *EqualTermPrimeContext {
	var p = new(EqualTermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTermPrime
	return p
}

func InitEmptyEqualTermPrimeContext(p *EqualTermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTermPrime
}

func (*EqualTermPrimeContext) IsEqualTermPrimeContext() {}

func NewEqualTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermPrimeContext {
	var p = new(EqualTermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTermPrime

	return p
}

func (s *EqualTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *EqualTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualTermPrimeContext) GetRelt() IRelationTermContext { return s.relt }

func (s *EqualTermPrimeContext) SetRelt(v IRelationTermContext) { s.relt = v }

func (s *EqualTermPrimeContext) RelationTerm() IRelationTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationTermContext)
}

func (s *EqualTermPrimeContext) EQ() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQ, 0)
}

func (s *EqualTermPrimeContext) NEQ() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEQ, 0)
}

func (s *EqualTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTermPrime(s)
	}
}

func (s *EqualTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTermPrime(s)
	}
}

func (p *GoliteParser) EqualTermPrime() (localctx IEqualTermPrimeContext) {
	localctx = NewEqualTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoliteParserRULE_equalTermPrime)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*EqualTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserEQ || _la == GoliteParserNEQ) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*EqualTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(317)

		var _x = p.RelationTerm()

		localctx.(*EqualTermPrimeContext).relt = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationTermContext is an interface to support dynamic dispatch.
type IRelationTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleTerm() ISimpleTermContext
	AllRelationTermPrime() []IRelationTermPrimeContext
	RelationTermPrime(i int) IRelationTermPrimeContext

	// IsRelationTermContext differentiates from other interfaces.
	IsRelationTermContext()
}

type RelationTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationTermContext() *RelationTermContext {
	var p = new(RelationTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTerm
	return p
}

func InitEmptyRelationTermContext(p *RelationTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTerm
}

func (*RelationTermContext) IsRelationTermContext() {}

func NewRelationTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationTermContext {
	var p = new(RelationTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationTerm

	return p
}

func (s *RelationTermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationTermContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *RelationTermContext) AllRelationTermPrime() []IRelationTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]IRelationTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationTermPrimeContext); ok {
			tst[i] = t.(IRelationTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *RelationTermContext) RelationTermPrime(i int) IRelationTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationTermPrimeContext)
}

func (s *RelationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationTerm(s)
	}
}

func (s *RelationTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationTerm(s)
	}
}

func (p *GoliteParser) RelationTerm() (localctx IRelationTermContext) {
	localctx = NewRelationTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoliteParserRULE_relationTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.SimpleTerm()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32212254720) != 0 {
		{
			p.SetState(320)
			p.RelationTermPrime()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationTermPrimeContext is an interface to support dynamic dispatch.
type IRelationTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetSmpt returns the smpt rule contexts.
	GetSmpt() ISimpleTermContext

	// SetSmpt sets the smpt rule contexts.
	SetSmpt(ISimpleTermContext)

	// Getter signatures
	SimpleTerm() ISimpleTermContext
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode

	// IsRelationTermPrimeContext differentiates from other interfaces.
	IsRelationTermPrimeContext()
}

type RelationTermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	smpt   ISimpleTermContext
}

func NewEmptyRelationTermPrimeContext() *RelationTermPrimeContext {
	var p = new(RelationTermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTermPrime
	return p
}

func InitEmptyRelationTermPrimeContext(p *RelationTermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTermPrime
}

func (*RelationTermPrimeContext) IsRelationTermPrimeContext() {}

func NewRelationTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationTermPrimeContext {
	var p = new(RelationTermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationTermPrime

	return p
}

func (s *RelationTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *RelationTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationTermPrimeContext) GetSmpt() ISimpleTermContext { return s.smpt }

func (s *RelationTermPrimeContext) SetSmpt(v ISimpleTermContext) { s.smpt = v }

func (s *RelationTermPrimeContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *RelationTermPrimeContext) GT() antlr.TerminalNode {
	return s.GetToken(GoliteParserGT, 0)
}

func (s *RelationTermPrimeContext) LT() antlr.TerminalNode {
	return s.GetToken(GoliteParserLT, 0)
}

func (s *RelationTermPrimeContext) LTE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLTE, 0)
}

func (s *RelationTermPrimeContext) GTE() antlr.TerminalNode {
	return s.GetToken(GoliteParserGTE, 0)
}

func (s *RelationTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationTermPrime(s)
	}
}

func (s *RelationTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationTermPrime(s)
	}
}

func (p *GoliteParser) RelationTermPrime() (localctx IRelationTermPrimeContext) {
	localctx = NewRelationTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoliteParserRULE_relationTermPrime)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*RelationTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32212254720) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*RelationTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(327)

		var _x = p.SimpleTerm()

		localctx.(*RelationTermPrimeContext).smpt = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleTermContext is an interface to support dynamic dispatch.
type ISimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	AllSimpleTermPrime() []ISimpleTermPrimeContext
	SimpleTermPrime(i int) ISimpleTermPrimeContext

	// IsSimpleTermContext differentiates from other interfaces.
	IsSimpleTermContext()
}

type SimpleTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTermContext() *SimpleTermContext {
	var p = new(SimpleTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTerm
	return p
}

func InitEmptySimpleTermContext(p *SimpleTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTerm
}

func (*SimpleTermContext) IsSimpleTermContext() {}

func NewSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermContext {
	var p = new(SimpleTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTerm

	return p
}

func (s *SimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleTermContext) AllSimpleTermPrime() []ISimpleTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ISimpleTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleTermPrimeContext); ok {
			tst[i] = t.(ISimpleTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *SimpleTermContext) SimpleTermPrime(i int) ISimpleTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermPrimeContext)
}

func (s *SimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTerm(s)
	}
}

func (s *SimpleTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTerm(s)
	}
}

func (p *GoliteParser) SimpleTerm() (localctx ISimpleTermContext) {
	localctx = NewSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoliteParserRULE_simpleTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Term()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserMINUS || _la == GoliteParserPLUS {
		{
			p.SetState(330)
			p.SimpleTermPrime()
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleTermPrimeContext is an interface to support dynamic dispatch.
type ISimpleTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Term() ITermContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsSimpleTermPrimeContext differentiates from other interfaces.
	IsSimpleTermPrimeContext()
}

type SimpleTermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptySimpleTermPrimeContext() *SimpleTermPrimeContext {
	var p = new(SimpleTermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTermPrime
	return p
}

func InitEmptySimpleTermPrimeContext(p *SimpleTermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTermPrime
}

func (*SimpleTermPrimeContext) IsSimpleTermPrimeContext() {}

func NewSimpleTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermPrimeContext {
	var p = new(SimpleTermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTermPrime

	return p
}

func (s *SimpleTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *SimpleTermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *SimpleTermPrimeContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleTermPrimeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, 0)
}

func (s *SimpleTermPrimeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *SimpleTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTermPrime(s)
	}
}

func (s *SimpleTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTermPrime(s)
	}
}

func (p *GoliteParser) SimpleTermPrime() (localctx ISimpleTermPrimeContext) {
	localctx = NewSimpleTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoliteParserRULE_simpleTermPrime)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*SimpleTermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserMINUS || _la == GoliteParserPLUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*SimpleTermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(337)
		p.Term()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryTerm() IUnaryTermContext
	AllTermPrime() []ITermPrimeContext
	TermPrime(i int) ITermPrimeContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *TermContext) AllTermPrime() []ITermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ITermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermPrimeContext); ok {
			tst[i] = t.(ITermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) TermPrime(i int) ITermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermPrimeContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoliteParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.UnaryTerm()
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFSLASH || _la == GoliteParserASTERISK {
		{
			p.SetState(340)
			p.TermPrime()
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermPrimeContext is an interface to support dynamic dispatch.
type ITermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetUnat returns the unat rule contexts.
	GetUnat() IUnaryTermContext

	// SetUnat sets the unat rule contexts.
	SetUnat(IUnaryTermContext)

	// Getter signatures
	UnaryTerm() IUnaryTermContext
	ASTERISK() antlr.TerminalNode
	FSLASH() antlr.TerminalNode

	// IsTermPrimeContext differentiates from other interfaces.
	IsTermPrimeContext()
}

type TermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	unat   IUnaryTermContext
}

func NewEmptyTermPrimeContext() *TermPrimeContext {
	var p = new(TermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_termPrime
	return p
}

func InitEmptyTermPrimeContext(p *TermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_termPrime
}

func (*TermPrimeContext) IsTermPrimeContext() {}

func NewTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermPrimeContext {
	var p = new(TermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_termPrime

	return p
}

func (s *TermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TermPrimeContext) GetOp() antlr.Token { return s.op }

func (s *TermPrimeContext) SetOp(v antlr.Token) { s.op = v }

func (s *TermPrimeContext) GetUnat() IUnaryTermContext { return s.unat }

func (s *TermPrimeContext) SetUnat(v IUnaryTermContext) { s.unat = v }

func (s *TermPrimeContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *TermPrimeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TermPrimeContext) FSLASH() antlr.TerminalNode {
	return s.GetToken(GoliteParserFSLASH, 0)
}

func (s *TermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTermPrime(s)
	}
}

func (s *TermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTermPrime(s)
	}
}

func (p *GoliteParser) TermPrime() (localctx ITermPrimeContext) {
	localctx = NewTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoliteParserRULE_termPrime)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TermPrimeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserFSLASH || _la == GoliteParserASTERISK) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TermPrimeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(347)

		var _x = p.UnaryTerm()

		localctx.(*TermPrimeContext).unat = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryTermContext is an interface to support dynamic dispatch.
type IUnaryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	SelectorTerm() ISelectorTermContext
	MINUS() antlr.TerminalNode

	// IsUnaryTermContext differentiates from other interfaces.
	IsUnaryTermContext()
}

type UnaryTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryTermContext() *UnaryTermContext {
	var p = new(UnaryTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryTerm
	return p
}

func InitEmptyUnaryTermContext(p *UnaryTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryTerm
}

func (*UnaryTermContext) IsUnaryTermContext() {}

func NewUnaryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryTermContext {
	var p = new(UnaryTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryTerm

	return p
}

func (s *UnaryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryTermContext) NOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOT, 0)
}

func (s *UnaryTermContext) SelectorTerm() ISelectorTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTermContext)
}

func (s *UnaryTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *UnaryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryTerm(s)
	}
}

func (s *UnaryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryTerm(s)
	}
}

func (p *GoliteParser) UnaryTerm() (localctx IUnaryTermContext) {
	localctx = NewUnaryTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GoliteParserRULE_unaryTerm)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(GoliteParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.SelectorTerm()
		}

	case GoliteParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.Match(GoliteParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.SelectorTerm()
		}

	case GoliteParserNIL, GoliteParserTRUE, GoliteParserFALSE, GoliteParserNEW, GoliteParserLPAREN, GoliteParserID, GoliteParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(353)
			p.SelectorTerm()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorTermContext is an interface to support dynamic dispatch.
type ISelectorTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	AllSelectorTermPrime() []ISelectorTermPrimeContext
	SelectorTermPrime(i int) ISelectorTermPrimeContext

	// IsSelectorTermContext differentiates from other interfaces.
	IsSelectorTermContext()
}

type SelectorTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermContext() *SelectorTermContext {
	var p = new(SelectorTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTerm
	return p
}

func InitEmptySelectorTermContext(p *SelectorTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTerm
}

func (*SelectorTermContext) IsSelectorTermContext() {}

func NewSelectorTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermContext {
	var p = new(SelectorTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTerm

	return p
}

func (s *SelectorTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectorTermContext) AllSelectorTermPrime() []ISelectorTermPrimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorTermPrimeContext); ok {
			len++
		}
	}

	tst := make([]ISelectorTermPrimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorTermPrimeContext); ok {
			tst[i] = t.(ISelectorTermPrimeContext)
			i++
		}
	}

	return tst
}

func (s *SelectorTermContext) SelectorTermPrime(i int) ISelectorTermPrimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermPrimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTermPrimeContext)
}

func (s *SelectorTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTerm(s)
	}
}

func (s *SelectorTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTerm(s)
	}
}

func (p *GoliteParser) SelectorTerm() (localctx ISelectorTermContext) {
	localctx = NewSelectorTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GoliteParserRULE_selectorTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Factor()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(357)
			p.SelectorTermPrime()
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorTermPrimeContext is an interface to support dynamic dispatch.
type ISelectorTermPrimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsSelectorTermPrimeContext differentiates from other interfaces.
	IsSelectorTermPrimeContext()
}

type SelectorTermPrimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermPrimeContext() *SelectorTermPrimeContext {
	var p = new(SelectorTermPrimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTermPrime
	return p
}

func InitEmptySelectorTermPrimeContext(p *SelectorTermPrimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTermPrime
}

func (*SelectorTermPrimeContext) IsSelectorTermPrimeContext() {}

func NewSelectorTermPrimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermPrimeContext {
	var p = new(SelectorTermPrimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTermPrime

	return p
}

func (s *SelectorTermPrimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermPrimeContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, 0)
}

func (s *SelectorTermPrimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *SelectorTermPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermPrimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTermPrime(s)
	}
}

func (s *SelectorTermPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTermPrime(s)
	}
}

func (p *GoliteParser) SelectorTermPrime() (localctx ISelectorTermPrimeContext) {
	localctx = NewSelectorTermPrimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GoliteParserRULE_selectorTermPrime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(GoliteParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FactorPrime1() IFactorPrime1Context
	FactorPrime2() IFactorPrime2Context
	NUMBER() antlr.TerminalNode
	NEW() antlr.TerminalNode
	ID() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NIL() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) FactorPrime1() IFactorPrime1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorPrime1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorPrime1Context)
}

func (s *FactorContext) FactorPrime2() IFactorPrime2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorPrime2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorPrime2Context)
}

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GoliteParserNUMBER, 0)
}

func (s *FactorContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserFALSE, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GoliteParserRULE_factor)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.FactorPrime1()
		}

	case GoliteParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.FactorPrime2()
		}

	case GoliteParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.Match(GoliteParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserNEW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(369)
			p.Match(GoliteParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(GoliteParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(371)
			p.Match(GoliteParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(372)
			p.Match(GoliteParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(373)
			p.Match(GoliteParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorPrime1Context is an interface to support dynamic dispatch.
type IFactorPrime1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFactorPrime1Context differentiates from other interfaces.
	IsFactorPrime1Context()
}

type FactorPrime1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorPrime1Context() *FactorPrime1Context {
	var p = new(FactorPrime1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factorPrime1
	return p
}

func InitEmptyFactorPrime1Context(p *FactorPrime1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factorPrime1
}

func (*FactorPrime1Context) IsFactorPrime1Context() {}

func NewFactorPrime1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorPrime1Context {
	var p = new(FactorPrime1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factorPrime1

	return p
}

func (s *FactorPrime1Context) GetParser() antlr.Parser { return s.parser }

func (s *FactorPrime1Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *FactorPrime1Context) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorPrime1Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *FactorPrime1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorPrime1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorPrime1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactorPrime1(s)
	}
}

func (s *FactorPrime1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactorPrime1(s)
	}
}

func (p *GoliteParser) FactorPrime1() (localctx IFactorPrime1Context) {
	localctx = NewFactorPrime1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GoliteParserRULE_factorPrime1)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Expression()
	}
	{
		p.SetState(378)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorPrime2Context is an interface to support dynamic dispatch.
type IFactorPrime2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Arguments() IArgumentsContext

	// IsFactorPrime2Context differentiates from other interfaces.
	IsFactorPrime2Context()
}

type FactorPrime2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorPrime2Context() *FactorPrime2Context {
	var p = new(FactorPrime2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factorPrime2
	return p
}

func InitEmptyFactorPrime2Context(p *FactorPrime2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factorPrime2
}

func (*FactorPrime2Context) IsFactorPrime2Context() {}

func NewFactorPrime2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorPrime2Context {
	var p = new(FactorPrime2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factorPrime2

	return p
}

func (s *FactorPrime2Context) GetParser() antlr.Parser { return s.parser }

func (s *FactorPrime2Context) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FactorPrime2Context) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FactorPrime2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorPrime2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorPrime2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactorPrime2(s)
	}
}

func (s *FactorPrime2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactorPrime2(s)
	}
}

func (p *GoliteParser) FactorPrime2() (localctx IFactorPrime2Context) {
	localctx = NewFactorPrime2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GoliteParserRULE_factorPrime2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(GoliteParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserLPAREN {
		{
			p.SetState(381)
			p.Arguments()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
