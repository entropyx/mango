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
	client   *mongo.Client
	database string
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
	connection := &Connection{client: client, database: config.Database}
	return connection, nil
}

func (c *Connection) UpdateOne(filter interface{}, model interface{}) error {
	bsonDoc := toBsonDoc(model)
	doc, err := getDocument(model)
	if err != nil {
		return err
	}
	collection := c.collection(model)
	collection.UpdateOne(ctx, filter, update)
}

func (c *Connection) collection(model interface{}) *mongo.Collection {
	client := c.client
	return client.Database(c.database).Collection(getCollection(model))
}
