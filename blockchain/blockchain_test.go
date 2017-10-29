package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	t.Log("Hex should be 64 characters")
	hash := Hash([]byte("Hello World"))
	assert.Equal(t, len(hash), 64)
}

func TestHashBlock(t *testing.T) {
	t.Log("Block should be correctly marshalled")
	//TODO: generate new block and test its hashing
}
