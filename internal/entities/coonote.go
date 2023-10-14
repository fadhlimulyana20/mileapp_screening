package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connote struct {
	ConnoteID              primitive.ObjectID `bson:"connote_id" json:"connote_id"`
	ConnoteNumber          int                `bson:"connote_number" json:"connote_number"`
	ConnoteService         string             `bson:"connote_service" json:"connote_service"`
	ConnoteServicePrice    int                `bson:"connote_service_price" json:"connote_service_price"`
	ConnoteAmount          int                `bson:"connote_amount" json:"connote_amount"`
	ConnoteCode            string             `bson:"connote_count" json:"connote_count"`
	ConnoteBookingCode     string             `bson:"connote_booking_code"`
	ConnoteOrder           int                `bson:"connote_order" json:"connote_order"`
	ConnoteState           string             `bson:"connote_state" json:"connote_state"`
	ConnoteStateID         int                `bson:"connote_state_id" json:"connote_state_id"`
	ZoneCodeFrom           string             `bson:"zone_code_from" json:"zone_code_from"`
	ZoneCodeTo             string             `bson:"zone_code_to" json:"zone_code_to"`
	SurchargeAmount        int                `bson:"surcharge_amount" json:"surcharge_amount"`
	TransactionID          primitive.ObjectID `bson:"transaction_id" json:"transaction_id"`
	ActualWeight           int                `bson:"actual_weight" json:"actual_weight"`
	VolumeWeight           int                `bson:"volume_weight" json:"volume_weight"`
	ChargeableWeight       int                `bson:"chargeable_weight" json:"chargeable_weight"`
	CreatedAt              time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt              time.Time          `bson:"updated_at" json:"updated_at"`
	OrganizationID         int                `bson:"organization_id" json:"organization_id"`
	LocationID             string             `bson:"location_id" json:"location_id"`
	ConnoteTotalPackage    string             `bson:"connote_total_pacakage" json:"connote_total_package"`
	ConnoteSurchargeAmount string             `bson:"connote_surcharge_amount" json:"connote_surcharge_amount"`
	ConnoteSlaDay          string             `bson:"connote_sla_day" json:"connote_sla_day"`
	LocationName           string             `bson:"location_name" json:"location_name"`
	LocationType           string             `bson:"location_type" json:"location_type"`
	SourceTarifDB          string             `bson:"source_tarif_db" json:"source_tarif_db"`
	IdSourceTarif          string             `bson:"id_source_tarif" json:"id_source_tarif"`
	Pod                    string             `bson:"pod" json:"pod"`
	History                primitive.A        `bson:"history" json:"history"`
}
