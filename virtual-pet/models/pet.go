package models

type Pet struct {
	ID        string  `json:"id" bson:"_id"`
	Name      string  `json:"name" bson:"name"`
	Happiness float64 `json:"happiness" bson:"happiness"`
	Hunger    float64 `json:"hunger" bson:"hunger"`
	Energy    float64 `json:"energy" bson:"energy"`
}
