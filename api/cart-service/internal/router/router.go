package router

import (
	"cart-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.CartHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cart/{user_id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{cart_id}/products/{product_id}", h.RemoveItemToCart).Methods("DELETE")
	r.HandleFunc("/cart/{cart_id}/products", h.AddItemToCart).Methods("PUT")

	return r
}
