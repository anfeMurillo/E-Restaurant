package modelpayment

import (
	enumpayment "e-restaurant/models/enums/enumPayment"
	"time"
)

type Payment struct {
	PaymentId int                       `db:"payment_id" json:"payment_id"`
	UserId    int                       `db:"user_id" json:"user_id"`
	OrderId   int                       `db:"order_id" json:"order_id"`
	Amount    float64                   `db:"amount" json:"amount"`
	Currency  string                    `db:"currency" json:"currency"`
	Method    enumpayment.PaymentMethod `db:"method" json:"method"`
	Status    enumpayment.PaymentStatus `db:"status" json:"status"`
	CreatedAt time.Time                 `db:"created_at" json:"created_at"`
}
