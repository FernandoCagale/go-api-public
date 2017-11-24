package models

type Tag struct {
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}
