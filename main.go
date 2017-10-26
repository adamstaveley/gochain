package gochain

import (
	"flag"
	"log"
)

var (
	host = flag.String("h", "localhost", "Specify host of node")
	port = flag.String("p", "8000", "Specify port of node")
)

func main() {
	flag.Parse()
	log.Printf("Listening on: %s:%s", *host, *port)
}
