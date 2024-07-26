package database

import (
	"context"
	"log"
	"os"

	"github.com/Peter-Immanuel/fox-alpine/pkg/domain"
	"github.com/ajclopez/mgs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Primitives struct{}

func (primitives *Primitives) ObjectID(oidStr string) (interface{}, error) {
	return primitive.ObjectIDFromHex(oidStr)
}

type mongoStore struct {
	Database *mongo.Database
}

func NewMongoStore() (domain.PetDB, error) {

	url := os.Getenv("DATABASE_URL")
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(url),
	)
	if err != nil {
		return nil, err
	}
	return mongoStore{Database: client.Database(os.Getenv("DATABASE"))}, nil
}

func (ms mongoStore) convertID(id domain.PetID) string {
	petID, correct := id.(string)
	if !correct {
		log.Fatal("Invalid ID")
	}
	return petID
}

func (ms mongoStore) parseQuery(query string) *options.FindOptions {

	queryHandler := mgs.NewQueryHandler(&Primitives{})
	opts := mgs.FindOption()
	result, _ := queryHandler.MongoGoSearch(query, opts)
	findOpts := options.Find()
	findOpts.SetLimit(result.Limit)
	findOpts.SetSkip(result.Skip)
	findOpts.SetSort(result.Sort)
	findOpts.SetProjection(result.Projection)

	return findOpts
}

func (ms mongoStore) Get(id domain.PetID) (*domain.Pet, error) {
	search_id, err := primitive.ObjectIDFromHex(ms.convertID(id))
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": search_id}
	pet := domain.Pet{}
	err = ms.Database.Collection("pets").FindOne(
		context.TODO(),
		filter,
	).Decode(&pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil

}

func (ms mongoStore) List(category string) ([]*domain.Pet, error) {
	var result []*domain.Pet
	cur, err := ms.Database.Collection("pets").Find(
		context.TODO(),
		bson.M{"category": category},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = cur.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result, nil
}

func (ms mongoStore) Create(pet *domain.Pet) (*domain.Pet, error) {
	pet.ID = primitive.NewObjectID()
	_, err := ms.Database.Collection("pets").InsertOne(context.TODO(), pet)
	return pet, err
}

func (ms mongoStore) Delete(id domain.PetID) error {
	objID := ms.convertID(id)
	petId, err := primitive.ObjectIDFromHex(objID)
	filter := bson.M{"_id": petId}
	_, err = ms.Database.Collection("pets").DeleteOne(
		context.TODO(),
		filter,
	)
	return err
}
