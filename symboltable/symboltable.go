package symboltable

import (
	"fmt"
	"golite/types"
)

type VarScope int

const (
	GLOBAL VarScope = iota
	LOCAL
)


//==============================================================================
// Symbol table entries
//==============================================================================

// SymbolTableEntry is an interface that all entries in the symbol table must implement
type SymbolTableEntry interface {
	fmt.Stringer
}

//------------------------------------------------
// VarEntry, FuncEntry, StructEntry structs
//------------------------------------------------

// VarEntry is the representation of a declared variable in the symbol table
type VarEntry struct {
	Name  string
	Ty    types.Type
	Scope VarScope

	// Used in IR:
	LLVMName string
	Reg      interface{} // can be a *ir.VirtualRegister or *ir.GlobalRegister	
	Initval  string
}

func NewVarEntry(name string, ty types.Type, scope VarScope) *VarEntry {
	return &VarEntry{name, ty, scope, "", nil, ""}
}

// FuncEntry is the representation of func definition in the symbol table
type FuncEntry struct {
	Name       		string
	RetTy      		types.Type
	Parameters 		[]*VarEntry             // separate parameters for easier type checking and code generation
	Variables  		*SymbolTable[*VarEntry] // define as a symbolTable so when child is looking up for variables, use the `Contains` method

	// Used in IR
	LLVMName  		string   // eg: @main, @foo
	RetValReg 		interface{} // can be a *ir.VirtualRegister or *ir.GlobalRegister


	// Used in assembly
	TempRegisters 		[]string
	StackRegisters 		map[string]string
}

func NewFuncEntry(name string, retTy types.Type) *FuncEntry {
	return &FuncEntry{name, retTy, make([]*VarEntry, 0), nil, "", nil, make([]string, 0), make(map[string]string)}
}

// StructEntry is the representation of type/struct definition in the symbol table
// consider eg: type Point struct {x int, y int}
type StructEntry struct {
	Ty		   types.StructType
	Fields 		map[string]*Field

	// fields used in IR
	LLVMName   		string 			// %struct.Point
	NilRegister 	interface{} 	// nil register for struct type
}

// Field struct is useful to get offset in LLVM translation while mantaining a hashmap for `Fields` in StructEntry
type Field struct {
	Ty 			types.Type
	Offset 		int
}


func NewStructEntry(ty types.StructType) *StructEntry {
	return &StructEntry{ty, make(map[string]*Field), "", nil}
}


//---------------------------------------------
// SymbolTableEntry interface implementations
//---------------------------------------------

func (entry *VarEntry) String() string { return "" }
func (entry *FuncEntry) String() string { return "" }
func (entry *StructEntry) String() string { return "" }

//==============================================================================
// SymbolTable structs and methods
//==============================================================================

// SymbolTable is a hashmap of string keys and generic type T that must implement SymbolTableEntry interface
// In our case, T =  VarEntry, FuncEntry, StructEntry
type SymbolTable[T SymbolTableEntry] struct {
	parent *SymbolTable[T]
	table  map[string]T		// map of entries in the symbol table (actual representation of the ST)
}

// wrapper for all symbol tables in a GoLite program
type SymbolTables struct {
	Globals *SymbolTable[*VarEntry]
	Funcs   *SymbolTable[*FuncEntry]
	Structs *SymbolTable[*StructEntry]
}

func NewSymbolTables() *SymbolTables {
	return &SymbolTables{
		Globals:  NewSymbolTable[*VarEntry](nil), 
		Funcs: 	  NewSymbolTable[*FuncEntry](nil), 
		Structs:  NewSymbolTable[*StructEntry](nil)}
}

// NewSymbolTable takes as argument a generic type T that implements the fmt.Stringer interface
// and returns a pointer to a SymbolTable[T] struct.
// note on syntax: `[T fmt.Stringer]` defines what `T` is about: a generic type that implements 
//
//	fmt.Stringer interface; `T` is then referenced in the parameters and return type of the function.
func NewSymbolTable[T SymbolTableEntry](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{parent, make(map[string]T)}
}

// Insert takes as argument a string `id` and a generic type `entry` that implements the SymbolTableEntry interface
// and inserts the pair (id, entry) into the symbol table.
func (st *SymbolTable[T]) Insert(id string, entry T) {
	st.table[id] = entry
}

// Contains takes as argument a string `id` and returns the entry associated 
// with `id` if it exists in the symbol table. Otherwise, it returns an empty entry.
func (st *SymbolTable[T]) Contains(id string) (T, bool) {
	
	// look for the entry in the given symbol table and all its parents
	// return the first entry found; if not found, return nil, false
	current := st
	for current != nil {
		if entry, ok := current.table[id]; ok {
			return entry, true
		}
		current = current.parent
	}
	var empty T
	return empty, false
}

// REVIEW: see if better design to look for parameters 
// have to consider that will need to typecheck each parameter with Expression in calls also
func (funcEntry *FuncEntry) ContainsParam(id string) (*VarEntry, bool) {
	for _, param := range funcEntry.Parameters {
		if param.Name == id {
			return param, true
		}
	}
	return nil, false
}

func (st *SymbolTable[T]) GetEntries() []T {
	entries := make([]T, 0)
	for _, entry := range st.table {
		entries = append(entries, entry)
	}
	return entries
}
