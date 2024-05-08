package product

import "time"

//MODEL OF PRODUCT
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
