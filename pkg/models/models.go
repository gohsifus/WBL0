package models

import "time"

type Item struct {
	Id          int    `json:"id,omitempty"` //omitempty чтобы поле в случае пустоты не вносилось в json при marshal
	ChrtId      int    `json:"chrt_id"validate:"required"`
	TrackNumber string `json:"track_number"validate:"required"`
	Price       int    `json:"price"validate:"required"`
	Rid         string `json:"rid"validate:"required"`
	Name        string `json:"name"validate:"required"`
	Sale        int    `json:"sale"validate:"required"`
	Size        string `json:"size"validate:"required"`
	TotalPrice  int    `json:"total_price"validate:"required"`
	NmId        int    `json:"nm_id"validate:"required"`
	Brand       string `json:"brand"validate:"required"`
	Status      int    `json:"status"validate:"required"`
	OrderId     int    `json:"order_id"validate:"required"`
}

type Delivery struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name"validate:"required"`
	Phone   string `json:"phone"validate:"required"`
	Zip     string `json:"zip"validate:"required"`
	City    string `json:"city"validate:"required"`
	Address string `json:"address"validate:"required"`
	Region  string `json:"region"validate:"required"`
	Email   string `json:"email"validate:"required"`
}

type Payment struct {
	Id           int    `json:"id,omitempty"`
	Transaction  string `json:"transaction"validate:"required"`
	RequestId    string `json:"request_id"validate:"required"`
	Currency     string `json:"currency"validate:"required"`
	Provider     string `json:"provider"validate:"required"`
	Amount       int    `json:"amount"validate:"required"`
	PaymentDt    int    `json:"payment_dt"validate:"required"`
	Bank         string `json:"bank"validate:"required"`
	DeliveryCost int    `json:"delivery_cost"validate:"required"`
	GoodsTotal   int    `json:"goods_total"validate:"required"`
	CustomFee    int    `json:"custom_fee"validate:"required"`
}

type Order struct {
	Id                int       `json:"id,omitempty"`
	OrderUid          string    `json:"order_uid"validate:"required"`
	TrackNumber       string    `json:"track_number"validate:"required"`
	Entry             string    `json:"entry"validate:"required"`
	Delivery          *Delivery `json:"delivery"validate:"required"`
	Payment           *Payment  `json:"payment"validate:"required"`
	Items             []Item    `json:"items"validate:"required"`
	Locale            string    `json:"locale"validate:"required"`
	InternalSignature string    `json:"internal_signature"validate:"required"`
	CustomerId        string    `json:"customer_id"validate:"required"`
	DeliveryService   string    `json:"delivery_service"validate:"required"`
	ShardKey          string    `json:"shardkey"validate:"required"`
	SmId              int       `json:"sm_id"validate:"required"`
	DateCreated       time.Time `json:"date_created"validate:"required"`
	OofShard          string    `json:"oof_shard"validate:"required"`
}
