package models

type Product struct {
	//gorm.Model
	ID          int64
	Name        string
	Description string
	Price       float64
}

func (p Product) TableName() string {
	return "product"
}
