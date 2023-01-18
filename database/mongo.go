package database

import (
	"context"
	"log"
	"time"

	"tea/domain"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func generateUUID() string {
	return uuid.New().String()
}

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

//Create a Mongo Client
func newMongClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerURL))
	if err != nil {
		return nil, err
	}
	//We could ping the server to test connectivity if we want

	return client, nil
}

//NewMongoRepository ...
func NewMongoRepository(mongoServerURL, mongoDb string, timeout int) (domain.Repository, error) {
	mongoClient, err := newMongClient(mongoServerURL, timeout)
	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil

}

func (r *mongoRepository) Store(Store *domain.Store) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"id":          generateUUID(),
			"name":        Store.Name,
			"description": Store.Description,
			"address":     Store.Address,
			"phone":       Store.Phone,
			"latitude":    Store.Latitude,
			"longitude":   Store.Longitude,
			"menuItems":   Store.MenuItems,
		},
	)
	if err != nil {
		return errors.Wrap(err, "Error writing to repository")
	}
	return nil

}
func (r *mongoRepository) Update(Store *domain.Store) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"code": Store.ID},
		bson.M{
			"$set": bson.M{
				"name":        Store.Name,
				"description": Store.Description,
				"address":     Store.Address,
				"phone":       Store.Phone,
				"latitude":    Store.Latitude,
				"longitude":   Store.Longitude,
				"menuItems":   Store.MenuItems,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) Find(code string) (*domain.Store, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	Store := &domain.Store{}
	collection := r.client.Database(r.db).Collection("items")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(Store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, errors.Wrap(err, "repository research")
	}
	return Store, nil

}

func (r *mongoRepository) FindAll() ([]*domain.Store, error) {

	var items []*domain.Store

	collection := r.client.Database(r.db).Collection("items")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item domain.Store
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil

}
func (r *mongoRepository) Delete(code string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"code": code}

	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
