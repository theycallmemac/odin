package odinlib

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"

)

func Log(varType string, desc string, value string, id string, timestamp string) bool {
    return FindAndInsert(varType, desc, value, id, timestamp)
}

func FindAndInsert(varType string, desc string, value string, id string, timestamp string) bool {
    client, err := SetupClient()
    if err != nil {
	return false
    }
    filter := bson.M{"id": id, "desc": desc, "type": varType}
    update := bson.M{"$set": bson.M{"type": varType, "desc": desc, "value": value, "id": id, "timestamp": timestamp}}
    collection := client.Database("odin").Collection("observability")
    _, err = collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
    return err == nil
}

func SetupClient() (*mongo.Client, error) {
    url, _ := os.LookupEnv("ODIN_MONGODB")
    c := getMongoClient(url)
    ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
    err := c.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your MongoDB instance is running")
        c.Disconnect(context.TODO())
    }
    return c, err
}

func getMongoClient(url string) *mongo.Client {
    clientOptions := options.Client().ApplyURI(url)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your `ODIN_MONGODB` environment variable")
        return nil
    }
    err = client.Connect(context.Background())
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your `ODIN_MONGODB` environment variable")
        return nil
    }
    return client
}

