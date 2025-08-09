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
	// Simple CLI: kaizen run [--dev]
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	dev := runCmd.Bool("dev", false, "Run in developer mode with sane defaults")

	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "run":
		_ = runCmd.Parse(os.Args[2:])
		run(*dev)
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
	fmt.Println("  run       Start the node")
	fmt.Println()
	fmt.Println("Options for 'run':")
	fmt.Println("  --dev     Start in developer mode")
}

func run(dev bool) {
	cfg := node.Config{DevMode: dev}
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
