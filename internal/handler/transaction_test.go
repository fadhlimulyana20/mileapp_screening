package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mileapp_screening/database"
)

var (
	mongoHost     = "localhost"
	mongoPort     = 27017
	mongoUser     = "root"
	mongoPassword = "password"
	mongoDatabase = "test_go_kit"
)

func TestGetListTransaction(t *testing.T) {
	req, err := http.NewRequest("GET", "/package", nil)
	if err != nil {
		t.Fatal(err)
	}

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.Get)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}

func TestGetDetailTransaction(t *testing.T) {
	req, err := http.NewRequest("GET", "/package/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "652a5a45d801013fa6cdc918")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.GetOne)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusOK, http.StatusNotFound)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}

func TestCreateTransaction(t *testing.T) {
	jsond := `
	{
		"customer_name": "agus",
		"customer_code": "24a",
		"transaction_amount": "70000",
		"transaction_discount": "0",
		"transaction_additional_field": {},
		"transaction_payment_type": "29",
		"transaction_state": "PAID",
		"location_id": "123",
		"organization_id": 25,
		"transaction_payment_type_name": "invoice",
		"transaction_cash_amount": 0,
		"transaction_cash_change": 0,
		"customer_attribute": {
			"Nama_Sales": "Radit Fitrawikarsa",
			"TOP": "14 Hari",
			"Jenis_Pelanggan": "B2B"
		},
		"connote": {
			"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
			"connote_number": 1,
			"connote_service": "ECO",
			"connote_service_price": 70700,
			"connote_amount": 70700,
			"connote_code": "AWB00100209082020",
			"connote_booking_code": "",
			"connote_order": 326931,
			"connote_state": "PAID",
			"connote_state_id": 2,
			"zone_code_from": "CGKFT",
			"zone_code_to": "SMG",
			"surcharge_amount": null,
			"transaction_id": "d0090c40-539f-479a-8274-899b9970bddc",
			"actual_weight": 20,
			"volume_weight": 0,
			"chargeable_weight": 20,
			"created_at": "2020-07-15T11:11:12+0700",
			"updated_at": "2020-07-15T11:11:22+0700",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74",
			"connote_total_package": "3",
			"connote_surcharge_amount": "0",
			"connote_sla_day": "4",
			"location_name": "Hub Jakarta Selatan",
			"location_type": "HUB",
			"source_tariff_db": "tariff_customers",
			"id_source_tariff": "1576868",
			"pod": null,
			"history": []
		},
		"connote_id": "123",
		"origin_data": {
			"customer_name": "PT. NARA OKA PRAKARSA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 100, SEMARANG TENGAH 12420",
			"customer_email": "info@naraoka.co.id",
			"customer_phone": "024-1234567",
			"customer_address_detail": null,
			"customer_zip_code": "12420",
			"zone_code": "CGKFT",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"destination_data": {
			"customer_name": "PT AMARIS HOTEL SIMPANG LIMA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 01, SEMARANG TENGAH",
			"customer_email": null,
			"customer_phone": "0248453499",
			"customer_address_detail": "KOTA SEMARANG SEMARANG TENGAH KARANGKIDUL",
			"customer_zip_code": "50241",
			"zone_code": "SMG",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"koli_data": [
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.1",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "e2cb6d86-0bb9-409b-a1f0-389ed4f2df2d",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.1"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.2",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "3600f10b-4144-4e58-a024-cc3178e7a709",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.2"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.3",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 2,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "LID HOT CUP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 2,
				"koli_id": "2937bdbf-315e-4c5e-b139-fd39a3dfd15f",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.3"
			}
		],
		"custom_field": {
			"catatan_tambahan": "JANGAN DI BANTING \/ DI TINDIH"
		},
		"current_location": {
			"name": "Hub Jakarta Selatan",
			"code": "JKTS01",
			"type": "Agent"
		}
	}
	`

	var bodyJson map[string]interface{}
	json.Unmarshal([]byte(jsond), &bodyJson)
	body, _ := json.Marshal(bodyJson)

	req, err := http.NewRequest("POST", "/package/", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.Insert)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusOK, http.StatusNotFound)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}

func TestUpdateTransaction(t *testing.T) {
	jsond := `
	{
		"customer_name": "agus",
		"customer_code": "24a",
		"transaction_amount": "70000",
		"transaction_discount": "0",
		"transaction_additional_field": {},
		"transaction_payment_type": "29",
		"transaction_state": "PAID",
		"location_id": "123",
		"organization_id": 25,
		"transaction_payment_type_name": "invoice",
		"transaction_cash_amount": 0,
		"transaction_cash_change": 0,
		"customer_attribute": {
			"Nama_Sales": "Radit Fitrawikarsa",
			"TOP": "14 Hari",
			"Jenis_Pelanggan": "B2B"
		},
		"connote": {
			"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
			"connote_number": 1,
			"connote_service": "ECO",
			"connote_service_price": 70700,
			"connote_amount": 70700,
			"connote_code": "AWB00100209082020",
			"connote_booking_code": "",
			"connote_order": 326931,
			"connote_state": "PAID",
			"connote_state_id": 2,
			"zone_code_from": "CGKFT",
			"zone_code_to": "SMG",
			"surcharge_amount": null,
			"transaction_id": "d0090c40-539f-479a-8274-899b9970bddc",
			"actual_weight": 20,
			"volume_weight": 0,
			"chargeable_weight": 20,
			"created_at": "2020-07-15T11:11:12+0700",
			"updated_at": "2020-07-15T11:11:22+0700",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74",
			"connote_total_package": "3",
			"connote_surcharge_amount": "0",
			"connote_sla_day": "4",
			"location_name": "Hub Jakarta Selatan",
			"location_type": "HUB",
			"source_tariff_db": "tariff_customers",
			"id_source_tariff": "1576868",
			"pod": null,
			"history": []
		},
		"connote_id": "123",
		"origin_data": {
			"customer_name": "PT. NARA OKA PRAKARSA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 100, SEMARANG TENGAH 12420",
			"customer_email": "info@naraoka.co.id",
			"customer_phone": "024-1234567",
			"customer_address_detail": null,
			"customer_zip_code": "12420",
			"zone_code": "CGKFT",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"destination_data": {
			"customer_name": "PT AMARIS HOTEL SIMPANG LIMA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 01, SEMARANG TENGAH",
			"customer_email": null,
			"customer_phone": "0248453499",
			"customer_address_detail": "KOTA SEMARANG SEMARANG TENGAH KARANGKIDUL",
			"customer_zip_code": "50241",
			"zone_code": "SMG",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"koli_data": [
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.1",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "e2cb6d86-0bb9-409b-a1f0-389ed4f2df2d",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.1"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.2",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "3600f10b-4144-4e58-a024-cc3178e7a709",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.2"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.3",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 2,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "LID HOT CUP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 2,
				"koli_id": "2937bdbf-315e-4c5e-b139-fd39a3dfd15f",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.3"
			}
		],
		"custom_field": {
			"catatan_tambahan": "JANGAN DI BANTING \/ DI TINDIH"
		},
		"current_location": {
			"name": "Hub Jakarta Selatan",
			"code": "JKTS01",
			"type": "Agent"
		}
	}
	`

	var bodyJson map[string]interface{}
	json.Unmarshal([]byte(jsond), &bodyJson)
	body, _ := json.Marshal(bodyJson)

	req, err := http.NewRequest("PUT", "/package/{id}", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "652a5a45d801013fa6cdc918")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.Update)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusOK, http.StatusNotFound)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}

