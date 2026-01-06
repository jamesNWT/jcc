package main

import (
	"os/exec"
	"fmt"
	"log"
	"flag"
)

func main() {
	lexFlagHelp := "compilation will run the lexer, but stop before parsing"
	parseFlagHelp := "compilation will run the lexer and parser, but stop before assembly generation"
	codeGenFlagHelp := "compilation will run the lexer, parser, and assembly generation, but stop before code emission"

	lexFlag := flag.Bool("lex", false, lexFlagHelp)
	parseFlag := flag.Bool("parse", false, parseFlagHelp)
	codegenFlag := flag.Bool("codegen", false, codeGenFlagHelp)

	flag.Parse()
	
	inputFileName := flag.Arg(0)
	outputFileName := inputFile[:len(inputFileName)-1] + "i"
	commandText = fmt.Sprintf("gcc -E -P $s -o $s", inputFileName, outputFileName)

	cmd := exec.Command("bash", "-c", commandText)

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(output))
}
