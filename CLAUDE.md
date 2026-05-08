# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

**jcc** is a C compiler written in Go, built following the book *Writing A C Compiler* by Nora Sandler. It is an educational/learning project.

## Commands

```bash
go build              # Build the 'jcc' executable
go test ./...         # Run all tests
go test ./... -run <TestName>  # Run a single test
go fmt ./...          # Format all Go source files
./jcc <file.c>        # Run the compiler on a C source file
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
- **`driver/driver.go`**: Orchestrates the full pipeline. Parses CLI flags, runs the gcc preprocessor, dispatches to each compiler stage (currently stubs), and links the final executable via `gcc`. The base filename is derived by stripping the last character of the input path (assumes a single-char extension like `.c`), then `addExtension` appends `.i`, `.s`, etc.
- **`lexer/lexer.go`**: Stub — `Lexer()` is not yet implemented.

**Current state:** The driver's compile-to-assembly step is a placeholder — it invokes `gcc -S` directly on the source file rather than routing through jcc's own stages. The `-lex`, `-parse`, and `-codegen` flag branches are stubbed out. The next step is implementing `Lexer()` and wiring it into the driver's `-lex` branch.

New compiler stages should be added as separate packages (e.g., `parser/`, `codegen/`) following the `lexer` package pattern, then dispatched from `driver.go` inside the corresponding flag branch.

**Test fixtures:** `driver/tests/` contains `.c` files for manual end-to-end testing.

The project has no third-party dependencies — only the Go standard library (`os/exec`, `flag`, `fmt`, `log`).
