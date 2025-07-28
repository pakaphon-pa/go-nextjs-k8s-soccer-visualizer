package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competition struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CompetitionID         string             `bson:"competition_id" json:"competition_id"`
	CompetitionCode       string             `bson:"competition_code" json:"competition_code"`
	Name                  string             `bson:"name" json:"name"`
	SubType               string             `bson:"sub_type" json:"sub_type"`
	Type                  string             `bson:"type" json:"type"`
	CountryID             int                `bson:"country_id" json:"country_id"`
	CountryName           string             `bson:"country_name" json:"country_name"`
	DomesticLeagueCode    string             `bson:"domestic_league_code" json:"domestic_league_code"`
	Confederation         string             `bson:"confederation" json:"confederation"`
	URL                   string             `bson:"url" json:"url"`
	IsMajorNationalLeague bool               `bson:"is_major_national_league,string" json:"is_major_national_league,string"`
}
