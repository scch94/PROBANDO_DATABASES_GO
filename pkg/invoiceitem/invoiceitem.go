package invoiceitem

import "time"

type invoiceitem struct {
	ID         uint
	invoice_id uint
	product_id uint
	created_at time.Time
	updated_at time.Time
}
