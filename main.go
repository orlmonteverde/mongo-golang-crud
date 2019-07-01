package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Name of the database.
	DBName          = "glottery"
	notesCollection = "notes"
	URI             = "mongodb://<user>:<password>@<host>/<name>"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

func main() {
	// Base context.
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	db := client.Database(DBName)
	coll := db.Collection(notesCollection)

	// Insert One Document.
	note := Note{}

	// An ID for MongoDB.
	note.ID = primitive.NewObjectID()
	note.Title = "First note"
	note.Body = "Some spam text"
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	result, err := coll.InsertOne(ctx, note)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	// Insert Many Documents.
	notes := []interface{}{}

	note2 := Note{}
	note2.ID = primitive.NewObjectID()
	note2.Title = "First note"
	note2.Body = "Some spam text"
	note2.CreatedAt = time.Now()
	note2.UpdatedAt = time.Now()

	note3 := Note{
		ID:        primitive.NewObjectID(),
		Title:     "Third note",
		Body:      "Some spam text",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	notes = append(notes, note2, note3)

	results, err := coll.InsertMany(ctx, notes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results.InsertedIDs)

	// Update a document.

	// Parsing a string ID to ObjectID from MongoDB.
	objID, err := primitive.ObjectIDFromHex("5d100d9c23affb7006dd9cff")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultUpdate, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"body":       "Some updated text",
				"updated_at": time.Now(),
			},
		},
	)

	fmt.Println(resultUpdate.ModifiedCount)

	// Delete one document.
	objID, err = primitive.ObjectIDFromHex("5d1017bda2dc2ce292e5f16c")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultDelete, err := coll.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultDelete.DeletedCount)

	// Find documents.

	objID, err = primitive.ObjectIDFromHex("5d100d9c23affb7006dd9cff")
	if err != nil {
		fmt.Println(err)
		return
	}

	findResult := coll.FindOne(ctx, bson.M{"_id": objID})
	if err := findResult.Err(); err != nil {
		fmt.Println(err)
		return
	}

	n := Note{}
	err = findResult.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n.Body)

	notesResult := []Note{}
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(ctx) {
		cursor.Decode(&n)
		notesResult = append(notesResult, n)
	}

	for _, el := range notesResult {
		fmt.Println(el.Title)
	}
}
