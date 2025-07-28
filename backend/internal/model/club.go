package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Club struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClubID                int64              `bson:"club_id" json:"club_id"`
	ClubCode              string             `bson:"club_code" json:"club_code"`
	Name                  string             `bson:"name" json:"name"`
	DomesticCompetitionID string             `bson:"domestic_competition_id" json:"domestic_competition_id"`
	TotalMarketValue      string             `bson:"total_market_value" json:"total_market_value"` // string because of +€3.05m format
	SquadSize             int                `bson:"squad_size" json:"squad_size"`
	AverageAge            float64            `bson:"average_age" json:"average_age"`
	ForeignersNumber      int                `bson:"foreigners_number" json:"foreigners_number"`
	ForeignersPercentage  float64            `bson:"foreigners_percentage" json:"foreigners_percentage"`
	NationalTeamPlayers   int                `bson:"national_team_players" json:"national_team_players"`
	StadiumName           string             `bson:"stadium_name" json:"stadium_name"`
	StadiumSeats          int                `bson:"stadium_seats" json:"stadium_seats"`
	NetTransferRecord     string             `bson:"net_transfer_record" json:"net_transfer_record"` // string format "+€3.05m"
	CoachName             string             `bson:"coach_name" json:"coach_name"`
	LastSeason            int                `bson:"last_season" json:"last_season"`
	Filename              string             `bson:"filename" json:"filename"`
	URL                   string             `bson:"url" json:"url"`
}
