package node

import (
	"encoding/json"
	"github.com/adamstaveley/gochain/blockchain"
	"fmt"
	"net/http"
)

func createGenesisBlock() {
	blockchain.NewBlock(100, "1")
	fmt.Println(blockchain.Chain)
}

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
		Message: "New block forged",
		Block: block,
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
	fmt.Println(len(blockchain.Chain))
	response := chainResponse{
		Chain: blockchain.Chain,
		Length: len(blockchain.Chain),
	}
	fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		Message: "New nodes have been added",
		TotalNodes: len(blockchain.Nodes),
	}

	json.NewEncoder(w).Encode(response)

}

func resolveNode(w http.ResponseWriter, r *http.Request) {
	currentLength := len(blockchain.Chain)
	blockchain.ResolveConflicts()
	newLength := len(blockchain.Chain)

	var response resolveNodeResponse
	if currentLength < newLength {
		response.Message = "Chain was replaced"
	} else {
		response.Message = "Chain remains authorative"
	}

	response.Chain = blockchain.Chain
	json.NewEncoder(w).Encode(response)

}