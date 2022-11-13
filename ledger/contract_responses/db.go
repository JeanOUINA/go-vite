package contract_responses

import (
	"path"

	chain_db "github.com/vitelabs/go-vite/v2/ledger/chain/db"
)

type ContractResponses struct {
	store *chain_db.Store
}

var ContractResponsesInstance *ContractResponses

func NewContractResponses(chainDir string) (*ContractResponses, error) {
	dataDir := path.Join(chainDir, "plugins")

	store, err := chain_db.NewStore(dataDir, "contract_responses")
	if err != nil {
		return nil, err
	}

	ContractResponsesInstance = &ContractResponses{
		store,
	}
	return ContractResponsesInstance, nil
}

func (cr *ContractResponses) GetResponse(tx []byte) ([]byte, error) {
	return cr.store.Get(tx)
}

func (cr *ContractResponses) SetResponse(tx []byte, value []byte) error {
	batch := cr.store.NewBatch()
	batch.Put(tx, value)
	cr.store.WriteDirectly(batch)
	return nil
}
