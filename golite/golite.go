package main

import (
	"flag"
	"fmt"
	"golite/lexer"
	"golite/parser"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	lexFlag      	bool
	astFlag      	bool
	llvmFlag     	bool
	sFlag      	 	bool
	eFlag       	bool
	outputFilename 	*string
	targetTriple = "x86_64-linux-gnu"
)

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n")
	fmt.Fprint(os.Stderr, "  golite.go [-lex] [-ast] [-llvm] [-e] [-o output_filename] path_to_file [target_triple]\n")
	fmt.Fprintln(os.Stderr, "\nArguments:")
	fmt.Fprintln(os.Stderr, "  path_to_file \n\tfull path to the golite program")
	fmt.Fprintln(os.Stderr, "\nOptions:")
	flag.PrintDefaults()
}


func init() {
	flag.BoolVar(&lexFlag, "lex", false, "Use this flag to print only the program tokens.")
	flag.BoolVar(&astFlag, "ast", false, "Use this flag to print the AST.")
	flag.BoolVar(&llvmFlag, "llvm", false, "Use this flag to print the LLVM IR. Optionally specify a target triple as second argument of function")
	flag.BoolVar(&sFlag, "s", false, "Use this flag to print the assembly code.")
	flag.BoolVar(&eFlag, "e", false, "Use this flag to compile and an executable. Optionally, specify the name of the using the o flag. Defaults to a.out if not provided")
	outputFilename = flag.String("o", "", "Specify the name of the executable. ")
}

func main() {
	//---------------------
	// input logic
	// --------------------
	// Add custom usage message
	flag.Usage = usage

	// Parse command line arguments
    flag.Parse()
	args := flag.Args()

	// // // NOTE: for debugging
	// llvmFlag = true
	// args = append(args, "/home/mshalimay/Compilers/proj-mashalimay/benchmarks/asm/simple1.golite")
	// asmFlag = true

	// Check correct number of arguments
    if len(args) < 1 {
        fmt.Println("\nError: Missing path to the golite program file.")
        flag.Usage()
        os.Exit(1)
    }

	// check if file exists
	inputSourcePath := args[0]
	if _, err := os.Stat(inputSourcePath); os.IsNotExist(err) {
		fmt.Println("Error: golite program file not found. Please check the path provided.")
		os.Exit(1)
	}

	// get target triple, if any
	if llvmFlag && len(args) > 1 {
		targetTriple = args[1]
	}

	// get the directory and name of the input source path provided
	dir := filepath.Dir(inputSourcePath)
	filenameWithoutExt := filepath.Base(inputSourcePath)
	ext := filepath.Ext(filenameWithoutExt)
	filenameWithoutExt = filenameWithoutExt[:len(filenameWithoutExt)-len(ext)]

	//---------------------
	// lexing phase
	// --------------------

	// Create lexer and lex program
	lexer := lexer.NewLexer(inputSourcePath)

	// if lex flag is set, print all tokens and exit
	if lexFlag{
		lexer.FillTokenStream()
		lexer.PrintAllTokens()
		return
	}

	//-------------------------------------
	// parsing - sintactic analysis
	// ------------------------------------

	// Create parser
	parser := parser.NewParser(lexer)
	
	// Parse the program file. This will return nil if there is a parsing error. 
	// This also populates the `program` struct with user defined types.
	program := parser.Parse()

	// If ast flag and no parsing error, print the AST.
	if astFlag  && program != nil{
		fmt.Println(program.String())
	}

	//-----------------------------------
	// parsing - semantic analysis
	// ----------------------------------

	// if no sintactic error, do semantic analysis.
	var errors []string
	if program != nil {
		errors = program.BuildSymbolTable()
		errors = program.TypeCheck(errors)
		for _, err := range errors {
			fmt.Println(err)
		}
	}

	//-----------------------------------
	// IR translation
	// ----------------------------------
	if program != nil && len(errors) == 0 {
		program.ToLLVM()
		if llvmFlag {
			llvmString := program.PrintLLVM(targetTriple, filenameWithoutExt)
			
			// Create ll file
			llvmFilename := dir + "/" + filenameWithoutExt + ".ll"
			file, err := os.OpenFile(llvmFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error opening/creating .ll file:", err)
				return
			}
			defer file.Close()

			// Write the string to the file
			_, err = file.WriteString(llvmString)
			if err != nil {
				// Handle the error
				fmt.Println("Error writing to .ll file:", err)
				return
			}
		}
	}

	//-----------------------------------
	// Code generation
	// ----------------------------------
	if program != nil && len(errors) == 0 && sFlag {
		asmString := program.TranslateToAssembly()
		
		// Create .s file
		asmFilename := dir + "/" + filenameWithoutExt + ".s"
		file, err := os.OpenFile(asmFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening/creating .s file:", err)
			return
		}
		defer file.Close()

		// Write the string to the file
		_, err = file.WriteString(asmString)
		if err != nil {
			// Handle the error
			fmt.Println("Error writing to .s file:", err)
			return
		}

		if eFlag {
			// create .out file
			var outFilename string
			if *outputFilename != "" {
				outFilename = dir + "/" + *outputFilename + ".out"
			} else {
				outFilename = dir + "/" + "a.out"
			}

			cmd := exec.Command("gcc", "-o", outFilename, asmFilename)
			cmd.Run()
		}
	}
}
