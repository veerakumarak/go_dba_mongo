package mongo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrSerialiseObjectIdToString = errors.New("serialise: unable to convert objectId to string")
	ErrSerialiseStringToObjectId = errors.New("serialise: unable to convert string to objectId")
	ErrSerialiseEntityToDoc      = errors.New("serialise: unable to convert entity to doc")
	ErrSerialiseDocToEntity      = errors.New("serialise: unable to convert doc to entity")
)

func convertStringToId(hexId string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hexId)
}

func convertIdToString(id primitive.ObjectID) (string, error) {
	return id.Hex(), nil
}
