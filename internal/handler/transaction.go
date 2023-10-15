package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/appctx"
	"github.com/mileapp_screening/internal/params"
	"github.com/mileapp_screening/internal/usecase"
	"github.com/mileapp_screening/utils/json"
	"github.com/mileapp_screening/utils/validator"

	"github.com/sirupsen/logrus"
)

type transactionHandler struct {
	trxUsecase usecase.TransactionUsecase
	name       string
	handler    Handler
}

type TransactionHandler interface {
	Insert(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
}

func NewTransactionHandler(mongo database.MongoDB) TransactionHandler {
	return &transactionHandler{
		trxUsecase: usecase.NewTransactionUsecase(mongo),
		name:       "Transaction Handler",
	}
}

// Create Package godoc
// @Summary      Create Package
// @Description  Create Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Param        transaction   body  params.TransactionCreateParam  true  "Create Package"
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/ [post]
func (t *transactionHandler) Insert(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Insert] is executed", t.name))
	startTime := time.Now()

	var param params.TransactionCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	logrus.Debug(param)

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := t.trxUsecase.Insert(param)
	t.handler.Response(w, resp, startTime, time.Now())
}

// Get Package godoc
// @Summary      Get Package
// @Description  Get Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Package id"
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/{id} [get]
func (t *transactionHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Get One] is executed", t.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")

	resp := t.trxUsecase.GetOne(id)
	t.handler.Response(w, resp, startTime, time.Now())
}

// Get List Package godoc
// @Summary      Get List Package
// @Description  Get List Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/ [get]
func (t *transactionHandler) Get(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Get] is executed", t.name))
	startTime := time.Now()

	resp := t.trxUsecase.Get()
	t.handler.Response(w, resp, startTime, time.Now())
}

// Delete Package godoc
// @Summary      Delete Package
// @Description  Delete Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Package id"
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/{id} [delete]
func (t *transactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Delete] is executed", t.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")

	resp := t.trxUsecase.Delete(id)
	t.handler.Response(w, resp, startTime, time.Now())
}

// Update Package godoc
// @Summary      Update Package
// @Description  Update Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Package ID"
// @Param        transaction   body  params.TransactionUpdateParam  true  "Update Package"
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/{id} [put]
func (t *transactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Update] is executed", t.name))
	startTime := time.Now()

	var param params.TransactionUpdateParam
	id := chi.URLParam(r, "id")
	param.TransactionID = id

	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	logrus.Debug(param)

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := t.trxUsecase.Update(param)
	t.handler.Response(w, resp, startTime, time.Now())
}

// Patch Package godoc
// @Summary      Patch Package
// @Description  Patch Package
// @Tags         Package
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Package ID"
// @Param        transaction   body  params.TransactionPatchParam  true  "Patch Package"
// @Success      200  {object}  interface{}
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /package/{id} [patch]
func (t *transactionHandler) Patch(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Patch] is executed", t.name))
	startTime := time.Now()

	var param params.TransactionPatchParam
	id := chi.URLParam(r, "id")
	param.TransactionID = id

	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	logrus.Debug(param)

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		t.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := t.trxUsecase.Patch(param)
	t.handler.Response(w, resp, startTime, time.Now())
}
