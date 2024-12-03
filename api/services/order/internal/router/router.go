package router

import (
	"product-service/internal/handler"

	"github.com/gorilla/mux"
)

func CreateRouter(h *handler.OrderHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orders", h.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/orders", h.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", h.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", h.DeleteOrder).Methods("DELETE")

	return r
}
