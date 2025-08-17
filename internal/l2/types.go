package l2

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Batch represents a collection of L2 transactions to be processed together
type Batch struct {
	ID           uint64               `json:"id"`
	Transactions []*types.Transaction `json:"transactions"`
	StateRoot    common.Hash          `json:"stateRoot"`
	Timestamp    time.Time            `json:"timestamp"`
	L1BlockHash  common.Hash          `json:"l1BlockHash"`
	ProofData    []byte               `json:"proofData,omitempty"`
}

// L2Block represents a block in the L2 chain
type L2Block struct {
	Header       *L2Header            `json:"header"`
	Transactions []*types.Transaction `json:"transactions"`
	Receipts     []*types.Receipt     `json:"receipts"`
}

// L2Header represents the header of an L2 block
type L2Header struct {
	Number      *big.Int    `json:"number"`
	ParentHash  common.Hash `json:"parentHash"`
	StateRoot   common.Hash `json:"stateRoot"`
	TxHash      common.Hash `json:"transactionsRoot"`
	ReceiptHash common.Hash `json:"receiptsRoot"`
	Timestamp   uint64      `json:"timestamp"`
	GasLimit    uint64      `json:"gasLimit"`
	GasUsed     uint64      `json:"gasUsed"`
	L1BlockHash common.Hash `json:"l1BlockHash"`
	BatchID     uint64      `json:"batchId"`
}

// StateCommitment represents a commitment to the L2 state
type StateCommitment struct {
	Root      common.Hash `json:"root"`
	BlockNum  uint64      `json:"blockNumber"`
	BatchID   uint64      `json:"batchId"`
	Proof     []byte      `json:"proof,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// L2Config holds configuration for the L2 chain
type L2Config struct {
	ChainID            uint64         `json:"chainId"`
	BatchSize          uint64         `json:"batchSize"`
	BatchTimeout       time.Duration  `json:"batchTimeout"`
	L1ChainID          uint64         `json:"l1ChainId"`
	BridgeAddress      common.Address `json:"bridgeAddress"`
	SequencerAddress   common.Address `json:"sequencerAddress"`
	EnableZK           bool           `json:"enableZk"`
	EnableConfidential bool           `json:"enableConfidential"`
}
