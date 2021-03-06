package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailConfirmation struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user" validate:"required" bson:"user"`
	Code      int                `json:"code" validate:"required" bson:"code"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
