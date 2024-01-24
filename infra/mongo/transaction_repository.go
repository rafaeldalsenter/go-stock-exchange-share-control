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

func (tr *transactionRepository) New(transaction *domain.Transaction) error {
	collection := tr.database.Collection(tr.collection)
	c := context.TODO()

	_, err := collection.InsertOne(c, transaction)

	return err
}

func (tr *transactionRepository) Get(code string) ([]domain.Transaction, error) {
	collection := tr.database.Collection(tr.collection)
	c := context.TODO()

	idHex, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(c, bson.M{"code": idHex})
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	err = cursor.All(c, &transactions)

	return transactions, err
}
