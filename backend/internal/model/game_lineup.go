package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameLineup struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GameLineupsID string             `bson:"game_lineups_id" json:"game_lineups_id"`
	Date          time.Time          `bson:"date" json:"date"`
	GameID        int64              `bson:"game_id" json:"game_id"`
	PlayerID      int64              `bson:"player_id" json:"player_id"`
	ClubID        int64              `bson:"club_id" json:"club_id"`
	PlayerName    string             `bson:"player_name" json:"player_name"`
	Type          string             `bson:"type" json:"type"` // e.g., "starting_lineup"
	Position      string             `bson:"position" json:"position"`
	Number        int                `bson:"number" json:"number"`
	TeamCaptain   int                `bson:"team_captain" json:"team_captain"` // 0 or 1; can use bool if preferred
}
