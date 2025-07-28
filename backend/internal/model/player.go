package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ObjectID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PlayerID                  int                `bson:"player_id" json:"player_id"`
	FirstName                 string             `bson:"first_name" json:"first_name"`
	LastName                  string             `bson:"last_name" json:"last_name"`
	Name                      string             `bson:"name" json:"name"`
	LastSeason                int                `bson:"last_season" json:"last_season"`
	CurrentClubID             int                `bson:"current_club_id" json:"current_club_id"`
	PlayerCode                string             `bson:"player_code" json:"player_code"`
	CountryOfBirth            string             `bson:"country_of_birth" json:"country_of_birth"`
	CityOfBirth               string             `bson:"city_of_birth" json:"city_of_birth"`
	CountryOfCitizenship      string             `bson:"country_of_citizenship" json:"country_of_citizenship"`
	DateOfBirth               time.Time          `bson:"date_of_birth" json:"date_of_birth"`
	SubPosition               string             `bson:"sub_position" json:"sub_position"`
	Position                  string             `bson:"position" json:"position"`
	Foot                      string             `bson:"foot" json:"foot"`
	HeightInCM                int                `bson:"height_in_cm" json:"height_in_cm"`
	ContractExpirationDate    string             `bson:"contract_expiration_date" json:"contract_expiration_date"`
	AgentName                 string             `bson:"agent_name" json:"agent_name"`
	ImageURL                  string             `bson:"image_url" json:"image_url"`
	URL                       string             `bson:"url" json:"url"`
	CurrentClubDomesticCompID string             `bson:"current_club_domestic_competition_id" json:"current_club_domestic_competition_id"`
	CurrentClubName           string             `bson:"current_club_name" json:"current_club_name"`
	MarketValueInEUR          int                `bson:"market_value_in_eur" json:"market_value_in_eur"`
	HighestMarketValueInEUR   int                `bson:"highest_market_value_in_eur" json:"highest_market_value_in_eur"`
}
