package handler

import (
	"fmt"
	"net/http"
	"time"

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
}

func NewTransactionHandler(mongo database.MongoDB) TransactionHandler {
	return &transactionHandler{
		trxUsecase: usecase.NewTransactionUsecase(mongo),
		name:       "Transaction Handler",
	}
}

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