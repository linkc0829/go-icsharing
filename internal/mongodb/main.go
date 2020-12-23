package mongodb

import (
	"context"
	"log"

	"github.com/linkc0829/go-ics/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbUser, pwd, dbName, dsn string

type MongoDB struct {
	Session       *mongo.Client
	Users         *mongo.Collection
	Deleted_users *mongo.Collection
	Income        *mongo.Collection
	Cost          *mongo.Collection
	IncomeHistory *mongo.Collection
	CostHistory   *mongo.Collection
}

//ConnectDB will build connection to MongoDB Atlas
func ConnectDB(cfg *utils.ServerConfig) *MongoDB {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDB.DSN))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoDB{
		Session:       client,
		Users:         client.Database("ics").Collection("users"),
		Deleted_users: client.Database("ics").Collection("deleted_users"),
		Income:        client.Database("ics").Collection("income"),
		Cost:          client.Database("ics").Collection("cost"),
		IncomeHistory: client.Database("ics").Collection("incomeHistory"),
		CostHistory:   client.Database("ics").Collection("costHistory"),
	}
}

//CloseDB will dissconnect to MongoDB
func CloseDB(db *MongoDB) {
	err := db.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}