package invoiceheader

import "time"

//modelo del invoiceheader
type Model struct {
	ID        uint
	Cliente   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//models slice of model
type Models []*Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
