package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/madtechworks/bpihub/x/bpihub/types"
	"strconv"
)

// GetProviderCount get the total number of provider
func (k Keeper) GetProviderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderCountKey))
	byteKey := types.KeyPrefix(types.ProviderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetProviderCount set the total number of provider
func (k Keeper) SetProviderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderCountKey))
	byteKey := types.KeyPrefix(types.ProviderCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendProvider appends a provider in the store with a new id and update the count
func (k Keeper) AppendProvider(
	ctx sdk.Context,
	creator string,
	name string,
	details string,
	website string,
) uint64 {
	// Create the provider
	count := k.GetProviderCount(ctx)
	var provider = types.Provider{
		Creator: creator,
		Id:      count,
		Name:    name,
		Details: details,
		Website: website,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	value := k.cdc.MustMarshalBinaryBare(&provider)
	store.Set(GetProviderIDBytes(provider.Id), value)

	// Update provider count
	k.SetProviderCount(ctx, count+1)

	return count
}

// SetProvider set a specific provider in the store
func (k Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	b := k.cdc.MustMarshalBinaryBare(&provider)
	store.Set(GetProviderIDBytes(provider.Id), b)
}

// GetProvider returns a provider from its id
func (k Keeper) GetProvider(ctx sdk.Context, id uint64) types.Provider {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	var provider types.Provider
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetProviderIDBytes(id)), &provider)
	return provider
}

// HasProvider checks if the provider exists in the store
func (k Keeper) HasProvider(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	return store.Has(GetProviderIDBytes(id))
}

// GetProviderOwner returns the creator of the provider
func (k Keeper) GetProviderOwner(ctx sdk.Context, id uint64) string {
	return k.GetProvider(ctx, id).Creator
}

// RemoveProvider removes a provider from the store
func (k Keeper) RemoveProvider(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	store.Delete(GetProviderIDBytes(id))
}

// GetAllProvider returns all provider
func (k Keeper) GetAllProvider(ctx sdk.Context) (list []types.Provider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Provider
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProviderIDBytes returns the byte representation of the ID
func GetProviderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetProviderIDFromBytes returns ID in uint64 format from a byte array
func GetProviderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
