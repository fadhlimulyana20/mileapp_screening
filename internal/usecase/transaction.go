package usecase

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/appctx"
	"github.com/mileapp_screening/internal/entities"
	"github.com/mileapp_screening/internal/params"
	"github.com/mileapp_screening/internal/repository"
	"github.com/sirupsen/logrus"
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
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithMessage("Package has been deleted")
}
