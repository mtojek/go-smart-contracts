package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

//Test initial message gets set up correctly
func TestGetMessage(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 0)

	//Deploy contract
	_, _, contract, _ := DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	// commit all pending transactions
	blockchain.Commit()

	if got, _ := contract.Message(nil); got != "Hello World" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}

}
