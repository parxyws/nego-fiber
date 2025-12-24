package domain

type Order struct {
}

func (order *Order) TableName() string {
	return "orders"
}
