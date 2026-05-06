package driver

import (
	"flag"
	//"fmt"
	"log"
	"os/exec"
)

func Driver() {
	lexFlagHelp := "compilation will run the lexer, but stop before parsing"
	parseFlagHelp := "compilation will run the lexer and parser, but stop before assembly generation"
	codeGenFlagHelp := "compilation will run the lexer, parser, and assembly generation, but stop before code emission"

	lexFlag := flag.Bool("lex", false, lexFlagHelp)
	parseFlag := flag.Bool("parse", false, parseFlagHelp)
	codeGenFlag := flag.Bool("codegen", false, codeGenFlagHelp)

	flag.Parse()
	 
	// intermediate file names
	sourceFileName := flag.Arg(0)
	sourceBaseName := sourceFileName[:len(sourceFileName)-1]
	preprocessedFileName := addExtension(sourceBaseName, "i")
	assemblyFileName := addExtension(sourceBaseName, "s")

	// Run the preprocessor
	cmd := exec.Command("gcc", "-E", "-P", sourceFileName, "-o", preprocessedFileName)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("Preprocessor error:", err)
	}

	// Compile the preprocessed soure file to assembly.
	if *lexFlag {
		// stubbed
	}
	if *parseFlag {
		// stubbed
	}
	if *codeGenFlag {
		// stubbed
	}
	// DUMMY - for now, to test that the driver is working, just invoke gcc
	// to generate the assembly file
	cmd = exec.Command("gcc", "-S", "-O", "-fno-asynchronous-unwind-tables", "-fcf-protection=none", sourceFileName)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("Compilation error:", err)
	}

	// Assemble and link the assembly file to produce an executable
	cmd = exec.Command("gcc", assemblyFileName, "-o", sourceBaseName)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("Linking error", err)
	}

}

func addExtension(baseName, extension string) string {
	return baseName + "." + extension
}
