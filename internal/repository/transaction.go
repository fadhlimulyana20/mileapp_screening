package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mileapp_screening/database"
	"github.com/mileapp_screening/internal/entities"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionRepo struct {
	mongo database.MongoDB
	name  string
}

type TransactionRepository interface {
	Insert(entities.Transaction) (entities.Transaction, error)
	GetOne(ID string) (entities.Transaction, error)
	Get() ([]entities.Transaction, int, error)
}

func NewTransactionRepository(mongo database.MongoDB) TransactionRepository {
	return &transactionRepo{
		mongo: mongo,
		name:  "Transaction Repository",
	}
}

func (t *transactionRepo) Insert(tx entities.Transaction) (entities.Transaction, error) {
	logrus.Info(fmt.Sprintf("[%s][Insert] is executed", t.name))

	transactionID := primitive.NewObjectID()
	tx.TransactionID = transactionID
	tx.TransactionOrder = 1
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	connoteID := primitive.NewObjectID()
	tx.Connote.ConnoteID = connoteID
	tx.ConnoteID = connoteID
	tx.Connote.TransactionID = transactionID

	for i, _ := range tx.KoliData {
		tx.KoliData[i].KoliID = primitive.NewObjectID()
		tx.KoliData[i].CreatedAt = time.Now()
		tx.KoliData[i].UpdatedAt = time.Now()
		tx.KoliData[i].ConnoteID = connoteID
	}

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

func (t *transactionRepo) GetOne(ID string) (entities.Transaction, error) {
	logrus.Info(fmt.Sprintf("[%s][Get One] is executed", t.name))
	var tx entities.Transaction

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get One] %s", t.name, err.Error()))
		return tx, err
	}

	db, client, err := t.mongo.Database()
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get One] %s", t.name, err.Error()))
		return tx, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", objID}}
	err = db.Collection("transaction").FindOne(ctx, filter).Decode(&tx)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get One] %s", t.name, err.Error()))
		return tx, err
	}
	defer t.mongo.Close(ctx, client)

	return tx, nil
}

func (t *transactionRepo) Get() ([]entities.Transaction, int, error) {
	logrus.Info(fmt.Sprintf("[%s][Get] is executed", t.name))

	var txs []entities.Transaction

	db, client, err := t.mongo.Database()
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", t.name, err.Error()))
		return txs, 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := db.Collection("transaction").Find(ctx, bson.D{})
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", t.name, err.Error()))
		return txs, 0, err
	}

	defer t.mongo.Close(ctx, client)

	for cur.Next(ctx) {
		var tx entities.Transaction
		err := cur.Decode(&tx)
		if err != nil {
			logrus.Error(fmt.Sprintf("[%s][Get] %s", t.name, err.Error()))
			return txs, 0, err
		}

		txs = append(txs, tx)
	}

	return txs, len(txs), nil
}
