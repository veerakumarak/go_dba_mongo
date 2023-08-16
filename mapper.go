package mongo

import (
	"errors"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrSerialiseObjectIdToString = errors.New("serialise: unable to convert objectId to string")
	ErrSerialiseStringToObjectId = errors.New("serialise: unable to convert string to objectId")
	ErrSerialiseEntityToDoc      = errors.New("serialise: unable to convert entity to doc")
	ErrSerialiseDocToEntity      = errors.New("serialise: unable to convert doc to entity")
)

type BaseDoc[ID primitive.ObjectID | uint64] struct {
	Id        ID    `bson:"_id,omitempty"`
	CreatedAt int64 `bson:"createdAt"`
	UpdatedAt int64 `bson:"updatedAt"`
}

type BaseEntity[ID string | uint64] struct {
	Id        ID
	CreatedAt int64
	UpdatedAt int64
}

func ConvertStringToId(hexId string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hexId)
}

func ConvertIdToString(id primitive.ObjectID) (string, error) {
	return id.Hex(), nil
}

func ConvertEntityToDocWithNumId[Entity BaseEntity[uint64], Doc BaseDoc[uint64]](request *Entity, doc *Doc) error {
	if err := copier.Copy(&doc, request); err != nil {
		return ErrSerialiseEntityToDoc
	}
	return nil
}

func ConvertEntityToDocWithObjectId(request *BaseEntity[string], doc *BaseDoc[primitive.ObjectID]) error {
	if err := copier.Copy(&doc, request); err != nil {
		return ErrSerialiseEntityToDoc
	}
	objectId, err := ConvertStringToId(request.Id)
	if err != nil {
		return err
	}
	doc.Id = objectId
	return nil
}

func ConvertDocToEntityObjectId(doc *BaseDoc[primitive.ObjectID], request *BaseEntity[string]) error {
	if err := copier.Copy(&request, doc); err != nil {
		return ErrSerialiseDocToEntity
	}
	entityId, err := ConvertIdToString(doc.Id)
	if err != nil {
		return err
	}
	request.Id = entityId
	return nil
}

func ConvertDocToEntityNumId(doc *BaseDoc[uint64], request *BaseEntity[uint64]) error {
	if err := copier.Copy(&request, doc); err != nil {
		return ErrSerialiseDocToEntity
	}
	return nil
}
