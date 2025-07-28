package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerMarketValue struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PlayerID                        int64              `bson:"player_id" json:"player_id"`
	Date                            time.Time          `bson:"date" json:"date"`
	MarketValueInEUR                int                `bson:"market_value_in_eur" json:"market_value_in_eur"`
	CurrentClubID                   int64              `bson:"current_club_id" json:"current_club_id"`
	PlayerClubDomesticCompetitionID string             `bson:"player_club_domestic_competition_id" json:"player_club_domestic_competition_id"`
}
