package mongodb

import (
	"context"
	"calcal/pkg/utils/constant"
	"time"

	"github.com/baac-tech/zlogwrap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository interface {
	Save(e interface{}) (*mongo.InsertOneResult, error)
	Update(data interface{}, valueFilter interface{}) (int, error)
	QueryOne(query map[string]interface{}, result interface{}, findOptions *options.FindOneOptions) error
	QueryAll(query map[string]interface{}, result interface{}, findOptions *options.FindOptions) error
	DeleteOne(query map[string]interface{}) (*mongo.DeleteResult, error)
}

type repositoryMongo struct {
	client *mongo.Client
	db     string
	coll   string
}

func NewRepositoryMongo(mongoClient *mongo.Client, db, coll string) MongoRepository {
	return &repositoryMongo{client: mongoClient, db: db, coll: coll}
}

func (r *repositoryMongo) Save(e interface{}) (*mongo.InsertOneResult, error) {

	collection := r.client.Database(r.db).Collection(r.coll)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return collection.InsertOne(ctx, e)
}

func (r *repositoryMongo) QueryOne(query map[string]interface{}, result interface{}, findOptions *options.FindOneOptions) error {

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: constant.QUERY_ONE,
	})

	bsonQuery := bson.M{}
	for k, v := range query {
		bsonQuery[k] = v
	}

	collection := r.client.Database(r.db).Collection(r.coll)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bsonQuery, findOptions).Decode(result)
	if err != nil {
		logger.Error("FindOne", err.Error())
	}

	return err
}

func (r *repositoryMongo) Update(data interface{}, valueFilter interface{}) (int, error) {

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: constant.UPDATE,
	})

	update, _ := ToBSonDoc(&data, bson.Marshal, bson.Unmarshal)
	opts := options.Update().SetUpsert(true)
	collection := r.client.Database(r.db).Collection(r.coll)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx, valueFilter, bson.D{{"$set", update}}, opts)
	if err != nil {
		// log.Println("UpdateOne", err.Error())
		logger.Error("UpdateOne", err.Error())
		return 0, err
	}

	return int(res.ModifiedCount), nil
}

func (r *repositoryMongo) DeleteOne(query map[string]interface{}) (*mongo.DeleteResult, error) {

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: constant.DELETE_ONE,
	})

	bsonQuery := bson.M{}
	for k, v := range query {
		bsonQuery[k] = v
	}

	collection := r.client.Database(r.db).Collection(r.coll)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, bsonQuery)
	if err != nil {
		logger.Error("DeleteOne", err.Error())
	}

	return result, nil
}

func (r *repositoryMongo) QueryAll(query map[string]interface{}, result interface{}, findOptions *options.FindOptions) error {

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: constant.QUERY_ALL,
	})

	bsonQuery := bson.M{}
	for k, v := range query {
		bsonQuery[k] = v
	}

	collection := r.client.Database(r.db).Collection(r.coll)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bsonQuery, findOptions)
	if err != nil {
		logger.Error("cursor: ", cursor)
		// log.Println("cursor: ", cursor)
		return err
	}

	err = cursor.All(ctx, result)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = cursor.Close(ctx)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil

}

func ToBSonDoc(v interface{}, marshaller func(interface{}) ([]byte, error), unmarshaller func([]byte, interface{}) error) (bson.D, error) {
	var doc bson.D
	data, err := marshaller(v)
	if err != nil {
		return doc, err
	}
	errUnmarshal := unmarshaller(data, &doc)
	return doc, errUnmarshal
}
