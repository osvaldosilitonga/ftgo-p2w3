package dto

import "ngc11/entity"

type TransactionReq struct {
	Data []TransactionReqBody `json:"orders" validate:"required"`
}

type TransactionReqBody struct {
	ProductID int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required"`
}

type ErrTransactionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TransactionResponse struct {
	Code     int               `json:"code"`
	Products []entity.Products `json:"products"`
}
