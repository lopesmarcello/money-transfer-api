package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *API) BindRoutes() {
	api.Router.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.Logger,
	)

	CSRF := csrf.Protect(
		[]byte("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
		csrf.Secure(false),
	)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(CSRF)
			r.Get("/csrf-token", api.getCsrfToken)
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
