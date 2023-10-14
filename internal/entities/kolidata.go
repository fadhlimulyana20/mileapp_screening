package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KoliData struct {
	KoliLength           int                `bson:"koli_length" json:"koli_length"`
	AWBUrl               string             `bson:"awb_url" json:"awb_url"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
	KoliChargeableWeight int                `bson:"koli_chargeable_weight" json:"koli_chargeable_weight"`
	KoliWidth            int                `bson:"koli_width" json:"koli_width"`
	KoliSurcharge        bson.A             `bson:"koli_surcharge" json:"koli_surcharge"`
	KoliHeight           int                `bson:"koli_height" json:"koli_height"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at"`
	KoliDescription      string             `bson:"koli_description" json:"koli_description"`
	KoliFormulaID        *int               `bson:"koli_formula_id" json:"koli_formula_id"`
	ConnoteID            primitive.ObjectID `bson:"connote_id" json:"connote_id"`
	KoliVolume           int                `bson:"koli_volume" json:"koli_volume"`
	KoliWeight           int                `bson:"koli_weight" json:"koli_weight"`
	KoliID               primitive.ObjectID `bson:"koli_id" json:"koli_id"`
	KoliCustomField      primitive.M        `bson:"koli_custom_field" json:"koli_custom_field"`
	KoliCode             string             `bson:"koli_code" json:"koli_code"`
}
