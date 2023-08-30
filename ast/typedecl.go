package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type TypeDecl struct {
	token   	*token.Token
	strucID 	string // eg: `myType` (obs: no asterisk)
	fields  	[]*Decl
}

func NewTypeDecl(token *token.Token, strucID string, fields []*Decl) *TypeDecl {
	return &TypeDecl{token, strucID, fields}
}

func (tyDecl *TypeDecl) String() string {
	// TypeDeclaration: 'type' 'id' 'struct' '{' Fields '}' ';'
	// eg: type myType struct { a int; b bool; c *myOtherType;}

	var out bytes.Buffer

	out.WriteString("type ")
	out.WriteString(tyDecl.strucID)
	out.WriteString(" struct")
	out.WriteString("{\n")

	for _, decl := range tyDecl.fields {
		out.WriteString(decl.id)
		out.WriteString(" ")
		out.WriteString(decl.ty)
		out.WriteString(";\n")
	}
	out.WriteString("}")
	out.WriteString(";\n")

	return out.String()
}

func (tyDecl *TypeDecl) BuildSymbolTable(errors []string, tables *st.SymbolTables, userTypes map[string]types.StructType) []string {

	// check if another type with the same name; if so, append error and return
	if _, contains := tables.Structs.Contains(tyDecl.strucID); contains {
		errors = append(errors, SemErr(tyDecl.token,
			fmt.Sprintf("duplicate type declaration `%s`.", tyDecl.strucID)))
		return errors
	}
	// REVIEW: doing this we are not looking for nested errors in the typeDecl.

	// create a new struct entry to add to the symbol table
	structEntry := st.NewStructEntry(userTypes[tyDecl.strucID])

	// fill struct entry `Fields`. @memo: `Fields` = map[string]types.Type
	var id, ty string
	var offset int
	for _, decl := range tyDecl.fields {
		id = decl.id
		ty = decl.ty

		// if field is duplicate append error and go to next
		if _, ok := structEntry.Fields[id]; ok {
			errors = append(errors, SemErr(decl.GetToken(),
				fmt.Sprintf("duplicate field declaration `%s` in `%s`.", id, tyDecl.strucID)))
			continue
		}

		// get the field type;
		// @memo: StrToTy return UnknownTySig if type does not exist (i.e., not base type or not declared)
		fieldType, fieldTyName := types.StrToTy(ty, userTypes)

		// if type for field does not exist, append error
		if fieldType == types.UnknownTySig {
			errors = append(errors, SemErr(tyDecl.token,
				fmt.Sprintf("unknown type `%s` in struct `%s` for field `%s`", fieldTyName, tyDecl.strucID, id)))
		}
		// else, insert the field type in the struct entry
		// NOTE: this will add fields with unknown types to the struct fields
		structEntry.Fields[id] = &st.Field{Ty: fieldType, Offset: offset}
		offset ++
	}

	// insert the new type to `Structs` symbol table
	tables.Structs.Insert(tyDecl.strucID, structEntry)

	return errors
}

func (tyDecl *TypeDecl) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	// eg: type myType struct { a int; b bool; c *SomeType; d *myType}

	// check for functions or variables with same name as this type
	if _, contains := tables.Globals.Contains(tyDecl.strucID); contains {
		errors = append(errors, SemErr(tyDecl.token, "struct and variable name conflict for `"+tyDecl.strucID+"`"))
	}

	if _, contains := tables.Funcs.Contains(tyDecl.strucID); contains {
		errors = append(errors, SemErr(tyDecl.token, "struct and function name conflict for `"+tyDecl.strucID+"`"))
	}
	return errors
}

func (tydecl *TypeDecl) ToLLVM(tables *st.SymbolTables) {
	// consider eg: type Node struct {x int, next *Node}

	// get the struct entry
	strucEntry, _ := tables.Structs.Contains(tydecl.strucID)

	// save llvm struc name into symbol table. Eg: %struct.Node
	// REVIEW IF NEED THIS
	strucEntry.LLVMName  = "%" + "struct." + tydecl.strucID

	// create register for nil value and store in symbol table
	strucEntry.NilRegister = ir.NewGlobalRegister(strucEntry.Ty, ".nil" + tydecl.strucID)

	
}

func (tydecl *TypeDecl) PrintLLVM(tables *st.SymbolTables) string {
	// consider eg: type Node struct {x int, next *Node}

	// get struct entry
	strucEntry, _ := tables.Structs.Contains(tydecl.strucID)

	var out bytes.Buffer	
	out.WriteString(fmt.Sprintf("%s = type {", strucEntry.LLVMName))
	
	// i64, %struct.Node*, ...
	var f *st.Field
	for i, decl := range tydecl.fields {
		f = strucEntry.Fields[decl.id]
		out.WriteString(types.ToLLVM(f.Ty))
		if i != len(tydecl.fields) - 1{
			out.WriteString(", ")
		}
	}
	out.WriteString("}\n")
	// out: %struct.Node = type {i64, %struct.Node*}


	// global variable for nil representation of this type
	nilReg := strucEntry.NilRegister.(ir.Operand)
	out.WriteString(fmt.Sprintf("%s = common global %s null\n", nilReg.String(), nilReg.LLVMTy()))
	// out:
	// %struct.Node = type {i64, %struct.Node*}
	// @.nilNode = common global %struct.Node* null

	return out.String()
}
