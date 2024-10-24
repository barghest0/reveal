package router

import (
	"cart-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.CartHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{userId}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart/{cartId}/item", h.AddItemToCart).Methods("PUT")
	r.HandleFunc("/cart/{cartId}/item/{itemId}", h.RemoveItemToCart).Methods("DELETE")

	return r
}
