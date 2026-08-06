# Minishell

A small Unix-like shell written in Go.

## Features

- Interactive REPL
- Built-in commands:
  - `cd`
  - `pwd`
  - `history`
  - `exit`
- External command execution
- Per-session command history
- Prompt that shortens the current directory relative to `$HOME`

## Requirements

- Go 1.24.5 or newer

## Run

```bash
go run ./cmd/mshell
```

## Project Layout

- `cmd/mshell` - application entrypoint
- `internal/shell` - REPL, prompt, and shell state
- `internal/executor` - built-in and external command execution
- `internal/parser` - tokenization and command parsing
- `internal/history` - history entry types
- `internal/process` - process-related helpers

## Notes

The parser currently accepts simple commands made of words only. Tokens for pipes and redirection are recognized, but pipeline and redirection execution are not implemented yet.
