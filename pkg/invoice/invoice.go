package invoice

import (
	"github.com/IanRivas/composition/pkg/customer"
	"github.com/IanRivas/composition/pkg/invoiceitem"
)

//factura
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.Total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
