package mango

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Operator struct {
	Operator string
	Value    interface{}
}

func Set(model interface{}) *Operator {
	return &Operator{
		Operator: "set",
		Value:    model,
	}
}

func (o *Operator) apply() bson.D {
	fmt.Println("value", toBsonDoc(o.Value))
	return bson.D{{fmt.Sprintf("$%s", o.Operator), toBsonDoc(o.Value)}}
}
