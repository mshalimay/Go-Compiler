package types

import "strings"

// Type includes information about working on Types
type Type interface {
	String() string
}

type BaseType interface {
	Type
}

type StructType interface {
	Type
	GetName() string
}

// ==============================================================================
// Base Types
// ==============================================================================
// 'int' type
type IntTy struct{}

func (intTy *IntTy) String() string {
	return "int"
}

// 'bool' type
type BoolTy struct{}

func (boolTy *BoolTy) String() string {
	return "bool"
}

// 'void' type
type VoidTy struct{}

func (voidTy *VoidTy) String() string {
	return "void"
}

type NilTy struct{}

func (nilTy *NilTy) String() string { return "nil" }

// 'unknown' type
type UnknownTy struct{}

func (unkTy *UnknownTy) String() string { return "unknown" }

// 'undeclared' type
type UndeclaredTy struct{}

func (undecTy *UndeclaredTy) String() string { return "undeclared" }

//==============================================================================
// Struct Types
//==============================================================================

type StrucType struct {
	Name string
}

func (sTy *StrucType) String() string {
	return sTy.Name
}

func (sTy *StrucType) GetName() string {
	return sTy.Name
}

//==============================================================================
// Sigleton definitions
//==============================================================================
// Each type will be associated to a singleton, so comparisons in semantic analysis
// are made with the addresses the singletons hold.

var IntTySig *IntTy
var BoolTySig *BoolTy
var NilTySig *NilTy
var VoidTySig *VoidTy
var UnknownTySig *UnknownTy
var StructSig *StrucType
var UndeclTySig *UndeclaredTy

// Set singletons for all types in `golite`. Called only once when the package is loaded.
func init() {
	IntTySig = &IntTy{}
	BoolTySig = &BoolTy{}
	NilTySig = &NilTy{}
	VoidTySig = &VoidTy{}
	UnknownTySig = &UnknownTy{}
	UndeclTySig = &UndeclaredTy{}
}

// ==============================================================================
// Helper functions
// ==============================================================================
func StrToTy(ty string, userTypes map[string]StructType) (Type, string) {

	// "int" -> IntTySig
	if ty == "int" {
		return IntTySig, ty

		// "bool" -> BoolTySig
	} else if ty == "bool" {
		return BoolTySig, ty

		// REVIEW this
	} else if ty == "" {
		return VoidTySig, ty

		// "nil" -> NilTySig
	} else if ty == "nil" {
		return NilTySig, ty

		// "*id" -> StrucType or UnknownTySig
	} else if strings.Contains(ty, "*") {
		ty = strings.Trim(ty, "*")
		if strucType, ok := userTypes[ty]; ok {
			return strucType, ty
		} else {
			return UnknownTySig, ty
		}
	} else {
		panic("Unrecognized conversion from string to type: " + ty)
	}
}

// Returns true if Type is an original type of the `golite` language
// (false if: a type created by the user or a unknown/undeclared type created here to handle errors)
func IsBaseType(ty Type) bool {
	if ty == IntTySig || ty == BoolTySig || ty == VoidTySig || ty == NilTySig {
		return true
	}
	return false
}

func IsStructType(ty Type) bool {
	_, ok := ty.(*StrucType)
	return ok
}

func ToLLVM(ty Type) string {
	switch ty.(type) {
	case *IntTy:
		return "i64"
	case *BoolTy:
		return "i64"
	case *VoidTy:
		return "void"
	case *NilTy:
		return "i64"
	case *StrucType:
		return "%" + "struct." + ty.String() + "*"
	default:
		panic("Invalid type for LLVM conversion: " + ty.String())
	}
}
