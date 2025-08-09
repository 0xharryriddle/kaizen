# Architecture overview

Kaizen aims to implement an EVM-compatible Ethereum client in modular layers:

- Node: process lifecycle, config, telemetry
- Networking: DevP2P (later), discovery, gossipsub (exploratory)
- Execution: EVM, state, tx pool
- Consensus: PoA dev mode first, pluggable interface for PoS
- RPC: JSON-RPC (HTTP/WebSocket)
- Storage: KV/Trie, blocks, receipts

This document will evolve as components solidify. Start with a minimal node skeleton and expand iteratively.
