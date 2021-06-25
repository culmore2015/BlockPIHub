package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers bpihub-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/bpihub/providers/{id}", getProviderHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/bpihub/providers", listProviderHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/bpihub/providers", createProviderHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bpihub/providers/{id}", updateProviderHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/bpihub/providers/{id}", deleteProviderHandler(clientCtx)).Methods("POST")

}
