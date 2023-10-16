package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/appctx"
	"github.com/mileapp_screening/internal/entities"
	"github.com/mileapp_screening/internal/params"
	"github.com/mileapp_screening/internal/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionUc struct {
	trxRepo repository.TransactionRepository
	name    string
}

type TransactionUsecase interface {
	Insert(param params.TransactionCreateParam) appctx.Response
	GetOne(ID string) appctx.Response
	Get() appctx.Response
	Delete(ID string) appctx.Response
	Update(param params.TransactionUpdateParam) appctx.Response
	Patch(param params.TransactionPatchParam) appctx.Response
}

func NewTransactionUsecase(mongo database.MongoDB) TransactionUsecase {
	return &transactionUc{
		trxRepo: repository.NewTransactionRepository(mongo),
		name:    "Transaction Usecase",
	}
}

func (t *transactionUc) Insert(param params.TransactionCreateParam) appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Insert] is executed", t.name))

	var trx entities.Transaction
	copier.Copy(&trx, &param)

	trx, err := t.trxRepo.Insert(trx)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Insert] %s", t.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(trx)
}

func (t *transactionUc) GetOne(ID string) appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Get One] is executed", t.name))

	trx, err := t.trxRepo.GetOne(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get One] %s", t.name, err.Error()))
		if errors.Is(err, mongo.ErrNoDocuments) {
			return *appctx.NewResponse().WithCode(http.StatusNotFound).WithMessage("data not found")
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(trx)
}

func (t *transactionUc) Get() appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Get] is executed", t.name))

	txs, count, err := t.trxRepo.Get()
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get One] %s", t.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(txs).WithMeta(1, 25, int64(count))
}

func (t *transactionUc) Delete(ID string) appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Delete] is executed", t.name))

	_, err := t.trxRepo.Delete(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", t.name, err.Error()))
		if errors.Is(err, mongo.ErrNoDocuments) {
			return *appctx.NewResponse().WithCode(http.StatusNotFound).WithMessage("data not found")
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithMessage("Package has been deleted")
}

func (t *transactionUc) Update(param params.TransactionUpdateParam) appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Update] is executed", t.name))

	var trx entities.Transaction
	trx, err := t.trxRepo.GetOne(param.TransactionID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", t.name, err.Error()))
		if errors.Is(err, mongo.ErrNoDocuments) {
			return *appctx.NewResponse().WithCode(http.StatusNotFound).WithMessage("data not found")
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}
	copier.Copy(&trx, &param)

	connoteID, err := primitive.ObjectIDFromHex(param.ConnoteID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", t.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}
	trx.ConnoteID = connoteID
	trx.Connote.ConnoteID = connoteID

	for i, _ := range trx.KoliData {
		koliID, err := primitive.ObjectIDFromHex(param.KoliData[i].KoliID)
		if err != nil {
			logrus.Error(fmt.Sprintf("[%s][Update] %s", t.name, err.Error()))
			return *appctx.NewResponse().WithErrors(err.Error())
		}
		trx.KoliData[i].ConnoteID = connoteID
		trx.KoliData[i].KoliID = koliID
	}

	trx, err = t.trxRepo.Update(trx)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", t.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(trx)
}

func (t *transactionUc) Patch(param params.TransactionPatchParam) appctx.Response {
	logrus.Info(fmt.Sprintf("[%s][Patch] is executed", t.name))

	var trx entities.Transaction
	trx, err := t.trxRepo.GetOne(param.TransactionID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Patch] %s", t.name, err.Error()))
		if errors.Is(err, mongo.ErrNoDocuments) {
			return *appctx.NewResponse().WithCode(http.StatusNotFound).WithMessage("data not found")
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}
	copier.CopyWithOption(&trx, &param, copier.Option{IgnoreEmpty: true})

	trx, err = t.trxRepo.Update(trx)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Patch] %s", t.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(trx)
}
