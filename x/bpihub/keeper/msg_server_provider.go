package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/madtechworks/bpihub/x/bpihub/types"
)

func (k msgServer) CreateProvider(goCtx context.Context, msg *types.MsgCreateProvider) (*types.MsgCreateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendProvider(
		ctx,
		msg.Creator,
		msg.Name,
		msg.Details,
		msg.Website,
	)

	return &types.MsgCreateProviderResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProvider(goCtx context.Context, msg *types.MsgUpdateProvider) (*types.MsgUpdateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var provider = types.Provider{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Details: msg.Details,
		Website: msg.Website,
	}

	// Checks that the element exists
	if !k.HasProvider(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetProviderOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetProvider(ctx, provider)

	return &types.MsgUpdateProviderResponse{}, nil
}

func (k msgServer) DeleteProvider(goCtx context.Context, msg *types.MsgDeleteProvider) (*types.MsgDeleteProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasProvider(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetProviderOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProvider(ctx, msg.Id)

	return &types.MsgDeleteProviderResponse{}, nil
}
