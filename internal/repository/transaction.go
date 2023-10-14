package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/entities"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionRepo struct {
	mongo database.MongoDB
	name  string
}

type TransactionRepository interface {
	Insert(entities.Transaction) (entities.Transaction, error)
}

func NewTransactionRepository(mongo database.MongoDB) TransactionRepository {
	return &transactionRepo{
		mongo: mongo,
		name:  "Transaction Repository",
	}
}

func (t *transactionRepo) Insert(tx entities.Transaction) (entities.Transaction, error) {
	logrus.Info(fmt.Sprintf("[%s][Insert] is executed", t.name))

	tx.TransactionID = primitive.NewObjectID()
	tx.TransactionOrder = 1
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	db, client, err := t.mongo.Database()
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Insert] %s", t.name, err.Error()))
		return tx, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection("transaction").InsertOne(ctx, tx)
	defer t.mongo.Close(ctx, client)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Insert] %s", t.name, err.Error()))
		return tx, err
	}

	return tx, nil
}
