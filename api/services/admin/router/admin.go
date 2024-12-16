package router

import (
	"admin-service/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.AdminHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/admin/products", h.GetAllProducts).Methods("GET")
	r.HandleFunc("/admin/users", h.GetAllUsers).Methods("GET")

	return r
}
