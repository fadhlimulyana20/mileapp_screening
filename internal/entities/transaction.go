package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrentLocation struct {
	Name string `bson:"name" json:"name"`
	Code string `bson:"code" json:"code"`
	Type string `bson:"type" json:"type"`
}

type TransactionCustomer struct {
	CustomerName          string `bson:"customer_name" json:"customer_name"`
	CustomerAddress       string `bson:"customer_address" json:"customer_address"`
	CustomerEmail         string `bson:"customer_email" json:"customer_email"`
	CustomerPhone         string `bson:"customer_phone" json:"customer_phone"`
	CustomerAddressDetail string `bson:"customer_address_detail" json:"customer_address_detail"`
	CustomerZipCode       string `bson:"customer_zip_code" json:"customer_zip_code"`
	ZoneCode              string `bson:"zone_code" json:"zone_code"`
	OrganizationID        int    `bson:"organization_id" json:"organization_id"`
	LocationID            string `bson:"location_id" json:"location_id"`
}

type Transaction struct {
	TransactionID              primitive.ObjectID  `bson:"_id" json:"transaction_id"`
	CustomerName               string              `bson:"customer_name" json:"customer_name"`
	CustomerCode               string              `bson:"customer_code" json:"customer_code"`
	TransactionAmount          string              `bson:"transaction_amount" json:"transaction_amount"`
	TransactionDiscount        string              `bson:"transaction_discount" json:"transaction_discount"`
	TransactionAdditionalField primitive.M         `bson:"transaction_additional_field" json:"transaction_additional_field"`
	TransactionPaymentType     string              `bson:"transaction_payment_type"`
	TransactionState           string              `bson:"transaction_state" json:"transaction_state"`
	TransactionCode            string              `bson:"transaction_code" json:"transaction_code"`
	TransactionOrder           int                 `bson:"transaction_order" json:"transaction_order"`
	LocationID                 primitive.ObjectID  `bson:"location_id" json:"location_id"`
	OrganizationID             int                 `bson:"organization_id" json:"organization_id"`
	CreatedAt                  time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt                  time.Time           `bson:"updated_at" json:"updated_at"`
	TransactionPaymentTypeName string              `bson:"transaction_payment_type_name"`
	TransactionCashAmount      int                 `bson:"transaction_cash_amount"`
	TransactionCashChange      int                 `bson:"transaction_cash_change"`
	CustomerAttribute          primitive.M         `bson:"customer_attribute" json:"customer_attribute"`
	ConnoteID                  primitive.ObjectID  `bson:"connote_id" json:"connote_id"`
	OriginData                 TransactionCustomer `bson:"origin_data" json:"origin_data"`
	DestinationData            TransactionCustomer `bson:"destination_data" json:"destination_data"`
	KoliData                   []KoliData          `bson:"koli_data" json:"koli_data"`
	CustomField                primitive.M         `bson:"custom_field" json:"custom_field"`
	CurrentLocation            CurrentLocation     `bson:"current_location" json:"current_location"`
}
