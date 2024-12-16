package handler

import (
	"admin-service/model"
	"admin-service/service"
	"encoding/json"
	"log"
	"net/http"
)

type AdminHandler struct {
	Service service.AdminService
}

func (h *AdminHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product

	err := h.Service.FetchProducts(&products)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding products to JSON: %v", err)
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User

	err := h.Service.FetchUsers(&users)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Error encoding users to JSON: %v", err)
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}
