package model

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type ContestPlayer struct {
	Id           primitive.ObjectID          `bson:"_id"`
	Contest      primitive.ObjectID          `bson:"contest"`
	User         primitive.ObjectID          `bson:"user"`
	RegisterTime time.Time                   `bson:"register_time"`
	Problems     map[string]ContestPlayerProblemEntry `bson:"problems"`
}

type ContestPlayerProblemEntry struct {
	Submissions []primitive.ObjectID `bson:"submissions"`
}
