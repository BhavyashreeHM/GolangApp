package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://Bhavya:admin123@localhost:27017"

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	client = c
	fmt.Println("✅ Connected to MongoDB")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server running + MongoDB connected"))

	collection := client.Database("testdb").Collection("users")

	doc := map[string]interface{}{
		"name": "Bhavya",
		"age":  25,
	}

	collection.InsertOne(context.TODO(), doc)

}

func main() {
	connectDB()

	http.HandleFunc("/", handler)

	fmt.Println("🚀 Server running on :30031")
	log.Fatal(http.ListenAndServe(":30031", nil))
}
