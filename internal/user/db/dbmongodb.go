package db

import (
	"awesomeProject/internal/user"
	"awesomeProject/package/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

//методы
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil{
		return "", fmt.Errorf("failed to create user %v", err)
	}

	d.logger.Debug("convert InsertedID to ObjectID") //https://stackoverflow.com/questions/59537212/how-to-convert-mongo-insertoneresults-insertedid-to-byte
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok { //если получилось
		return oid.Hex(), nil
	}
	//nCtx, cancel :+ context.WithTimeout(ctx, 1*time.Second)
	d.logger.Trace(user) //если что-то отвалится в трейсе будет пользователь
	return "", fmt.Errorf("failed to convert object id to hex  %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return u, fmt.Errorf("failed to create convertx hex to oid %v", id)
	}
	//mongo.getDataBase("test").getCollection("docs").find({})
	filter := bson.M{" id": oid}

	//options.FindOneOptions{}()

	result := d.collection.FindOne(ctx, filter)
	if result != nil {
		return u, fmt.Errorf("failed to find user by id %s due to error %v", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user by id %s due to error %v", id, err)
	}

	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

//конструктор
func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger : logger
	}

}

//
