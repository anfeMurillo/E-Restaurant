package orderstatus

type OrderStatus string

const (
	OrderStatusReceived     OrderStatus = "received"
	OrderStatusInProcessing OrderStatus = "in_processing"
	OrderStatusCompleted    OrderStatus = "completed"
	OrderStatusCanceled     OrderStatus = "canceled"
)
