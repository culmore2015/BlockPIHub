package bpihub

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/madtechworks/bpihub/x/bpihub/keeper"
	"github.com/madtechworks/bpihub/x/bpihub/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the provider
	for _, elem := range genState.ProviderList {
		k.SetProvider(ctx, *elem)
	}

	// Set provider count
	k.SetProviderCount(ctx, uint64(len(genState.ProviderList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all provider
	providerList := k.GetAllProvider(ctx)
	for _, elem := range providerList {
		elem := elem
		genesis.ProviderList = append(genesis.ProviderList, &elem)
	}

	return genesis
}
