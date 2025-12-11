package payment

import (
	paymentstatus "e-restaurant/models/enums/paymentStatus"
	"time"
)

type Payment struct {
	PaymentId int                         `db:"payment_id" json:"payment_id"`
	UserId    int                         `db:"user_id" json:"user_id"`
	OrderId   int                         `db:"order_id" json:"order_id"`
	Amount    float64                     `db:"amount" json:"amount"`
	Currency  string                      `db:"currency" json:"currency"`
	Method    paymentstatus.PaymentMethod `db:"method" json:"method"`
	Status    paymentstatus.PaymentStatus `db:"status" json:"status"`
	CreatedAt time.Time                   `db:"created_at" json:"created_at"`
}
