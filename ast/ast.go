package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

//==============================================================================
// Metadata
//==============================================================================

// Statement nodes
// block, assignment, print, delete, conditional, loop, ret, invocation

// Expression nodes
// intlit, nillit, boollit, factorPrime1, factorPrime2, binary, unary, variable, field, new

// nodes with `BuildSymbleTable` method: declaration, function and typedecl

// nodes with `TypeCheck` method:

// nodes with `GetType` method:

// nodes with `HasReturn` method: statement nodes, function

//==============================================================================
// usage for all nodes:
//==============================================================================
// REVIEW use this to add padding to functions, typedecls, blocks
const padding = 2

// SemErr returns a string containing a semantic error message.
func SemErr(token *token.Token, message string) string {
	var out bytes.Buffer

	out.WriteString("Semantic error at ")
	out.WriteString(token.String())
	out.WriteString(": ")
	out.WriteString(message)
	return out.String()
}

// Get a variable entry from parameters, local or global variables
func GetVarEntry(tables *st.SymbolTables, funCallee *st.FuncEntry, id string) *st.VarEntry{
	var varEntry *st.VarEntry

	// check if is parameter
	if entry, contains := funCallee.ContainsParam(id); contains {
		varEntry = entry
	} else if entry, contains := funCallee.Variables.Contains(id); contains {
		varEntry = entry
	}
	return varEntry

}