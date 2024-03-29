package factories

import (
	entities "saleinvoice/App/Domain/Entities"
	"time"
)

type SaleFactory struct{}

func (pf *SaleFactory) create(attributes map[string]interface{}) *entities.Sale {
	return &entities.Sale{
		Id:        attributes["id"].(int),
		InvoiceId: attributes["invoice_id"].(int),
		ProductId: attributes["product_id"].(int),
		Qty:       attributes["qty"].(int),
		Price:     attributes["price"].(float32),
		Dis:       attributes["dis"].(float32),
		Amount:    attributes["amount"].(float32),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
