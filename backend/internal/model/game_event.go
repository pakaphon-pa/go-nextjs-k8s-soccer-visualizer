package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameEvent struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GameEventID    string             `bson:"game_event_id" json:"game_event_id"`
	Date           time.Time          `bson:"date" json:"date"`
	GameID         int64              `bson:"game_id" json:"game_id"`
	Minute         int                `bson:"minute" json:"minute"`
	Type           string             `bson:"type" json:"type"` // e.g., "Cards"
	ClubID         int64              `bson:"club_id" json:"club_id"`
	PlayerID       int64              `bson:"player_id" json:"player_id"`
	Description    string             `bson:"description" json:"description"`
	PlayerInID     string             `bson:"player_in_id,omitempty" json:"player_in_id,omitempty"`
	PlayerAssistID string             `bson:"player_assist_id,omitempty" json:"player_assist_id,omitempty"`
}
