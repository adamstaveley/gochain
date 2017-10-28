package main

import (
	"flag"
	"fmt"
	"github.com/adamstaveley/gochain/node"
	"log"
)

var (
	host = flag.String("h", "localhost", "Specify host of node")
	port = flag.String("p", "8000", "Specify port of node")
)

func main() {
	flag.Parse()
	nodeHost := fmt.Sprintf("%s:%s", *host, *port) 
	log.Printf("Listening on: http://%s", nodeHost)
	node.Init(nodeHost)
}
