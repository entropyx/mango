package mango

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type Config struct {
	options.ClientOptions
	Address  string
	Port     uint
	Timeout  *time.Time
	Database string
	Username string
	Password string
	Source   string
	Context  context.Context
}

type Connection struct {
	client *mongo.Client
}

func Connect(config *Config) (*Connection, error) {
	u := url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(config.Username, config.Password),
		Host:   fmt.Sprintf("%s:%s", config.Address, config.Port),
	}
	if config.Context == nil {
		config.Context = context.Background()
	}
	client, err := mongo.Connect(config.Context, u.String(), &config.ClientOptions)
	if err != nil {
		return nil, err
	}
	connection := &Connection{client}
	return connection, nil
}

func (c *Connection) UpsertOne(filter interface{}, model interface{}) error {
	bsonDoc := toBsonDoc(model)
	doc, err := getDocument(model)
	if err != nil {
		return err
	}
	client := doc.Context.Value(clientKey).(*mongo.Client)
	client.Database(getCollection(model))
}
