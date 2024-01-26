package mongo

import (
	"context"
	"go-stock-exchange-shares-control/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepository struct {
	database   *mongo.Database
	collection string
}

func NewTransactionRepositoryMongo(connectionString string, databaseName string, collection string) domain.TransactionRepository {

	cl := NewMongoDatabase(connectionString)

	db := cl.Database(databaseName)

	return &transactionRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *transactionRepository) New(c context.Context, transaction *domain.Transaction) (string, error) {
	collection := tr.database.Collection(tr.collection)

	result, err := collection.InsertOne(c, transaction)

	return result.InsertedID.(primitive.ObjectID).String(), err
}

func (tr *transactionRepository) Get(c context.Context, code string) ([]domain.Transaction, error) {
	collection := tr.database.Collection(tr.collection)

	filter := bson.D{{Key: "code", Value: code}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	err = cursor.All(c, &transactions)

	return transactions, err
}
