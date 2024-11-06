package router

import (
	"cart-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.CartHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cart/{user_id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart/{user_id}/products", h.AddProductToCart).Methods("POST")
	r.HandleFunc("/cart/{user_id}/products/{product_id}", h.UpdateProductQuantity).Methods("PUT")
	r.HandleFunc("/cart/{user_id}/products/{product_id}", h.RemoveProductFromCart).Methods("DELETE")

	return r
}
