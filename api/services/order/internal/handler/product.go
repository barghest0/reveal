package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"product-service/internal/model"
	"product-service/internal/service"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service service.OrderService
}

func CreateProductHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var product model.Order
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateOrder(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	idsParam := r.URL.Query().Get("ids")

	// Если параметр ids пустой, передаем пустой срез
	var idsInt []int
	if idsParam != "" {
		// Преобразуем строку в срез целых чисел
		ids := strings.Split(idsParam, ",")
		for _, idStr := range ids {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid id format", http.StatusBadRequest)
				return
			}
			idsInt = append(idsInt, id)
		}
	}

	products, err := h.service.GetOrders(idsInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetOrder(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	var product model.Order
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = uint(id)
	if err := h.service.UpdateOrder(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteOrder(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
