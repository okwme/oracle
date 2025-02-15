package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/relevant-community/oracle/x/atom/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// AtomUsd implements the Query/AtomUsd gRPC method
func (k Keeper) AtomUsd(c context.Context, req *types.QueryAtomUsdRequest) (*types.QueryAtomUsdResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	atomUsd := k.GetAtomUsd(ctx)

	return &types.QueryAtomUsdResponse{AtomUsd: atomUsd}, nil
}
