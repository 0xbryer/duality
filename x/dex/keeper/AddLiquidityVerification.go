package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddLiquidityVerification(goCtx context.Context, msg *types.MsgAddLiquidity) (string, string, sdk.AccAddress, sdk.AccAddress, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB)

	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	amount, err := sdk.NewDecFromStr(msg.Amount)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	//var decAmounts []sdk.Dec
	// for i := 0; i < len(msg.Amount); i++ {

	//amount, err := sdk.NewDecFromStr(msg.Amount)
	// // Error checking for valid sdk.Dec
	//if err != nil {
	// return "", "", nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	// }
	// decAmounts = append(decAmounts, amount)
	// }

	return token0, token1, callerAddr, receiverAddr, amount, nil
}