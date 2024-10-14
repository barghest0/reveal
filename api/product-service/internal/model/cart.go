package model

type Cart struct {
	ID       string
	Items    []CartItem
	Total    float64
	Currency string
}

type CartItem struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}
