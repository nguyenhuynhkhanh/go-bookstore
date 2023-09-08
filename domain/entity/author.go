package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Author ...
type Author struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	LastName    string             `json:"lastName" bson:"lastName"`
	BirthDate   time.Time          `json:"birthDate" bson:"birthDate"`
	Nationality string             `json:"nationality" bson:"nationality"`
	CreatedAt   string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string             `json:"updatedAt" bson:"updatedAt"`
}
