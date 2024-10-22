package arbitrum

import (
	"context"

	"github.com/kryptogo/go-ethereum-arbitrum/arbitrum_types"
	"github.com/kryptogo/go-ethereum-arbitrum/core"
	"github.com/kryptogo/go-ethereum-arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
