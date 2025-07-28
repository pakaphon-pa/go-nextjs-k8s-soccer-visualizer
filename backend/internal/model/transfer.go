package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerTransfer struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PlayerID         int64              `bson:"player_id" json:"player_id"`
	TransferDate     time.Time          `bson:"transfer_date" json:"transfer_date"`
	TransferSeason   string             `bson:"transfer_season" json:"transfer_season"` // e.g., "26/27"
	FromClubID       int64              `bson:"from_club_id" json:"from_club_id"`
	ToClubID         int64              `bson:"to_club_id" json:"to_club_id"`
	FromClubName     string             `bson:"from_club_name" json:"from_club_name"`
	ToClubName       string             `bson:"to_club_name" json:"to_club_name"`
	TransferFee      int64              `bson:"transfer_fee" json:"transfer_fee"`               // in EUR
	MarketValueInEUR int64              `bson:"market_value_in_eur" json:"market_value_in_eur"` // at time of transfer
	PlayerName       string             `bson:"player_name" json:"player_name"`
}
