package repo

import (
	"context"
	"crawl/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Db *mongo.Database
}

type IRepo interface {
	CreateCollection(collect *model.Collection) error
	CreateProduct(collect *model.Item) error
}

func NewRepo(db *mongo.Database) IRepo {
	return &Repo{
		Db: db,
	}
}

func (mongo *Repo) CreateCollection(collect *model.Collection) error {
	bbytes, _ := bson.Marshal(collect)

	_, err := mongo.Db.Collection("collections").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *Repo) CreateProduct(product *model.Item) error {
	bbytes, _ := bson.Marshal(product)

	_, err := mongo.Db.Collection("products").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}
