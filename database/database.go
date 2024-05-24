package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/anhhuy1010/cms-order/config"
)

var db *mongo.Database

func Init() (*mongo.Database, error) {
	if db == nil {
		cfg := config.GetConfig()
		host := cfg.GetString("database.host")
		port := cfg.GetString("database.port")
		user := cfg.GetString("database.username")
		password := cfg.GetString("database.password")
		database := cfg.GetString("database.db_name")
		ssl := cfg.GetBool("database.ssl")

		var uri string
		if ssl == true {
			uri = fmt.Sprintf(`mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&readPreference=secondaryPreferred`, user, password, host)
		} else {
			uri = fmt.Sprintf(`mongodb://%s:%s@%s:%s/?authMechanism=SCRAM-SHA-256`, user, password, host, port)
		}

		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()
		optionsClient := options.Client()
		optionsClient.ApplyURI(uri)
		client, err := mongo.Connect(ctx, optionsClient)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		db = client.Database(database)
	}

	return db, nil
}

func GetInstance() *mongo.Database {
	return db
}
