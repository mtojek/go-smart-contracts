package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mtojek/go-smart-contracts/inbox/contracts"
	"log"
)

func main() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/aa872bbe81694920ae9544a191a39999")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := contracts.NewInbox(common.HexToAddress("0xb5a030d6cf528c0a054208dfcf8c78c2d9ca9999"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	fmt.Println(contract.Message(nil))

}
