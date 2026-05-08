package driver

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Driver() {
	// Define flags
	lexFlagHelp := "compilation will run the lexer, but stop before parsing"
	parseFlagHelp := "compilation will run the lexer and parser, but stop before assembly generation"
	codeGenFlagHelp := "compilation will run the lexer, parser, and assembly generation, but stop before code emission"

	lexFlag := flag.Bool("lex", false, lexFlagHelp)
	parseFlag := flag.Bool("parse", false, parseFlagHelp)
	codeGenFlag := flag.Bool("codegen", false, codeGenFlagHelp)

	// resolve flags
	flag.Parse()

	// check flags are valid
	stopFlags := 0
	if *lexFlag {
		stopFlags++
	}
	if *parseFlag {
		stopFlags++
	}
	if *codeGenFlag {
		stopFlags++
	}
	if stopFlags > 1 {
		log.Fatal("Only one of -lex, -parse, -codegen may be specified")
	}

	// check our argument is valid
	if flag.NArg() == 0 {
		log.Fatal("no source file was specified")
	}

	sourceFilePath := flag.Arg(0)

	_, err := os.Stat(sourceFilePath)

	if err != nil {
		log.Fatal("Argument error: ", err)
	}

	// intermediate file names
	preprocessedFilePath := changeExtension(sourceFilePath, "i")
	assemblyFilePath := changeExtension(sourceFilePath, "s")
	outputFilePath := sourceFilePath[:len(assemblyFilePath)-2]

	// Run the preprocessor
	cmd := exec.Command("gcc", "-E", "-P", sourceFilePath, "-o", preprocessedFilePath)
	_, err = cmd.Output()
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
	cmd = exec.Command("gcc", "-S", "-O", "-fno-asynchronous-unwind-tables", "-fcf-protection=none", sourceFilePath, "-o", assemblyFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("Compilation error:", err)
	}

	// Assemble and link the assembly file to produce an executable
	cmd = exec.Command("gcc", assemblyFilePath, "-o", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("Linking error: ", err)
	}

}

func changeExtension(fileName, newExtension string) string {
	// TODO: this is brittle, should be more robust
	parts := strings.Split(fileName, ".")

	return parts[0] + "." + newExtension
}
