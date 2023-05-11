package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserLimitOrdersAll(
	goCtx context.Context,
	req *types.QueryAllUserLimitOrdersRequest,
) (*types.QueryAllUserLimitOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	profile := NewUserProfile(addr)
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryAllUserLimitOrdersResponse{
		LimitOrders: profile.GetAllLimitOrders(ctx, k),
	}, nil
}