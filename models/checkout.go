package models

type CheckoutReqeust struct {
	CartId            string `json:"cart_id" form:"cart_id" validate:"required"`
	ShippingPrice     int    `json:"shipping_price" form:"shipping_price" validate:"required"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status" validate:"required"`
	Resi              string `json:"resi" form:"resi" validate:"required"`
	PaymentId         int    `json:"payment_id" form:"payment_id" validate:"required"`
}

type Checkout struct {
	CartId            string `json:"cart_id" form:"cart_id"`
	ShippingPrice     int    `json:"shipping_price" form:"shipping_price"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	Resi              string `json:"resi" form:"resi"`
	PaymentId         int    `json:"payment_id" form:"payment_id"`
	UserId            int    `json:"user_id" form:"user_id"`
}
