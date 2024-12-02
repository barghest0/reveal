package service

import (
	"admin-service/internal/config"
	"admin-service/model"
	"encoding/json"
	"log"
	"net/http"
)

type AdminService interface {
	FetchProducts(products *[]model.Product) error
	FetchUsers(users *[]model.User) error
	FetchCarts(carts *[]model.Cart) error // Заглушка, получить все корзины пока нельзя
}

type adminService struct {
	cfg config.Config
}

func CreateAdminService(cfg config.Config) AdminService {
	return &adminService{cfg: cfg}
}

func (s *adminService) FetchProducts(products *[]model.Product) error {
	res, err := http.Get(s.cfg.ProductAPI + "/products")

	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&products); err != nil {
		log.Printf("Error decoding products response: %v", err)
		return err
	}

	return nil
}

func (s *adminService) FetchUsers(users *[]model.User) error {
	res, err := http.Get(s.cfg.UserAPI + "/users")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		log.Printf("Error decoding users response: %v", err)
		return err
	}

	return nil
}

func (s *adminService) FetchCarts(carts *[]model.Cart) error {
	res, err := http.Get(s.cfg.CartAPI + "/cart")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&carts); err != nil {
		log.Printf("Error decoding users response: %v", err)
		return err
	}

	return nil
}