func TestPatchTransaction(t *testing.T) {
	jsond := `
	{
		"customer_code": "24a",
		"transaction_amount": "70000",
		"transaction_discount": "0",
		"transaction_additional_field": {},
		"transaction_payment_type": "29",
		"transaction_state": "PAID",
		"location_id": "123",
		"organization_id": 25,
		"transaction_payment_type_name": "invoice",
		"transaction_cash_amount": 0,
		"transaction_cash_change": 0,
		"customer_attribute": {
			"Nama_Sales": "Radit Fitrawikarsa",
			"TOP": "14 Hari",
			"Jenis_Pelanggan": "B2B"
		},
		"connote": {
			"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
			"connote_number": 1,
			"connote_service": "ECO",
			"connote_service_price": 70700,
			"connote_amount": 70700,
			"connote_code": "AWB00100209082020",
			"connote_booking_code": "",
			"connote_order": 326931,
			"connote_state": "PAID",
			"connote_state_id": 2,
			"zone_code_from": "CGKFT",
			"zone_code_to": "SMG",
			"surcharge_amount": null,
			"transaction_id": "d0090c40-539f-479a-8274-899b9970bddc",
			"actual_weight": 20,
			"volume_weight": 0,
			"chargeable_weight": 20,
			"created_at": "2020-07-15T11:11:12+0700",
			"updated_at": "2020-07-15T11:11:22+0700",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74",
			"connote_total_package": "3",
			"connote_surcharge_amount": "0",
			"connote_sla_day": "4",
			"location_name": "Hub Jakarta Selatan",
			"location_type": "HUB",
			"source_tariff_db": "tariff_customers",
			"id_source_tariff": "1576868",
			"pod": null,
			"history": []
		},
		"connote_id": "123",
		"origin_data": {
			"customer_name": "PT. NARA OKA PRAKARSA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 100, SEMARANG TENGAH 12420",
			"customer_email": "info@naraoka.co.id",
			"customer_phone": "024-1234567",
			"customer_address_detail": null,
			"customer_zip_code": "12420",
			"zone_code": "CGKFT",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"destination_data": {
			"customer_name": "PT AMARIS HOTEL SIMPANG LIMA",
			"customer_address": "JL. KH. AHMAD DAHLAN NO. 01, SEMARANG TENGAH",
			"customer_email": null,
			"customer_phone": "0248453499",
			"customer_address_detail": "KOTA SEMARANG SEMARANG TENGAH KARANGKIDUL",
			"customer_zip_code": "50241",
			"zone_code": "SMG",
			"organization_id": 6,
			"location_id": "5cecb20b6c49615b174c3e74"
		},
		"koli_data": [
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.1",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "e2cb6d86-0bb9-409b-a1f0-389ed4f2df2d",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.1"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.2",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 9,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "V WARP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 9,
				"koli_id": "3600f10b-4144-4e58-a024-cc3178e7a709",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.2"
			},
			{
				"koli_length": 0,
				"awb_url": "https:\/\/tracking.mile.app\/label\/AWB00100209082020.3",
				"created_at": "2020-07-15 11:11:13",
				"koli_chargeable_weight": 2,
				"koli_width": 0,
				"koli_surcharge": [],
				"koli_height": 0,
				"updated_at": "2020-07-15 11:11:13",
				"koli_description": "LID HOT CUP",
				"koli_formula_id": null,
				"connote_id": "f70670b1-c3ef-4caf-bc4f-eefa702092ed",
				"koli_volume": 0,
				"koli_weight": 2,
				"koli_id": "2937bdbf-315e-4c5e-b139-fd39a3dfd15f",
				"koli_custom_field": {
					"awb_sicepat": null,
					"harga_barang": null
				},
				"koli_code": "AWB00100209082020.3"
			}
		],
		"custom_field": {
			"catatan_tambahan": "JANGAN DI BANTING \/ DI TINDIH"
		},
		"current_location": {
			"name": "Hub Jakarta Selatan",
			"code": "JKTS01",
			"type": "Agent"
		}
	}
	`

	var bodyJson map[string]interface{}
	json.Unmarshal([]byte(jsond), &bodyJson)
	body, _ := json.Marshal(bodyJson)

	req, err := http.NewRequest("PATCH", "/package/{id}", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "652a5a45d801013fa6cdc918")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.Patch)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusOK, http.StatusNotFound)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}

func TestGetDeleteTransaction(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/package/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "652a5a45d801013fa6cdc918")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	mongoDB := database.NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)

	rr := httptest.NewRecorder()
	trxHandler := NewTransactionHandler(mongoDB)
	handler := http.HandlerFunc(trxHandler.Delete)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v or %v",
			status, http.StatusOK, http.StatusNotFound)
	}

	if rr.Body.String() == "" {
		t.Errorf("empty body string")
	}
}
