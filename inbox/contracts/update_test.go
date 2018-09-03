package contracts

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)


// Test message gets updated correctly
func TestSetMessage(t *testing.T) {

	//Setup simulated blockchain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 0)

	//Deploy contract

	_, _, contract, _ :=DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	// commit all pending transactions
	blockchain.Commit()
	contract.SetMessage(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "Hello from Mars")

	blockchain.Commit()

	if got, _ := contract.Message(nil); got != "Hello from Mars" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}
