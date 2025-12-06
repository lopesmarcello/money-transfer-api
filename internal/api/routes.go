package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *API) BindRoutes() {
	api.Router.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.Logger,
	)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/conta", api.handleCreateUser)
			r.Get("/conta/{id}/saldo", api.handleGetSaldo)
		})
	})
}

// Criar Conta (POST /conta) -> Testar
//
// Consultar Saldo (GET /conta/{id}/saldo)
// Depositar Dinheiro (POST /conta/{id}/deposito)
// Sacar Dinheiro (POST /conta/{id}/saque)
// Transferência (POST /conta/transferencia)
//
// Fechar Conta (DELETE /conta/{id})
