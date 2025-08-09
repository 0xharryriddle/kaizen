# Kaizen

Kaizen is an Ethereum client implementation written in Go. This repo is a starting scaffold to build an EVM-compatible decentralized network node.

## Features (WIP)

- Minimal node bootstrap with CLI
- Pluggable components structure (p2p, rpc, consensus, execution)
- Ready-to-run dev mode stub

## Requirements

- Go 1.22+
- Git
- Windows PowerShell, macOS, or Linux shell

## Quick start

Build and run the minimal stub:

```powershell
# From repo root
go mod tidy
go build -o bin/kaizen ./cmd/kaizen
./bin/kaizen --help
./bin/kaizen run --dev
```

On Windows PowerShell you can run the binary as:

```powershell
./bin/kaizen.exe run --dev
```

## Project layout

```
cmd/
  kaizen/          # CLI entrypoint
internal/
  node/            # Node composition and lifecycle (WIP)
.docs/
  architecture/    # High-level docs
  development/     # Dev guides
```

## Development

- Use `go test ./...` to run tests
- Use `go vet ./...` and `golangci-lint` (optional) for static checks

### Common commands

```powershell
# Format
go fmt ./...

# Lint (optional)
# golangci-lint run

# Test
go test ./...

# Build CLI
go build -o bin/kaizen ./cmd/kaizen
```

## Roadmap (suggested)

1. Execution layer skeleton (EVM integration)
2. JSON-RPC server (eth, net, web3 minimal)
3. P2P networking (DevP2P protocol)
4. Consensus (PoA dev mode, then PoS interface)
5. Storage (state, blocks, tx pool)

## License

Apache-2.0 (suggested) – add `LICENSE` if desired.
