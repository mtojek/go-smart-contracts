package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mtojek/go-smart-contracts/inbox/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
)

const key  = `--json file--`

func main(){
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
	contract, err :=contracts.NewInbox(common.HexToAddress("0xb5a030d6cf528c0a054208dfcf8c78c2d9ca9999"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	contract.SetMessage(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "Hello From Mars")
}
