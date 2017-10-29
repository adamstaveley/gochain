package node

import (
	"github.com/satori/go.uuid"
	"net/http"
)

var (
	nodeIdentifier string
)

// Init initialises node with desired host and port
func Init(host string) {
	createGenesisBlock()
	nodeIdentifier = uuid.NewV4().String()

	http.HandleFunc("/mine", mine)
	http.HandleFunc("/transactions/new", newTx)
	http.HandleFunc("/chain", chain)
	http.HandleFunc("/nodes/register", registerNode)
	http.HandleFunc("/nodes/resolve", resolveNode)
	http.ListenAndServe(host, nil)
}
