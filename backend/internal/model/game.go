package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GameID              int64              `bson:"game_id" json:"game_id"`
	CompetitionID       string             `bson:"competition_id" json:"competition_id"`
	Season              int                `bson:"season" json:"season"`
	Round               string             `bson:"round" json:"round"`
	Date                time.Time          `bson:"date" json:"date"`
	HomeClubID          int64              `bson:"home_club_id" json:"home_club_id"`
	AwayClubID          int64              `bson:"away_club_id" json:"away_club_id"`
	HomeClubGoals       int                `bson:"home_club_goals" json:"home_club_goals"`
	AwayClubGoals       int                `bson:"away_club_goals" json:"away_club_goals"`
	HomeClubPosition    int                `bson:"home_club_position" json:"home_club_position"`
	AwayClubPosition    int                `bson:"away_club_position" json:"away_club_position"`
	HomeClubManagerName string             `bson:"home_club_manager_name" json:"home_club_manager_name"`
	AwayClubManagerName string             `bson:"away_club_manager_name" json:"away_club_manager_name"`
	Stadium             string             `bson:"stadium" json:"stadium"`
	Attendance          int                `bson:"attendance" json:"attendance"`
	Referee             string             `bson:"referee" json:"referee"`
	URL                 string             `bson:"url" json:"url"`
	HomeClubFormation   string             `bson:"home_club_formation" json:"home_club_formation"`
	AwayClubFormation   string             `bson:"away_club_formation" json:"away_club_formation"`
	HomeClubName        string             `bson:"home_club_name" json:"home_club_name"`
	AwayClubName        string             `bson:"away_club_name" json:"away_club_name"`
	Aggregate           string             `bson:"aggregate" json:"aggregate"`
	CompetitionType     string             `bson:"competition_type" json:"competition_type"`
}
