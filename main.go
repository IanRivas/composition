package main

import (
	"fmt"

	"github.com/IanRivas/composition/pkg/customer"
	"github.com/IanRivas/composition/pkg/invoice"
	"github.com/IanRivas/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Argentina",
		"Buenos Aires",
		customer.New("Ian", "thompson", "1550501020"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de go", 9.99),
			invoiceitem.New(2, "Curso de go base de datos", 8.99),
			invoiceitem.New(3, "Curso de go apis", 12.99),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
