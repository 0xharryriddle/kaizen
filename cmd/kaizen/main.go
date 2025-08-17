package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/0xharryriddle/kaizen/internal/node"
)

func main() {
	// Enhanced CLI: kaizen run [--dev] [--l2-mode]
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	dev := runCmd.Bool("dev", false, "Run in developer mode with sane defaults")
	l2Mode := runCmd.Bool("l2-mode", false, "Run as Layer 2 node (default: true)")

	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "run":
		_ = runCmd.Parse(os.Args[2:])
		run(*dev, *l2Mode)
	case "help", "--help", "-h":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Println("kaizen <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  run       Start the Layer 2 node")
	fmt.Println()
	fmt.Println("Options for 'run':")
	fmt.Println("  --dev       Start in developer mode")
	fmt.Println("  --l2-mode   Run as Layer 2 node (default behavior)")
}

func run(dev bool, l2Mode bool) {
	cfg := node.Config{
		DevMode: dev,
		L2Mode:  l2Mode || true, // Default to L2 mode
	}
	n := node.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := n.Start(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start node: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("node started; press Ctrl+C to exit")
	// Minimal wait loop to simulate running server.
	for {
		time.Sleep(1 * time.Second)
	}
}
