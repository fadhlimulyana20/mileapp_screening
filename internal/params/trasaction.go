package params

type TransactionCurrentLocationCreateParam struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	Type string `json:"type" validate:"required"`
}

type TransactionCurrentLocationPatchParam struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`
}

type TransactionCustomerCreateParam struct {
	CustomerName          string `json:"customer_name"`
	CustomerAddress       string `json:"customer_address"`
	CustomerEmail         string `json:"customer_email"`
	CustomerPhone         string `json:"customer_phone"`
	CustomerAddressDetail string `json:"customer_address_detail"`
	CustomerZipCode       string `json:"customer_zip_code"`
	ZoneCode              string `json:"zone_code"`
	OrganizationID        int    `json:"organization_id"`
	LocationID            string `json:"location_id"`
}

type TransactionCreateParam struct {
	CustomerName               string                                `json:"customer_name" validate:"required"`
	CustomerCode               string                                `json:"customer_code" validate:"required"`
	TransactionAmount          string                                `json:"transaction_amount" validate:"required"`
	TransactionDiscount        string                                `json:"transaction_discount" validate:"required"`
	TransactionAdditionalField map[string]interface{}                `json:"transaction_additional_field"`
	TransactionPaymentType     string                                `json:"transaction_payment_type" validate:"required"`
	TransactionState           string                                `json:"transaction_state" validate:"required"`
	LocationID                 string                                `json:"location_id" validate:"required"`
	OrganizationID             int                                   `json:"organization_id" validate:"required"`
	TransactionPaymentTypeName string                                `json:"transaction_payment_type_name" validate:"required"`
	TransactionCashAmount      *int                                  `json:"transaction_cash_amount" validate:"required"`
	TransactionCashChange      *int                                  `json:"transaction_cash_change" validate:"required"`
	CustomerAttribute          map[string]interface{}                `json:"customer_attribute"`
	Connote                    ConnoteCreateParam                    `json:"connote" validate:"required"`
	ConnoteID                  string                                `json:"connote_id" validate:"required"`
	OriginData                 TransactionCustomerCreateParam        `json:"origin_data" validate:"required"`
	DestinationData            TransactionCustomerCreateParam        `json:"destination_data" validate:"required"`
	KoliData                   []KoliDataCreateParam                 `json:"koli_data" validate:"required"`
	CustomField                map[string]interface{}                `json:"custom_field"`
	CurrentLocation            TransactionCurrentLocationCreateParam `json:"current_location" validate:"required"`
}

type TransactionUpdateParam struct {
	TransactionID              string                                `json:"transaction_id"`
	CustomerName               string                                `json:"customer_name" validate:"required"`
	CustomerCode               string                                `json:"customer_code" validate:"required"`
	TransactionAmount          string                                `json:"transaction_amount" validate:"required"`
	TransactionDiscount        string                                `json:"transaction_discount" validate:"required"`
	TransactionAdditionalField map[string]interface{}                `json:"transaction_additional_field"`
	TransactionPaymentType     string                                `json:"transaction_payment_type" validate:"required"`
	TransactionState           string                                `json:"transaction_state" validate:"required"`
	LocationID                 string                                `json:"location_id" validate:"required"`
	OrganizationID             int                                   `json:"organization_id" validate:"required"`
	TransactionPaymentTypeName string                                `json:"transaction_payment_type_name" validate:"required"`
	TransactionCashAmount      *int                                  `json:"transaction_cash_amount" validate:"required"`
	TransactionCashChange      *int                                  `json:"transaction_cash_change" validate:"required"`
	CustomerAttribute          map[string]interface{}                `json:"customer_attribute"`
	Connote                    ConnoteUpdateParam                    `json:"connote"`
	ConnoteID                  string                                `json:"connote_id"`
	OriginData                 TransactionCustomerCreateParam        `json:"origin_data"`
	DestinationData            TransactionCustomerCreateParam        `json:"destination_data"`
	KoliData                   []KoliDataUpdateParam                 `json:"koli_data"`
	CustomField                map[string]interface{}                `json:"custom_field"`
	CurrentLocation            TransactionCurrentLocationCreateParam `json:"current_location" validate:"required"`
}

type TransactionPatchParam struct {
	TransactionID              string                               `json:"transaction_id"`
	CustomerName               string                               `json:"customer_name"`
	CustomerCode               string                               `json:"customer_code"`
	TransactionAmount          string                               `json:"transaction_amount"`
	TransactionDiscount        string                               `json:"transaction_discount"`
	TransactionAdditionalField map[string]interface{}               `json:"transaction_additional_field"`
	TransactionPaymentType     string                               `json:"transaction_payment_type"`
	TransactionState           string                               `json:"transaction_state"`
	LocationID                 string                               `json:"location_id"`
	OrganizationID             int                                  `json:"organization_id"`
	TransactionPaymentTypeName string                               `json:"transaction_payment_type_name"`
	TransactionCashAmount      *int                                 `json:"transaction_cash_amount"`
	TransactionCashChange      *int                                 `json:"transaction_cash_change"`
	CustomerAttribute          map[string]interface{}               `json:"customer_attribute"`
	Connote                    ConnoteUpdateParam                   `json:"connote"`
	ConnoteID                  string                               `json:"connote_id"`
	OriginData                 TransactionCustomerCreateParam       `json:"origin_data"`
	DestinationData            TransactionCustomerCreateParam       `json:"destination_data"`
	KoliData                   []KoliDataUpdateParam                `json:"koli_data"`
	CustomField                map[string]interface{}               `json:"custom_field"`
	CurrentLocation            TransactionCurrentLocationPatchParam `json:"current_location"`
}
