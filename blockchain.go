package gochain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	chain               Chain
	currentTransactions []Tx
	nodes               []string
)

// Hash creates SHA-256 with base64 string encoding
func Hash(text []byte) string {
	newsha256 := sha256.New()
	hash := newsha256.Sum(text)
	return base64.StdEncoding.EncodeToString(hash)
}

// HashBlock creates SHA-256 hash of block
func HashBlock(block Block) string {
	blockString, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	return Hash(blockString)
}

// ValidProof validates proof of Hash(p, p')
func ValidProof(lastProof, proof int) bool {
	guess := strconv.Itoa(lastProof) + strconv.Itoa(proof)
	guessHash := Hash([]byte(guess))
	return strings.HasSuffix(guessHash, "0000")
}

// LastBlock retreives block at final chain index
func LastBlock() Block {
	return chain[len(chain)-1]
}

// NewBlock creates new block on the chain
func NewBlock(proof int, previousHash string) Block {
	block := Block{
		len(chain),
		previousHash,
		proof,
		time.Now().Unix(),
		currentTransactions,
	}

	// reset current list of transactions
	currentTransactions = currentTransactions[:0]
	chain = append(chain, block)
	return block
}

// NewTransaction creates a new tx to go into next mined block
// returns index of the block that will hold this tx
func NewTransaction(tx Tx) int {
	currentTransactions = append(currentTransactions, tx)
	for index, block := range chain {
		if reflect.DeepEqual(block, LastBlock()) {
			return index
		}
	}
	return -1
}

// ProofOfWork finds p' such that Hash(pp') contains 4 leading zeros
// where p is previous p'
func ProofOfWork(lastProof int) int {
	proof := 0
	for !ValidProof(lastProof, proof) {
		proof++
	}
	return proof
}

// RegisterNode adds new node to list of known nodes
func RegisterNode(address string) {
	parsedURL, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	nodes = append(nodes, parsedURL.Host)
	return
}

// ResolveConflicts is the consensus algorithm
// replaces chain with longest in network
func ResolveConflicts() {
	maxLength := len(chain)
	var newChain []Block

	for index, node := range nodes {
		endpoint := fmt.Sprintf("http://%s/chain", node)
		response, err := http.Get(endpoint)
		if err != nil {
			panic(err)
		}

		if response.StatusCode == 200 {
			bodyString, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			var body ChainResponseBody
			json.Unmarshal(bodyString, body)

			length := body.length
			chain := body.chain

			if length > maxLength && ValidChain(chain) {
				maxLength = length
				newChain = chain
			}

			if index == len(nodes)-1 {
				if len(newChain) > len(chain) {
					chain = newChain
				}
			}
		}
	}
}

// ValidChain determines if given chain is valid
func ValidChain(chain Chain) bool {
	lastBlock := chain[0]
	currentIndex := 1

	for currentIndex < len(chain) {
		block := chain[currentIndex]
		log.Println(lastBlock)
		log.Println(block)
		log.Println("----------")
		// check if hash of block is correct
		if !reflect.DeepEqual(block.previousHash, HashBlock(lastBlock)) {
			return false
		}
		// check if PoW is correct
		if !ValidProof(lastBlock.proof, block.proof) {
			return false
		}

		lastBlock = block
		currentIndex++
	}
	return true
}
