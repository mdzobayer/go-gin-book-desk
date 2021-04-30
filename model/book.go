package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"Title"`
	Description string             `json:"Description"`
	ISBN        string             `json:"ISBN"`
}
