package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Entry describes a MongoDB entry.
type Entry struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Short     string             `bson:"short"`
	URL       string             `bson:"url"`
}

// Model interface describes what functions are available to the File object
type Model interface {
	Init() error

	Store(string, string) error
	Delete(int) error
	Get(interface{}) (Entry, error)

	Contains(string) bool
}

type Database struct {
	Key        string
	collection *mongo.Collection
	context    context.Context
}

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

func (d *Database) Store(short string, full string) error {
	err := d.storeEntry(d.createEntry(full, short))
	return err
}
func (d *Database) Delete(id int) error {
	_, err := d.collection.DeleteOne(d.context, bson.M{"_id": id})
	return err
}

func (d *Database) Get(filter interface{}) (Entry, error) {
	res := d.collection.FindOne(d.context, filter)

	if res.Err() != nil {
		return Entry{}, res.Err()
	}

	entry := Entry{}
	err := res.Decode(&entry)

	if err != nil {
		return Entry{}, err
	}

	return entry, nil
}
func (d *Database) Contains(str string) bool {
	res := d.collection.FindOne(d.context, bson.M{"url": str})

	if res.Err() != nil {
		return false
	}
	return true
}

func (d *Database) storeEntry(entry *Entry) error {
	_, err := d.collection.InsertOne(d.context, entry)
	return err
}

func (d *Database) createEntry(short string, full string) *Entry {
	return &Entry{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		Short:     short,
		URL:       full,
	}
}
