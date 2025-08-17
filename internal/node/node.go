package node

import (
	"context"
	"fmt"
)

// Config holds node configuration.
type Config struct {
	DevMode bool
	L2Mode  bool // Run as Layer 2 node
}

// Node is the top-level process holder.
type Node struct {
	cfg Config
}

// New creates a new Node.
func New(cfg Config) *Node {
	return &Node{cfg: cfg}
}

// Start brings the node up.
func (n *Node) Start(ctx context.Context) error {
	if n.cfg.DevMode {
		fmt.Println("[kaizen] starting Layer 2 node in DEV mode: using ephemeral config")
	} else {
		fmt.Println("[kaizen] starting Layer 2 node")
	}

	if n.cfg.L2Mode {
		fmt.Println("[kaizen] initializing Layer 2 components:")
		fmt.Println("  - EVM execution engine")
		fmt.Println("  - Transaction batcher")
		fmt.Println("  - ZK prover (placeholder)")
		fmt.Println("  - L1-L2 bridge (placeholder)")
		fmt.Println("  - Storage layer")
	}

	// TODO: initialize L2 subsystems (l2, evm, zk, bridge, storage, rpc, p2p)
	return nil
}

// Stop gracefully shuts down the node.
func (n *Node) Stop(ctx context.Context) error {
	fmt.Println("[kaizen] stopping Layer 2 node")
	// TODO: gracefully shutdown L2 components
	return nil
}
