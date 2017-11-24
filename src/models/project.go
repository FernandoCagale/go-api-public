package models

import (
	shared "github.com/FernandoCagale/go-api-shared/src/validation"
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Tags        []Tag         `json:"tags" bson:"tags"`
}

func (t Project) Validate() (errors map[string][]shared.Validation, ok bool) {
	errors = make(map[string][]shared.Validation)

	if t.Name == "" {
		errors["name"] = append(errors["name"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
	}

	if len(t.Tags) == 0 {
		errors["tags"] = append(errors["tags"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
	}

	return errors, len(errors) == 0
}
