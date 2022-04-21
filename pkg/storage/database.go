package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Database implements LinkStorage interface using a (mongo) database for storing data.
type Database struct {
	Key        string
	collection *mongo.Collection
	context    context.Context
}

// DatabaseEntry describes a MongoDB entry.
type DatabaseEntry struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Short     string             `bson:"short"`
	URL       string             `bson:"url"`
}

// createEntry used to create a database entry.
func createEntry(hash string, URL string) *DatabaseEntry {
	return &DatabaseEntry{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		Short:     hash,
		URL:       URL,
	}
}

// Init connects to the database.
func (d *Database) Init() error {
	d.context = context.TODO()
	uri := fmt.Sprintf("mongodb+srv://suicedek:%s@mongocluster.qkbsn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", d.Key)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(d.context, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(d.context, nil)
	if err != nil {
		return err
	}

	d.collection = client.Database("shorturlDB").Collection("urls")
	return nil
}

// Store stores the data in storage.
func (d *Database) Store(hash string, URL string) error {
	entry := createEntry(hash, URL)
	_, err := d.collection.InsertOne(d.context, entry)
	return err
}

// Delete removes data from the storage.
func (d *Database) Delete(id int) error {
	_, err := d.collection.DeleteOne(d.context, bson.M{"_id": id})
	return err
}

// Get searches the storage for the entry using given filter.
func (d *Database) Get(filter interface{}) (DatabaseEntry, error) {
	res := d.collection.FindOne(d.context, filter)

	if res.Err() != nil {
		return DatabaseEntry{}, res.Err()
	}

	entry := DatabaseEntry{}
	err := res.Decode(&entry)

	if err != nil {
		return DatabaseEntry{}, err
	}

	return entry, nil
}

// Contains checks if the storage contains the URL.
func (d *Database) Contains(str string) bool {
	res := d.collection.FindOne(d.context, bson.M{"url": str})

	if res.Err() != nil {
		return false
	}
	return true
}
