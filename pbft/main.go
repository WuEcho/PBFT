package main

import (
	"os"
	"kongyixueyuan.com/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
