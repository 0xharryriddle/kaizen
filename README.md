# Kaizen

Kaizen is an EVM-compatible, confidential, and high-performance Layer 2 (L2) rollup implementation written in Go. Inspired by leading L2 solutions like Arbitrum and ZkSync, Kaizen aims to provide scalable, private, and efficient transaction processing on top of Ethereum.

## Features (WIP)

- Modular Layer 2 architecture (execution, settlement, data availability, zk modules)
- EVM compatibility (using go-ethereum as a library)
- Confidential transactions (zk-SNARK/zk-STARK integration)
- High-performance rollup logic (batch processing, parallel execution)
- Bridge to Ethereum L1 (deposits, withdrawals, cross-chain messaging)
- JSON-RPC API for L2 node interaction
- P2P networking for L2 validator nodes
- Ready-to-run dev mode stub

## Requirements

- Go 1.22+
- Git
- Windows PowerShell, macOS, or Linux shell
- (Future) Dependencies: go-ethereum, gnark (for zk-proofs), leveldb/rocksdb

## Quick start

Build and run the minimal Layer 2 node stub:

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
  l2/              # Layer 2 rollup logic (batching, state management)
  evm/             # EVM integration and execution engine
  zk/              # Zero-knowledge proof system for confidentiality
  bridge/          # L1-L2 bridge logic (deposits/withdrawals)
  consensus/       # L2 consensus mechanism
  storage/         # State storage and data availability
.docs/
  architecture/    # High-level architecture docs
  development/     # Development guides and tutorials
  l2-design/       # Layer 2 specific design documents
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

## Development

- Use `go test ./...` to run tests
- Use `go vet ./...` and `golangci-lint` (optional) for static checks
- Build and test L2 components incrementally
- Use dev mode for local testing without L1 connection

### Common commands

```powershell
# Format
go fmt ./...

# Lint (optional)
# golangci-lint run

# Test all modules
go test ./...

# Build L2 node
go build -o bin/kaizen ./cmd/kaizen

# Run in development mode
./bin/kaizen.exe run --dev --l2-mode
```

## Layer 2 Migration Roadmap

### Phase 1: Foundation (Weeks 1-2)

1. **Project Structure Setup**

   - Create modular directory structure for L2 components
   - Update go.mod with L2-specific dependencies
   - Set up basic CLI for L2 operations

2. **EVM Integration Skeleton**
   - Integrate go-ethereum as a library
   - Create basic transaction execution framework
   - Implement state management interfaces

### Phase 2: Core L2 Logic (Weeks 3-5)

3. **Rollup Mechanism**

   - Implement batch processing logic
   - Create state commitment and merkle tree structures
   - Add transaction sequencing and ordering

4. **Bridge to Ethereum L1**
   - Design deposit/withdrawal contracts
   - Implement cross-chain message passing
   - Add L1 state synchronization

### Phase 3: Confidentiality (Weeks 6-8)

5. **Zero-Knowledge Integration**

   - Integrate zk-SNARK library (gnark)
   - Implement private transaction processing
   - Add confidential state transitions

6. **Privacy Features**
   - Anonymous transaction pools
   - Confidential smart contract execution
   - Privacy-preserving bridges

### Phase 4: Performance & Networking (Weeks 9-11)

7. **High Performance Optimizations**

   - Parallel transaction execution
   - Efficient state storage (LevelDB/RocksDB)
   - Batch compression and optimization

8. **P2P Networking**
   - Custom L2 networking protocol
   - Validator node communication
   - Consensus mechanism for L2

### Phase 5: APIs & Integration (Weeks 12-14)

9. **JSON-RPC API**

   - Ethereum-compatible RPC endpoints
   - L2-specific API extensions
   - WebSocket support for real-time updates

10. **Testing & Security**
    - Comprehensive unit and integration tests
    - Fuzz testing for critical components
    - Security audit preparation

## Step-by-Step Build Tutorials

### Tutorial 1: Setting Up L2 Project Structure

```powershell
# 1. Create new module directories
mkdir internal\l2, internal\evm, internal\zk, internal\bridge
mkdir internal\consensus, internal\storage
mkdir .docs\l2-design

# 2. Update go.mod with L2 dependencies
go mod edit -require=github.com/ethereum/go-ethereum@latest
go mod edit -require=github.com/consensys/gnark@latest
go mod tidy
```

### Tutorial 2: EVM Integration

```go
// internal/evm/engine.go - Basic EVM execution engine
package evm

import (
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/core/vm"
)

type Engine struct {
    vmConfig vm.Config
    chainConfig *params.ChainConfig
}

func (e *Engine) ExecuteTransaction(tx *types.Transaction) (*types.Receipt, error) {
    // L2-specific transaction execution logic
}
```

### Tutorial 3: Rollup Batch Processing

```go
// internal/l2/batcher.go - Transaction batching
package l2

type Batcher struct {
    pendingTxs []Transaction
    batchSize  int
}

func (b *Batcher) CreateBatch() (*Batch, error) {
    // Create optimized transaction batches
}
```

### Tutorial 4: ZK Integration

```go
// internal/zk/prover.go - Zero-knowledge proofs
package zk

import "github.com/consensys/gnark/frontend"

func GenerateProof(circuit frontend.Circuit, witness frontend.Circuit) (Proof, error) {
    // Generate zk-SNARK proofs for confidential transactions
}
```

### Tutorial 5: L1-L2 Bridge

```go
// internal/bridge/bridge.go - Cross-chain bridge
package bridge

type Bridge struct {
    l1Client *ethclient.Client
    l2Engine *evm.Engine
}

func (b *Bridge) ProcessDeposit(deposit *L1Deposit) error {
    // Process deposits from L1 to L2
}
```

## Roadmap (Layer 2 Focused)

1. Layer 2 architecture foundation and project restructuring
2. EVM integration using go-ethereum library
3. Rollup mechanism with batch processing and state commitments
4. Bridge to Ethereum L1 (deposits, withdrawals, messaging)
5. Zero-knowledge integration for confidential transactions
6. High-performance optimizations (parallel execution, efficient storage)
7. P2P networking and L2 consensus mechanism
8. JSON-RPC API with Ethereum compatibility
9. Comprehensive testing, security audits, and documentation
10. Mainnet deployment and ecosystem integration

## Next Steps

1. **Immediate Actions:**

   - Review and understand Arbitrum and ZkSync architectures
   - Set up the new project structure with L2 modules
   - Begin EVM integration planning

2. **Week 1 Goals:**

   - Complete project restructuring
   - Implement basic L2 node skeleton
   - Start EVM integration with go-ethereum

3. **Resources:**
   - Study Arbitrum's Nitro architecture
   - Review ZkSync's zkEVM implementation
   - Explore Polygon's zkEVM for additional insights

## License

Apache-2.0 (suggested) – add `LICENSE` if desired.
