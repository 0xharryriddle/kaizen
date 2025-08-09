package node

import (
	"context"
	"testing"
)

func TestNodeStartStop(t *testing.T) {
	n := New(Config{DevMode: true})
	ctx := context.Background()
	if err := n.Start(ctx); err != nil {
		t.Fatalf("start failed: %v", err)
	}
	if err := n.Stop(ctx); err != nil {
		t.Fatalf("stop failed: %v", err)
	}
}
