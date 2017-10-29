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
		Index        int // configure more specific type e.g. uint64
		PreviousHash string
		Proof        int
		Timestamp    int64
		Transactions []Tx
	}

	// ChainResponseBody is the JSON response body from other node's chain request
	ChainResponseBody struct {
		Message string
		Length  int
		Chain   []Block
	}
)