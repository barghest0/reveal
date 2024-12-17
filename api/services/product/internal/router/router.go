package router

import (
	"product-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.ProductHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/products", h.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/products", h.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}/purchase", h.PurchaseProduct).Methods("POST")

	return r
}
