package handler

import (
	"cart-service/internal/model"
	"cart-service/internal/service"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CartHandler struct {
	Service service.CartService
}

func CreateCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{Service: service}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart model.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateCart(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["userId"])

	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}
	cartId, err := h.Service.GetCart(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cartId)
}

func (h *CartHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["cartId"])

	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	var newItem model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem.CartId = uint(id)

	cart, err := h.Service.GetCart(uint(id))
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	cart.Items = append(cart.Items, newItem)

	if err := h.Service.UpdateCart(cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func (h *CartHandler) RemoveItemToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId, err := strconv.Atoi(vars["cartId"])
	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	itemId, err := strconv.Atoi(vars["itemId"])
	if err != nil {
		http.Error(w, "Invalid item id", http.StatusBadRequest)
		return
	}

	err = h.Service.RemoveItemToCart(uint(cartId), uint(itemId))
	if err != nil {
		if err.Error() == "item not found" {
			http.Error(w, "Item not found in cart", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to remove item from cart", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item removed successfully"})
}
