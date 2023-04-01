package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
)

func loggerMessage(msg string) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}

	//defer logger.Sync()

	logger.Info(msg)

}

func InitConnection() *mongo.Client {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	loggerMessage("MongoDB conected!")

	return client
}
