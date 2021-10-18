package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/crypto"
)

func TestMakeGenesisAccountConfig(t *testing.T) {
	cfg := makeMainnetGenesisCfg()
	if !IsCompleteGenesisConfig(cfg) {
		t.Fatalf("convert genesis config failed")
	}
}

func TestGenesisCfg(t *testing.T) {
	hash, err := types.BytesToHash(crypto.Hash256([]byte(GenesisJson())))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hash)
	assert.Equal(t, "79803fa6fc50f7e7ce18366fd6595394d6d257a191ed47850ca09609749e8f21", hash.String())
}
