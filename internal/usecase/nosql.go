package usecase

import (
	"fmt"
	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/appctx"
	"github.com/mileapp_screening/internal/entities"
	"github.com/mileapp_screening/internal/params"
	"github.com/mileapp_screening/internal/repository"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type nosql struct {
	repo repository.NoSQLInf
	name string
}

type NosqlUsecase interface {
	Insert(param params.CreateNosqlParam) appctx.Response
}

func NewNoSqlUsecase(mongo database.MongoDB) NosqlUsecase {
	return &nosql{
		repo: repository.NewNoSQLInf(mongo),
		name: "NoSQLUsecase",
	}
}

func (n *nosql) Insert(param params.CreateNosqlParam) appctx.Response {
	var nosqle entities.Nosql
	copier.Copy(&nosqle, &param)

	nosqle, err := n.repo.InsertOne(nosqle)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Insert] %s", n.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(nosqle)
}
