package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Establish connection with MongoDB.
	client, err := mongo.NewClient(options.
		Client().ApplyURI("mongodb://localhost:27017").
		SetAuth(options.Credential{
			Username: "satan",
			Password: "666",
		}))
	handleError(err)
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()
	err = client.Connect(ctx)
	handleError(err)

	// Upload the file to MongoDB.
	file, err := os.Open("last_occultism_by_perldrop.mid")
	handleError(err)
	bucket, err := gridfs.NewBucket(client.Database("test"))
	handleError(err)
	fileID, err := bucket.UploadFromStream("last_occultism.mid", file)
	handleError(err)

	fmt.Println("File uploaded. It's ID:", fileID.String())

	// Download the file from MongoDB.
	newFile, err := os.OpenFile("last_occultism.mid", os.O_CREATE|os.O_WRONLY, 0777)
	handleError(err)
	_, err = bucket.DownloadToStream(fileID, newFile)
	handleError(err)

	// Delete the uploaded file.
	err = bucket.Delete(fileID)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
