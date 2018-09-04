package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mtojek/go-smart-contracts/inbox/contracts"
	"log"
	"strings"
)

const key = `--json file--`

func main() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/aa872bbe81694920ae9544a191a39999")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "XYZ")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := contracts.DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
}
