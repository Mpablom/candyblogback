package work

import "go.mongodb.org/mongo-driver/bson/primitive"

type Work struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Image       string             `bson:"image"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Gallery     []Gallery          `bson:"gallery"`
}

type Gallery struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Image string             `bson:"image"`
}
