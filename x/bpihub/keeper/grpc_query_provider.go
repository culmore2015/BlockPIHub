package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/madtechworks/bpihub/x/bpihub/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProviderAll(c context.Context, req *types.QueryAllProviderRequest) (*types.QueryAllProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var providers []*types.Provider
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	providerStore := prefix.NewStore(store, types.KeyPrefix(types.ProviderKey))

	pageRes, err := query.Paginate(providerStore, req.Pagination, func(key []byte, value []byte) error {
		var provider types.Provider
		if err := k.cdc.UnmarshalBinaryBare(value, &provider); err != nil {
			return err
		}

		providers = append(providers, &provider)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProviderResponse{Provider: providers, Pagination: pageRes}, nil
}

func (k Keeper) Provider(c context.Context, req *types.QueryGetProviderRequest) (*types.QueryGetProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var provider types.Provider
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasProvider(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetProviderIDBytes(req.Id)), &provider)

	return &types.QueryGetProviderResponse{Provider: &provider}, nil
}
