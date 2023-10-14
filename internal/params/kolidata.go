package params

type KoliDataCreateParam struct {
	KoliLength           int                    `json:"koli_length"`
	AWBUrl               string                 `json:"awb_url"`
	KoliChargeableWeight int                    `json:"koli_chargeable_weight"`
	KoliWidth            int                    `json:"koli_width"`
	KoliSurcharge        []interface{}          `json:"koli_surcharge"`
	KoliHeight           int                    `json:"koli_height"`
	KoliDescription      string                 `json:"koli_description"`
	KoliFormulaID        *int                   `json:"koli_formula_id"`
	ConnoteID            string                 `json:"connote_id"`
	KoliVolume           int                    `bson:"koli_volume" json:"koli_volume"`
	KoliWeight           int                    `bson:"koli_weight" json:"koli_weight"`
	KoliID               string                 `json:"koli_id"`
	KoliCustomField      map[string]interface{} `json:"koli_custom_field"`
	KoliCode             string                 `json:"koli_code"`
}
