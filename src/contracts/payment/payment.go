package payment

import (
	paymentstatus "e-restaurant/models/enums/paymentStatus"
	"e-restaurant/models/payment"
)

type Repository interface {
	CreatePayment(payment.Payment)

	GetPaymentById(paymentId int)

	UpdatePaymentStatus(paymentId int, status paymentstatus.PaymentStatus)

	UpdatePaymentMethod(paymentId int, method paymentstatus.PaymentMethod)
}
