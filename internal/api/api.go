// Package api
package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/lopesmarcello/money-transfer/internal/services"
)

type API struct {
	Router      *chi.Mux
	UserService services.UserService
}
