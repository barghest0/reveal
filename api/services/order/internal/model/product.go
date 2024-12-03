package model

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"   // Заказ создан, ожидает обработки
	OrderStatusConfirmed OrderStatus = "confirmed" // Заказ подтверждён
	OrderStatusShipped   OrderStatus = "shipped"   // Заказ отправлен
	OrderStatusDelivered OrderStatus = "delivered" // Заказ доставлен
	OrderStatusCanceled  OrderStatus = "canceled"  // Заказ отменён
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`               // Уникальный идентификатор заказа
	UserID     uint        `json:"user_id"`                            // ID пользователя, создавшего заказ
	Products   []OrderItem `gorm:"foreignKey:OrderID" json:"products"` // Список товаров в заказе
	TotalPrice float64     `json:"total_price"`                        // Итоговая стоимость заказа
	Status     OrderStatus `json:"status"`                             // Статус заказа
	// CreatedAt   time.Time     `json:"created_at"`                  // Время создания заказа
	// UpdatedAt   time.Time     `json:"updated_at"`                  // Время последнего изменения
	// Address     ShippingInfo  `json:"address"`                     // Информация о доставке
	// PaymentInfo PaymentDetail `json:"payment_info"`                // Информация об оплате
}

// OrderItem представляет товар в составе заказа.
type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`  // Уникальный идентификатор записи
	OrderID   uint    `gorm:"index" json:"order_id"` // Внешний ключ для связи с заказом
	ProductID uint    `json:"product_id"`            // ID товара
	Quantity  uint    `json:"quantity"`              // Количество единиц товара
	Price     float64 `json:"price"`                 // Цена за единицу товара
}

// // ShippingInfo содержит данные о доставке.
// type ShippingInfo struct {
// 	AddressLine1 string `json:"address_line_1"` // Основной адрес
// 	AddressLine2 string `json:"address_line_2"` // Дополнительный адрес (если есть)
// 	City         string `json:"city"`           // Город
// 	State        string `json:"state"`          // Штат/регион
// 	PostalCode   string `json:"postal_code"`    // Почтовый индекс
// 	Country      string `json:"country"`        // Страна
// }
//
// // PaymentDetail представляет информацию об оплате.
// type PaymentDetail struct {
// 	PaymentMethod string     `json:"payment_method"` // Метод оплаты (например, "credit_card", "paypal")
// 	TransactionID string     `json:"transaction_id"` // ID транзакции
// 	Amount        float64    `json:"amount"`         // Сумма оплаты
// 	PaidAt        *time.Time `json:"paid_at"`        // Время оплаты (если оплачено)
// }
