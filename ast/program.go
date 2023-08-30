package ast

import (
	"bytes"
	"fmt"
	// "fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// // Program node is interface for ...
// type ProgramNode interface {
// 	String() string
// }


type Program struct {
	token 		 *token.Token
	Types 		 []*TypeDecl
	Declarations []*Declaration
	Functions 	 []*Function
	UserTypes 	 map[string]types.StructType
	SymbolTables *st.SymbolTables
}

func NewProgram(token *token.Token, types []*TypeDecl, declarations []*Declaration, functions []*Function) *Program{
	return &Program{token, types, declarations, functions, nil, nil}
}


func (program *Program) String() string{
	// Program = Types Declarations Functions 'eof'  
	// obs: dont need to print 'eof'

	var out bytes.Buffer
	
	for _, typeDecl := range program.Types {
		out.WriteString(typeDecl.String())
	}

	for _, declaration := range program.Declarations {
		out.WriteString(declaration.String())
	}

	for _, function := range program.Functions {
		out.WriteString(function.String())
	}

	return out.String()
}

func (p *Program) BuildSymbolTable() []string{
	errors := make([]string, 0)

	p.SymbolTables = st.NewSymbolTables()


	// build `Structs` symbol table
	for _, typeDecl := range p.Types {
		errors = typeDecl.BuildSymbolTable(errors, p.SymbolTables, p.UserTypes)
	}

	// build `Globals` variables symbol table
	for _, declaration := range p.Declarations {
		errors = declaration.BuildSymbolTable(errors, p.SymbolTables, p.UserTypes)
	}

	// build `Functions` symbol table
	for _, function := range p.Functions {
		errors = function.BuildSymbolTable(errors, p.SymbolTables, p.UserTypes)
	}

	return errors
}


func (p *Program) TypeCheck(errors []string) []string{
	for _, typeDecl := range p.Types {
		errors = typeDecl.TypeCheck(errors, p.SymbolTables)
	}

	for _, declaration := range p.Declarations {
		errors = declaration.TypeCheck(errors, p.SymbolTables)
	}
	
	for _, function := range p.Functions {
		errors = function.TypeCheck(errors, p.SymbolTables)
	}

	// check if contains main function, if it has no return type and no parameters
	if mainEntry, contains := p.SymbolTables.Funcs.Contains("main"); !contains{
		errors = append(errors, SemErr(p.token, "No `main` function."))
	} else {
		if mainEntry.RetTy != types.VoidTySig {
			errors = append(errors, SemErr(p.token, "Main function must have no return type."))
		}

		if len(mainEntry.Parameters) != 0 {
			errors = append(errors, SemErr(p.token, "Main function must have no parameters."))
		}
	}
	return errors
}

func (p *Program) ToLLVM() []*ir.BasicBlock{
	basicBlocks := []*ir.BasicBlock{}

	for _, tyDecl := range p.Types {
		tyDecl.ToLLVM(p.SymbolTables)
	}

	for _, declaration := range p.Declarations {
		declaration.ToLLVM(p.SymbolTables)
	}

	for _, function := range p.Functions {
		basicBlocks = append(basicBlocks, function.ToLLVM(p.SymbolTables))
	}

	return basicBlocks
}


func (p *Program) PrintLLVM(targetTriple, filename string) string {

	var out bytes.Buffer
	out.WriteString("source_filename = \"" + filename + "\"\n")
	out.WriteString("target triple = \"" + targetTriple + "\"\n\n")

	for _, tyDecl := range p.Types {
		out.WriteString(tyDecl.PrintLLVM(p.SymbolTables))
	}
	out.WriteString("\n")

	for _, decl := range p.Declarations {
		out.WriteString(decl.PrintLLVM(p.SymbolTables))
	}
	out.WriteString("\n")

	for _, f := range p.Functions {
		out.WriteString(f.PrintLLVM(p.SymbolTables))
		out.WriteString("\n")
	}

	out.WriteString(ir.EOFIns.String())

	return out.String()
}

func (p *Program) TranslateToAssembly() string {
	var out bytes.Buffer

	out.WriteString("\t .arch armv8-a\n")
	
	for _, global := range p.SymbolTables.Globals.GetEntries() {
		out.WriteString(fmt.Sprintf("\t .comm g%s, 8, 8\n", global.Name))
	}

	out.WriteString("\t .text\n")

	for _, f := range p.Functions {
		out.WriteString(f.TranslateToAssembly(p.SymbolTables))
		out.WriteString("\n")
	}

	out.WriteString(ir.EOFIns.ARMString())
	return out.String()
}
