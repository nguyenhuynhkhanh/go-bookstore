package repository

import (
	"context"
	"log"
	"time"

	entities "bookstore.com/domain/entity"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const AuthorCollectionName = "authors"

//Repository ...
type AuthorRepository interface {
	Find(id primitive.ObjectID) (*entities.Author, error)
	Store(author *entities.Author) error
	FindAll() ([]*entities.Author, error)
	Delete(id primitive.ObjectID) error
}

//NewMongoRepository ...
func NewAuthorMongoRepository(mongoServerURL, mongoDb string, timeout int) (AuthorRepository, error) {
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

func (r *mongoRepository) Store(author *entities.Author) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection(AuthorCollectionName)

	now := time.Now()
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"_id":         primitive.NewObjectID(),
			"firstName":   author.FirstName,
			"lastName":    author.LastName,
			"birthDate":   author.BirthDate,
			"nationality": author.Nationality,
			"createdAt":   now,
			"updatedAt":   now,
		},
	)
	if err != nil {
		return errors.Wrap(err, "Add author error!")
	}
	return nil
}

func (r *mongoRepository) Find(id primitive.ObjectID) (*entities.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	author := &entities.Author{}
	collection := r.client.Database(r.db).Collection("authors")

	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(author)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, errors.Wrap(err, "repository research")
	}
	return author, nil

}

func (r *mongoRepository) FindAll() ([]*entities.Author, error) {
	var authors []*entities.Author
	collection := r.client.Database(r.db).Collection(AuthorCollectionName)
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item entities.Author
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		authors = append(authors, &item)
	}
	return authors, nil
}

func (r *mongoRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"_id": id}

	collection := r.client.Database(r.db).Collection(AuthorCollectionName)
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
