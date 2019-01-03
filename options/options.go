package options

import "github.com/mongodb/mongo-go-driver/mongo/options"

type Update struct {
	Upsert bool
}

type Client struct {
	options.ClientOptions
}
