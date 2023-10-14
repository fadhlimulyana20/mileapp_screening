package params

type ConnoteCreateParam struct {
	ConnoteNumber          int           `json:"connote_number"`
	ConnoteService         string        `json:"connote_service"`
	ConnoteServicePrice    int           `json:"connote_service_price"`
	ConnoteAmount          int           `json:"connote_amount"`
	ConnoteCode            string        `json:"connote_count"`
	ConnoteBookingCode     string        `json:"connote_booking_code"`
	ConnoteOrder           int           `json:"connote_order"`
	ConnoteState           string        `json:"connote_state"`
	ConnoteStateID         int           `json:"connote_state_id"`
	ZoneCodeFrom           string        `json:"zone_code_from"`
	ZoneCodeTo             string        `json:"zone_code_to"`
	SurchargeAmount        int           `json:"surcharge_amount"`
	TransactionID          string        `json:"transaction_id"`
	ActualWeight           int           `json:"actual_weight"`
	VolumeWeight           int           `json:"volume_weight"`
	ChargeableWeight       int           `json:"chargeable_weight"`
	OrganizationID         int           `json:"organization_id"`
	LocationID             string        `json:"location_id"`
	ConnoteTotalPackage    string        `json:"connote_total_package"`
	ConnoteSurchargeAmount string        `json:"connote_surcharge_amount"`
	ConnoteSlaDay          string        `json:"connote_sla_day"`
	LocationName           string        `json:"location_name"`
	LocationType           string        `json:"location_type"`
	SourceTarifDB          string        `json:"source_tarif_db"`
	IdSourceTarif          string        `json:"id_source_tarif"`
	Pod                    string        `json:"pod"`
	History                []interface{} `json:"history"`
}
