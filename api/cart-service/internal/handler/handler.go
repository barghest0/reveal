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
	id, err := strconv.Atoi(vars["user_id"])

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
	id, err := strconv.Atoi(vars["cart_id"])

	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	var newProduct model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct.CartId = uint(id)

	cart, err := h.Service.GetCart(uint(id))
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	cart.Products = append(cart.Products, newProduct)

	if err := h.Service.UpdateCart(cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func (h *CartHandler) RemoveItemToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId, err := strconv.Atoi(vars["cart_id"])
	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		http.Error(w, "Invalid item id", http.StatusBadRequest)
		return
	}

	err = h.Service.RemoveItemToCart(uint(cartId), uint(productId))
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
