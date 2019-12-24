package mango

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/entropyx/mango/options"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	options.Client
	Address  string
	Port     uint
	Srv      bool
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
	scheme := "mongodb"
	if config.Srv {
		scheme += "+srv"
	}
	u := url.URL{
		Scheme:   scheme,
		User:     url.UserPassword(config.Username, config.Password),
		Host:     hostString(config.Address, config.Port),
		Path:     config.Database,
		RawQuery: "authMechanism=SCRAM-SHA-1&authSource=" + config.Source,
	}
	if config.Context == nil {
		config.Context = context.Background()
	}
	config.ClientOptions.ApplyURI(u.String())
	client, err := mongo.Connect(config.Context, &config.ClientOptions)
	if err != nil {
		return nil, err
	}
	connection := &Connection{client: client, database: config.Database}
	return connection, nil
}

func (c *Connection) Register(ctx context.Context, models ...interface{}) {
	newCtx := context.WithValue(ctx, keyConnection, c)
	for _, model := range models {
		doc := getDocument(model)
		doc.Context = newCtx
	}
}

func (c *Connection) GetClient() *mongo.Client {
	return c.client
}

func (c *Connection) collection(model interface{}) *mongo.Collection {
	client := c.client
	return client.Database(c.database).Collection(getCollection(model))
}

func hostString(address string, port uint) string {
	return fmt.Sprintf("%s%s", address, portString(port))
}

func portString(port uint) string {
	s := ":%d"
	if port > 0 {
		return fmt.Sprintf(s, port)
	}
	return ""
}
