package driver

import (
	"context"
	"crawl/conf"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongo() *MongoDB {
	config, err := conf.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.duth3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		config.DBUser, config.DBPassword)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary()) // ping vào bản chính
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
	Mongo.Client = client

	return Mongo
}
