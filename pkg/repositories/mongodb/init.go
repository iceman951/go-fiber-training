package mongodb

import (
	"context"
	"fmt"
	"calcal/pkg/utils/constant"
	"time"

	"github.com/baac-tech/zlogwrap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(dbData DB) (client *mongo.Client, db string) {

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: constant.INIT_MONGO,
	})

	dsn := fmt.Sprintf("mongodb://%v:%v",
		dbData.IP,
		dbData.Port,
	)
	auth := options.Credential{
		AuthMechanism: dbData.AuthMechanism,
		Username:      dbData.User,
		Password:      dbData.Password,
		AuthSource:    dbData.DBName,
	}

	db = dbData.DBName

	clientOptions := options.Client()
	clientOptions.ApplyURI(dsn)
	clientOptions.SetDirect(true)
	if dbData.User != "" && dbData.Password != "" {
		clientOptions.SetAuth(auth)
	}

	err := clientOptions.Validate()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	return client, db
}
