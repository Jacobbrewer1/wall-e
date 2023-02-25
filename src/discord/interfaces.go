package discord

import (
	"database/sql"
	"encoding"
	"go.mongodb.org/mongo-driver/bson"
	"go/types"
)

type (
	iEnum interface {
		encoding.TextMarshaler
		encoding.TextUnmarshaler
		bson.Marshaler
		bson.Unmarshaler
		sql.Scanner
		types.Type
		Display() string
	}
)
