package database

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"strconv"
	"strings"
)

func CheckoutChart(checkout *models.Checkout) (interface{}, error) {

	// rubah cart id jadi array

	// fmt.Println(cartsId)
	// var transactionDetail models.TransactionDetails
	// totalBelanja := 0
	// for _, v := range cartsId {
	// 	var cart models.CartCheckout
	// 	e := config.DB.Table("carts").Select("carts.id, carts.products_id, carts.users_id, carts.qty, products.price").Joins("join products on products.id = carts.products_id").Where("carts.id = ?", v).Where("carts.deleted_at is null").First(&cart).Error
	// 	if e != nil {
	// 		fmt.Println("ada error")
	// 	} else {
	// 		fmt.Println(cart.Price)
	// 		//
	// 		temp := cart.Qty * cart.Price
	// 		totalBelanja += temp
	// 		// transactionDetail.
	// 	}
	// }
	var carts []models.Carts
	cartsId := strings.Split(checkout.CartId, ",")
	config.DB.Preload("Products").Find(&carts, cartsId)
	totalBelanja := 0
	for _, v := range carts {
		l, _ := strconv.Atoi(v.Qty)
		temp := l * v.Products.Price
		totalBelanja += temp
	}

	transaction := models.Transactions{
		UserId:             uint(checkout.UserId),
		ShippingPrice:      checkout.ShippingPrice,
		TotalPrice:         totalBelanja + checkout.ShippingPrice,
		TransactionsStatus: checkout.TransactionStatus,
		Resi:               checkout.Resi,
		PaymentId:          uint(checkout.PaymentId),
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}
	// msukan ka dalam cart detail

	for _, v := range carts {
		transactionDetail := models.TransactionDetails{
			TransactionId: transaction.ID,
			ProductsId:    v.ProductsId,
			Qty:           v.Qty,
		}
		if err := config.DB.Create(&transactionDetail).Error; err != nil {
			return nil, err
		}
		if e := config.DB.Delete(&carts, v.ID).Error; e != nil {
			return nil, e
		}

	}

	return transaction, nil

}
