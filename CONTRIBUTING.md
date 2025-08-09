# Contributing to Kaizen

Thanks for your interest in contributing!

## How to contribute

- Discuss significant changes in an issue first
- Fork the repo and create a branch
- Keep PRs focused and small when possible
- Add or update tests

## Development

- Go 1.22+
- `go test ./...` before pushing
- Run the dev node locally: `go build -o bin/kaizen.exe ./cmd/kaizen` then `./bin/kaizen.exe run --dev`

## Code style

- `go fmt ./...`
- Prefer clear, documented interfaces between packages

## Commit messages

- Conventional style appreciated (feat:, fix:, chore:, docs:, test:)
