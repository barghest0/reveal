package router

import (
	"cart-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.CartHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/carts", h.GetAllCarts).Methods("GET")
	r.HandleFunc("/cart", h.GetCart).Methods("GET")
	r.HandleFunc("/cart/products", h.AddProductToCart).Methods("POST")
	r.HandleFunc("/cart/products/{product_id}", h.UpdateProductQuantity).Methods("PUT")
	r.HandleFunc("/cart/products/{product_id}", h.RemoveProductFromCart).Methods("DELETE")

	return r
}
