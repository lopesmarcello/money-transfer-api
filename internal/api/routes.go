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
			r.Route("/conta", func(r chi.Router) {
				r.Post("/", api.handleCreateUser)
				r.Get("/{id}/saldo", api.handleGetSaldo)
				r.Post("/{id}/deposito", api.handleDeposit)
				r.Post("/{id}/saque", api.handleWithdraw)
				r.Post("/transferencia", api.handleTransfer)
				r.Delete("/{id}", api.handleCloseAccount)
			})
		})
	})
}
