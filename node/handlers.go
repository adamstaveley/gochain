package node

import (
	"encoding/json"
	"github.com/adamstaveley/gochain/blockchain"
	"fmt"
	"net/http"
)

func mine(w http.ResponseWriter, r *http.Request) {
	lastBlock := blockchain.LastBlock()
	lastProof := lastBlock.Proof
	proof := blockchain.ProofOfWork(lastProof)

	tx := blockchain.Tx{
		Sender: "0", 
		Recipient: nodeIdentifier, 
		Amount: 1,
	}

	blockchain.NewTransaction(tx)
	previousHash := blockchain.HashBlock(blockchain.LastBlock())
	block := blockchain.NewBlock(proof, previousHash)

	response := mineResponse{
		message: "New block forged",
		block: block,
	}

	json.NewEncoder(w).Encode(response) // return JSON object
}

func newTx(w http.ResponseWriter, r *http.Request) {
	var body blockchain.Tx
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	index := blockchain.NewTransaction(body)
	fmt.Fprintf(w, "Transaction will be added to next block %d", index)
}

func chain(w http.ResponseWriter, r *http.Request) {
	response := chainResponse{
		chain: blockchain.Chain,
		length: len(blockchain.Chain),
	}
	json.NewEncoder(w).Encode(response)
}

func registerNode(w http.ResponseWriter, r *http.Request) {
	var body registerNodeRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close() // contentious
	nodes := body.nodes

	if len(nodes) < 1 {
		fmt.Fprintf(w, "Error: please supply valid list of nodes")
		return
	}

	for _, node := range nodes {
		blockchain.RegisterNode(node)
	}

	response := registerNodeResponse{
		message: "New nodes have been added",
		totalNodes: len(blockchain.Nodes),
	}

	json.NewEncoder(w).Encode(response)

}

func resolveNode(w http.ResponseWriter, r *http.Request) {
	currentLength := len(blockchain.Chain)
	blockchain.ResolveConflicts()
	newLength := len(blockchain.Chain)

	var response resolveNodeResponse
	if currentLength < newLength {
		response.message = "Chain was replaced"
	} else {
		response.message = "Chain remains authorative"
	}

	response.chain = blockchain.Chain
	json.NewEncoder(w).Encode(response)

}