package eth

import (
	"context"

	"github.com/kryptogo/go-ethereum-arbitrum/core"
	"github.com/kryptogo/go-ethereum-arbitrum/core/state"
	"github.com/kryptogo/go-ethereum-arbitrum/core/types"
	"github.com/kryptogo/go-ethereum-arbitrum/core/vm"
	"github.com/kryptogo/go-ethereum-arbitrum/eth/tracers"
	"github.com/kryptogo/go-ethereum-arbitrum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
