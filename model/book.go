package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"Title,omitempty"`
	Description string             `json:"Description,omitempty"`
	ISBN        string             `json:"ISBN,omitempty"`
}

func (b *Book) PreparePut() {

	if b.ID.IsZero() || b.ID.Hex() == "" {
		b.ID = primitive.NewObjectID()
	}
}
