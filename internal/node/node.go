package node

import (
	"context"
	"fmt"
)

// Config holds node configuration.
type Config struct {
	DevMode bool
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
		fmt.Println("[kaizen] starting in DEV mode: using ephemeral config")
	} else {
		fmt.Println("[kaizen] starting")
	}
	// TODO: initialize subsystems (storage, rpc, p2p, execution)
	return nil
}

// Stop gracefully shuts down the node.
func (n *Node) Stop(ctx context.Context) error {
	fmt.Println("[kaizen] stopping")
	return nil
}
