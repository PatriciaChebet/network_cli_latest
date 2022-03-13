package main

import (
	"fmt"

	"github.com/PatriciaChebet/network_cli_latest/internal/network"
)

func main() {
	fmt.Println("Awesome CLI v0.0.1")
	network.Ping("127.0.0.1")
}
