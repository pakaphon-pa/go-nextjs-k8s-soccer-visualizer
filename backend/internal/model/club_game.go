package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClubGame struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GameID              int64              `bson:"game_id" json:"game_id"`
	ClubID              int64              `bson:"club_id" json:"club_id"`
	OwnGoals            int                `bson:"own_goals" json:"own_goals"`
	OwnPosition         string             `bson:"own_position" json:"own_position"`
	OwnManagerName      string             `bson:"own_manager_name" json:"own_manager_name"`
	OpponentID          int64              `bson:"opponent_id" json:"opponent_id"`
	OpponentGoals       int                `bson:"opponent_goals" json:"opponent_goals"`
	OpponentPosition    string             `bson:"opponent_position" json:"opponent_position"`
	OpponentManagerName string             `bson:"opponent_manager_name" json:"opponent_manager_name"`
	Hosting             string             `bson:"hosting" json:"hosting"`
	IsWin               int                `bson:"is_win" json:"is_win"` // 0 or 1, consider bool if preferred
}
