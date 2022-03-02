package test_tools

import (
	"github.com/vitelabs/go-vite/v2/common/types"
	ledger "github.com/vitelabs/go-vite/v2/interfaces/core"
	"github.com/vitelabs/go-vite/v2/ledger/consensus/core"
)

type MockConsensus struct{}

func (c *MockConsensus) VerifyABsProducer(abs map[types.Gid][]*ledger.AccountBlock) ([]*ledger.AccountBlock, error) {
	panic("implement me")
}

func (c *MockConsensus) SBPReader() core.SBPStatReader {
	return nil
}

func (c *MockConsensus) VerifyAccountProducer(block *ledger.AccountBlock) (bool, error) {
	return true, nil
}

type MockCssVerifier struct{}

func (c *MockCssVerifier) VerifyABsProducer(abs map[types.Gid][]*ledger.AccountBlock) ([]*ledger.AccountBlock, error) {
	return nil, nil
}

func (c *MockCssVerifier) VerifySnapshotProducer(block *ledger.SnapshotBlock) (bool, error) {
	return true, nil
}

func (c *MockCssVerifier) VerifyAccountProducer(block *ledger.AccountBlock) (bool, error) {
	return true, nil
}
