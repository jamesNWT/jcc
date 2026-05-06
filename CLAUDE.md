# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

**jcc** is a C compiler written in Go, built following the book *Writing A C Compiler* by Nora Sandler. It is an educational/learning project.

## Commands

```bash
go build          # Build the 'jcc' executable
go test ./...     # Run all tests
go fmt ./...      # Format all Go source files
./jcc <file.c>   # Run the compiler on a C source file
```

## Compiler Pipeline

The compiler follows a classic multi-stage pipeline, controlled by command-line flags:

```
*.c → Preprocessor (gcc -E -P) → *.i → Lexer → Parser → Code Generator → *.s → Assembler/Linker (gcc) → executable
```

Flags to stop at a given stage:
- `-lex`: stop after lexing
- `-parse`: stop after parsing
- `-codegen`: stop after code generation

Intermediate files: `.i` (preprocessed), `.s` (assembly). The output executable has no extension.

## Architecture

- **`main.go`**: Entry point; delegates entirely to the driver.
- **`driver/driver.go`**: Orchestrates the full pipeline. Handles preprocessing via `gcc -E -P`, invokes each compiler stage (currently stubs), creates the `.s` file, and links via `gcc`. This is where pipeline control logic and file management live.
- **`lexer/lexer.go`**: Stub — `Lexer()` is not yet implemented.

Parser and code generation packages do not yet exist; they are next to be created following the driver's existing stage-dispatch pattern.

The project has no third-party dependencies — only the Go standard library (`os/exec`, `flag`, `fmt`, `log`).
