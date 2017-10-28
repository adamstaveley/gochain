package blockchain

type (
	// Tx describes a single transaction
	Tx struct {
		Sender    string
		Recipient string
		Amount    int
	}

	// Block describes a single block
	Block struct {
		index        int // configure more specific type e.g. uint64
		previousHash string
		Proof        int
		timestamp    int64
		transactions []Tx
	}

	// ChainResponseBody is the JSON response body from other node's chain request
	ChainResponseBody struct {
		message string
		length  int
		chain   []Block
	}
)