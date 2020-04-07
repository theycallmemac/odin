package odinlib

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"

)

func Log(varType string, desc string, value string, id string) bool {
    return FindAndInsert(varType, desc, value, id)
}

func FindAndInsert(varType string, desc string, value string, id string) bool {
    client := SetupClient()
    filter := bson.M{"id": id, "desc": desc, "type": varType}
    update := bson.M{"$set": bson.M{"type": varType, "desc": desc, "value": value, "id": id,}}
    collection := client.Database("odin").Collection("observability")
    _, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
    return err != nil
}

func SetupClient() *mongo.Client {
    c := getMongoClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your MongoDB instance is running")
        os.Exit(2)
    }
    return c
}

func getMongoClient() *mongo.Client {
    url, _ := os.LookupEnv("ODIN_MONGODB")
    clientOptions := options.Client().ApplyURI(url)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your `ODIN_MONGODB` environment variable")
    }
    err = client.Connect(context.Background())
    if err != nil {
        fmt.Println("Cannot connect to MongoDB - check your `ODIN_MONGODB` environment variable")
    }
    return client
}

