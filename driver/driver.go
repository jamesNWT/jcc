package driver

import (
	"os/exec"
	"fmt"
	"log"
	"flag"
)

func Driver() {
	lexFlagHelp := "compilation will run the lexer, but stop before parsing"
	parseFlagHelp := "compilation will run the lexer and parser, but stop before assembly generation"
	codeGenFlagHelp := "compilation will run the lexer, parser, and assembly generation, but stop before code emission"

	lexFlag := flag.Bool("lex", false, lexFlagHelp)
	parseFlag := flag.Bool("parse", false, parseFlagHelp)
	codeGenFlag := flag.Bool("codegen", false, codeGenFlagHelp)

	flag.Parse()
	
	// Run the preprocessor
	sourceFileName := flag.Arg(0)
	preprocFileName := sourceFileName[:len(sourceFileName)-1] + "i"
	commandText := fmt.Sprintf("gcc -E -P $s -o $s", sourceFileName, preprocFileName)

	cmd := exec.Command("bash", "-c", commandText)

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	// Compile the preprocessed soure file to assembly.
	assemblyFileName := sourceFileName[:len(sourceFileName)-1] + "s"
	if *lexFlag {
		// stubbed
	}
	if *parseFlag {
		// stubbed
	}
	if *codeGenFlag {
		// stubbed 
	}
	// DUMMY
	commandText = fmt.Sprintf("touch $s", assemblyFileName)
	cmd = exec.Command("bash", "-c", commandText)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Assemble and link the assembly file to produce an executable
	outputFileName := sourceFileName[:len(sourceFileName)-2]

	commandText = fmt.Sprintf("gcc $s -o $s", assemblyFileName, outputFileName)

	cmd = exec.Command("bash", "-c", commandText)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

}
