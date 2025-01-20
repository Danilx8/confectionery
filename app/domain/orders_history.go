package domain

type OrdersHistory struct {
	Id        int    `gorm:"id;primary_key"`
	Date      string `gorm:"date;primary_key"`
	OrderId   string `gorm:"order_id"`
	Order     Order  `gorm:"foreignkey:OrderId"`
	OldStatus string `gorm:"old_status"`
	NewStatus string `gorm:"new_status"`
	Time      string `gorm:"time"`
}

type OrdersHistoryRepository interface {
	Create(history *OrdersHistory) error
	FetchAll() ([]OrdersHistory, error)
}
