# Getting started (development)

## Prereqs

- Go 1.22+
- Make (optional)
- Windows: PowerShell 5.1+ or PowerShell 7

## Setup

```powershell
# Ensure modules resolved
go mod tidy

# Build
go build -o bin/kaizen ./cmd/kaizen

# Run dev node
./bin/kaizen run --dev
```

## VS Code

Recommended extensions:

- Go (golang.go)
- EditorConfig (optional)
- Markdown All in One

Debug configuration can be added later to launch the CLI with args.
