package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appearance struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ObjectID
	AppearanceID        int64              `bson:"appearance_id" json:"appearance_id"`
	GameID              int64              `bson:"game_id" json:"game_id"`
	PlayerID            int64              `bson:"player_id" json:"player_id"`
	PlayerClubID        int64              `bson:"player_club_id" json:"player_club_id"`
	PlayerCurrentClubID int64              `bson:"player_current_club_id" json:"player_current_club_id"`
	Date                time.Time          `bson:"date" json:"date"` // Use ISO8601 format when marshaling JSON
	PlayerName          string             `bson:"player_name" json:"player_name"`
	CompetitionID       string             `bson:"competition_id" json:"competition_id"`
	YellowCards         int                `bson:"yellow_cards" json:"yellow_cards"`
	RedCards            int                `bson:"red_cards" json:"red_cards"`
	Goals               int                `bson:"goals" json:"goals"`
	Assists             int                `bson:"assists" json:"assists"`
	MinutesPlayed       int                `bson:"minutes_played" json:"minutes_played"`
}
