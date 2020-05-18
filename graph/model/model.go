package model

import (
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MarshalID redefine the base ID type to use an id from an external library
func MarshalID(id primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.Hex()))
	})
}

// UnmarshalID returns primitive.ObjectID and error
func UnmarshalID(v interface{}) (primitive.ObjectID, error) {
	str, ok := v.(string)
	if !ok {
		return primitive.ObjectIDFromHex("")
	}
	return primitive.ObjectIDFromHex(str)
}
