package node

import (
	"github.com/adamstaveley/gochain/blockchain"
)

type (
	mineResponse struct {
		Message string
		Block blockchain.Block
	}

	chainResponse struct {
		Chain []blockchain.Block
		Length int
	}

	registerNodeRequest struct {
		nodes []string
	}

	registerNodeResponse struct {
		Message string
		TotalNodes int
	}

	resolveNodeResponse struct {
		Message string
		Chain []blockchain.Block
	}

)