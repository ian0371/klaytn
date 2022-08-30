package governance

import (
	"testing"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractConnector(t *testing.T) {
	var (
		name   = "istanbul.committeesize"
		valueA = []byte{0xa}
		valueB = []byte{0xbb, 0xbb}
		p      = map[string][]byte{
			name: valueA,
		}
	)

	accounts, sim, _, contract := prepareSimulatedContractWithParams(t, p)

	owner := accounts[0]

	// Value exists after SetParam()
	names, values, err := contract.GetAllParams(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(names))
	assert.Equal(t, 1, len(values))
	assert.Equal(t, name, names[0])
	assert.Equal(t, valueA, values[0])

	// Call SetParam() again
	ab := sim.BlockChain().CurrentHeader().Number.Uint64() + 2
	tx, err := contract.SetParam(owner, name, valueB, ab)
	require.Nil(t, err)

	// increase block number to reach activation block
	for sim.BlockChain().CurrentHeader().Number.Uint64() < ab {
		sim.Commit()
	}

	receipt, _ := sim.TransactionReceipt(nil, tx.Hash())
	require.NotNil(t, receipt)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// Value changed after SetParam()
	names, values, err = contract.GetAllParams(nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(names))
	assert.Equal(t, 1, len(values))
	assert.Equal(t, name, names[0])
	assert.Equal(t, valueB, values[0])
}

func TestContractConnector_getAddressAt(t *testing.T) {
	accounts, sim, gpaddr, _ := prepareSimulatedContract(t)
	_, contract := prepareRegistry(t, accounts[0], sim, gpaddr)
	ret, err := contract.GetAddress(nil, GOVPARAM_NAME)
	assert.Nil(t, err)
	assert.Equal(t, gpaddr, ret)
}
