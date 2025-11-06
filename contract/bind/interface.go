package bind

import (
	"context"

	"github.com/mavryk-network/gomavryk/contract"
	"github.com/mavryk-network/gomavryk/mavryk"
	"github.com/mavryk-network/gomavryk/micheline"
	"github.com/mavryk-network/gomavryk/rpc"
)

type Contract interface {
	Address() mavryk.Address
	Call(ctx context.Context, args contract.CallArguments, opts *rpc.CallOptions) (*rpc.Receipt, error)
	RunView(ctx context.Context, name string, args micheline.Prim) (micheline.Prim, error)
}

type RPC interface {
	GetContractStorage(ctx context.Context, addr mavryk.Address, id rpc.BlockID) (micheline.Prim, error)
	GetBigmapValue(ctx context.Context, bigmap int64, hash mavryk.ExprHash, id rpc.BlockID) (micheline.Prim, error)
}

var (
	_ Contract = &contract.Contract{}
	_ RPC      = &rpc.Client{}
)
