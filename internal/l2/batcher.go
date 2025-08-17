package l2

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// Batcher manages transaction batching for the L2
type Batcher struct {
	config     *L2Config
	pendingTxs []*types.Transaction
	mutex      sync.RWMutex
	batchChan  chan *Batch
	stopChan   chan struct{}
	wg         sync.WaitGroup
}

// NewBatcher creates a new transaction batcher
func NewBatcher(config *L2Config) *Batcher {
	return &Batcher{
		config:     config,
		pendingTxs: make([]*types.Transaction, 0),
		batchChan:  make(chan *Batch, 100),
		stopChan:   make(chan struct{}),
	}
}

// Start begins the batching process
func (b *Batcher) Start(ctx context.Context) {
	b.wg.Add(1)
	go b.batchingLoop(ctx)
}

// Stop stops the batching process
func (b *Batcher) Stop() {
	close(b.stopChan)
	b.wg.Wait()
}

// AddTransaction adds a transaction to the pending pool
func (b *Batcher) AddTransaction(tx *types.Transaction) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.pendingTxs = append(b.pendingTxs, tx)

	// Create batch immediately if we reach the target size
	if len(b.pendingTxs) >= int(b.config.BatchSize) {
		batch := b.createBatch()
		select {
		case b.batchChan <- batch:
			log.Info("Created batch due to size limit", "txCount", len(batch.Transactions))
		default:
			log.Warn("Batch channel full, dropping batch")
		}
	}

	return nil
}

// GetBatchChannel returns the channel for receiving new batches
func (b *Batcher) GetBatchChannel() <-chan *Batch {
	return b.batchChan
}

// batchingLoop runs the main batching logic
func (b *Batcher) batchingLoop(ctx context.Context) {
	defer b.wg.Done()

	ticker := time.NewTicker(b.config.BatchTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.stopChan:
			return
		case <-ticker.C:
			b.createTimedBatch()
		}
	}
}

// createTimedBatch creates a batch when the timeout is reached
func (b *Batcher) createTimedBatch() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.pendingTxs) > 0 {
		batch := b.createBatch()
		select {
		case b.batchChan <- batch:
			log.Info("Created batch due to timeout", "txCount", len(batch.Transactions))
		default:
			log.Warn("Batch channel full, dropping timed batch")
		}
	}
}

// createBatch creates a new batch from pending transactions
func (b *Batcher) createBatch() *Batch {
	txs := make([]*types.Transaction, len(b.pendingTxs))
	copy(txs, b.pendingTxs)

	batch := &Batch{
		ID:           uint64(time.Now().UnixNano()), // TODO: Use proper ID generation
		Transactions: txs,
		Timestamp:    time.Now(),
	}

	// Clear pending transactions
	b.pendingTxs = b.pendingTxs[:0]

	return batch
}

// GetPendingCount returns the number of pending transactions
func (b *Batcher) GetPendingCount() int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return len(b.pendingTxs)
}
