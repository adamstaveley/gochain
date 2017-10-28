package node

import (
	"github.com/adamstaveley/gochain/blockchain"
)

type (
	mineResponse struct {
		message string
		block blockchain.Block
	}

	chainResponse struct {
		chain []blockchain.Block
		length int
	}

	registerNodeRequest struct {
		nodes []string
	}

	registerNodeResponse struct {
		message string
		totalNodes int
	}

	resolveNodeResponse struct {
		message string
		chain []blockchain.Block
	}

)