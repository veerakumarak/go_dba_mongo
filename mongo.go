package mongo

import (
	"context"
	dba "github.com/veerakumarak/go_dba_core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository[Entity any, Id string | uint64] struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func (r *mongoRepository[Entity, Id]) Count() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	return uint64(count), nil
}

func (r *mongoRepository[Entity, Id]) FindById(entity *Entity, id Id) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	// ToDo check if this is required
	//objectId, err := convertStringToId(id)
	//if err != nil {
	//	return err
	//}

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(entity)
	if err == mongo.ErrNoDocuments {
		return dba.ErrRepositoryEntityNotFound
	}
	if err != nil {
		return dba.ErrRepositoryInternalError
	}
	return nil
}

//func (r *mongoRepository[Entity]) Find(search dba.Search) ([]Entity, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (r *mongoRepository[Entity, Id]) Save(entity *Entity) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return dba.ErrRepositoryInternalError
	}

	return nil
}

func (r *mongoRepository[Entity, Id]) SaveAll(entity []Entity) error {
	//TODO implement me
	panic("implement me")
}

func (r *mongoRepository[Entity, Id]) ExistsById(id Id) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	filter := bson.M{"_id": id}
	count, err := r.collection.CountDocuments(ctx, filter, options.Count().SetLimit(1))
	if err != nil {
		return false, dba.ErrRepositoryInternalError
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *mongoRepository[Entity, Id]) DeleteById(id Id) error {
	//TODO implement me
	panic("implement me")
}

//func (r *mongoRepository[Entity]) Delete(search dba.Search) ([]Entity, error) {
//	//TODO implement me
//	panic("implement me")
//}
