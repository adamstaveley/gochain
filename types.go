package gochain

type (
	// Tx describes a single transaction
	Tx struct {
		sender    string
		recipient string
		amount    int
	}

	// Block describes a single block
	Block struct {
		index        int // configure more specific type e.g. uint64
		previousHash string
		proof        int
		timestamp    int64
		transactions []Tx
	}

	// Chain holds all of blocks
	Chain []Block

	// ChainResponseBody is the JSON response body from other node's chain request
	ChainResponseBody struct {
		message string
		length  int
		chain   Chain
	}
)