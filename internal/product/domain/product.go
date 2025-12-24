package domain

type Product struct{}

func (p *Product) TableName() string {
	return "products"
}
