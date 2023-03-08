package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

// SetUserStake set a specific userStake in the store from its index
func (k Keeper) SetUserStake(ctx sdk.Context, userStake types.UserStake) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserStakeKeyPrefix))
	b := k.cdc.MustMarshal(&userStake)
	store.Set(types.UserStakeKey(
		userStake.Creator,
	), b)
}

// GetUserStake returns a userStake from its index
func (k Keeper) GetUserStake(
	ctx sdk.Context,
	creator string,

) (val types.UserStake, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserStakeKeyPrefix))

	b := store.Get(types.UserStakeKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserStake removes a userStake from the store
func (k Keeper) RemoveUserStake(
	ctx sdk.Context,
	creator string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserStakeKeyPrefix))
	store.Delete(types.UserStakeKey(
		creator,
	))
}

// GetAllUserStake returns all userStake
func (k Keeper) GetAllUserStake(ctx sdk.Context) (list []types.UserStake) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserStakeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserStake
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
