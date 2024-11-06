package handler

import (
	"cart-service/internal/model"
	"cart-service/internal/service"
	"encoding/json"
	"log"

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

func (h *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		http.Error(w, "Invalid cart id", http.StatusBadRequest)
		return
	}

	var product model.CartProduct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart, err := h.Service.GetCart(uint(user_id))
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	product.CartId = uint(cart.ID)

	cart.Products = append(cart.Products, product)

	// Сохраняем обновленную корзину
	if err := h.Service.UpdateCart(cart); err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *CartHandler) UpdateProductQuantity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var quantityData struct {
		Quantity uint `json:"quantity"`
	}

	// Декодируем новое количество из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&quantityData); err != nil {
		http.Error(w, "Invalid quantity data", http.StatusBadRequest)
		return
	}

	// Получаем корзину пользователя
	cart, err := h.Service.GetCart(uint(userID))
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	productFound := false
	for i, product := range cart.Products {
		if product.ProductID == uint(productID) {
			cart.Products[i].Quantity = quantityData.Quantity
			productFound = true
			break
		}
	}

	if !productFound {
		http.Error(w, "Product not found in cart", http.StatusNotFound)
		return
	}

	// Сохраняем обновленную корзину
	if err := h.Service.UpdateCart(cart); err != nil {
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart.Products)
}

func (h *CartHandler) RemoveProductFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	cart, err := h.Service.GetCart(uint(user_id))
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	err = h.Service.RemoveProductFromCart(uint(cart.UserId), uint(productId))
	if err != nil {
		log.Printf("Failed to remove item from cart: %v", err)
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
