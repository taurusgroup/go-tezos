// +build integration

package gotezos

import (
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Head_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)
	head, err := gt.Head()
	assert.Nil(t, err)
	assert.NotNil(t, head)
}

func Test_Block_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)
	head, err := gt.Block(100000)
	assert.Nil(t, err)
	assert.NotNil(t, head)
}

func Test_Operation_Hashes_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	hashes, err := gt.OperationHashes(head.Hash)
	assert.Nil(t, err)
	assert.NotNil(t, hashes)
}

func Test_Balance_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	hashes, err := gt.Balance(head.Hash, mockAddressTz1)
	assert.Nil(t, err)
	assert.NotNil(t, hashes)
}

func Test_Blocks_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	hashes, err := gt.Blocks(&BlocksInput{})
	assert.Nil(t, err)
	assert.NotNil(t, hashes)
}

func Test_ChainID_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	chainid, err := gt.ChainID()
	assert.Nil(t, err)
	assert.NotNil(t, chainid)
}

func Test_Checkpoint_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	checkpoint, err := gt.Checkpoint()
	assert.Nil(t, err)
	assert.NotNil(t, checkpoint)
}

func Test_Invalid_Blocks_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	invalidBlocks, err := gt.InvalidBlocks()
	assert.Nil(t, err)
	assert.NotNil(t, invalidBlocks)
}

func Test_DelegatedContracts_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegatedContracts, err := gt.DelegatedContracts(head.Hash, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, delegatedContracts)
}

func Test_DelegatedContractsAtCycle_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	delegatedContracts, err := gt.DelegatedContractsAtCycle(100, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, delegatedContracts)
}

func Test_FrozenBalance_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	frozenBalance, err := gt.FrozenBalance(100, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, frozenBalance)
}

func Test_Delegate_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegate, err := gt.Delegate(head.Hash, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_StakingBalance_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegate, err := gt.StakingBalance(head.Hash, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_StakingBalanceAtCycle_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	delegate, err := gt.StakingBalanceAtCycle(100, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_BakingRights_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegate, err := gt.BakingRights(&BakingRightsInput{
		BlockHash: &head.Hash,
	})
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_EndorsingRights_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegate, err := gt.EndorsingRights(&EndorsingRightsInput{
		BlockHash: &head.Hash,
	})
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_Delegates_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	delegate, err := gt.Delegates(&DelegatesInput{
		BlockHash: &head.Hash,
	})
	assert.Nil(t, err)
	assert.NotNil(t, delegate)
}

func Test_Version_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	version, err := gt.Version()
	assert.Nil(t, err)
	assert.NotNil(t, version)
}

func Test_Constants_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	head, err := gt.Block(100000)
	assert.Nil(t, err)

	constants, err := gt.Constants(head.Hash)
	assert.Nil(t, err)
	assert.NotNil(t, constants)
}

func Test_Connections_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	connections, err := gt.Connections()
	assert.Nil(t, err)
	assert.NotNil(t, connections)
}

func Test_Bootstrap_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	bootstrap, err := gt.Bootstrap()
	assert.Nil(t, err)
	assert.NotNil(t, bootstrap)
}

func Test_Commit_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	bootstrap, err := gt.Commit()
	assert.Nil(t, err)
	assert.NotNil(t, bootstrap)
}

func Test_Cycle_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	cycle, err := gt.Cycle(100)
	assert.Nil(t, err)
	assert.NotNil(t, cycle)
}

func Test_ActiveChains_Integration(t *testing.T) {
	gt := newGoTezosMainnet(t)

	activeChains, err := gt.ActiveChains()
	assert.Nil(t, err)
	assert.NotNil(t, activeChains)
}

func newGoTezosMainnet(t *testing.T) *GoTezos {
	gt, err := New(getMainnetIntegration())
	assert.Nil(t, err)
	gt.SetClient(&http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	})

	return gt
}

func getMainnetIntegration() string {
	return os.Getenv("TEZOS_MAINNET_URL")
}
