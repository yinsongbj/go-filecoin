package validation

import (
	"math/big"
	"testing"

	"github.com/filecoin-project/chain-validation/pkg/chain"
	"github.com/filecoin-project/chain-validation/pkg/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-filecoin/address"
	"github.com/filecoin-project/go-filecoin/types"
)

func TestMessageFactory(t *testing.T) {
	signer, keys := types.NewMockSignersAndKeyInfo(1)
	factory := NewMessageFactory(signer)
	p := chain.NewMessageProducer(factory)

	gasPrice := big.NewInt(1)
	gasLimit := big.NewInt(1000)

	sender, err := keys[0].Address()
	require.NoError(t, err)
	require.NoError(t, p.Transfer(state.Address(sender.Bytes()), state.BurntFundsAddress, big.NewInt(1), gasPrice, gasLimit))

	messages := p.Messages()
	assert.Equal(t, 1, len(messages))
	msg := messages[0].(*types.SignedMessage)
	assert.Equal(t, sender, msg.From)
	assert.Equal(t, address.BurntFundsAddress, msg.To)
	assert.Equal(t, types.NewAttoFIL(big.NewInt(1)), msg.Value)
}
