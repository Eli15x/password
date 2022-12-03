package models

type Rule struct {
	Rule string `json:"rule,omitempty" bson:"rule"`
	Value int `json:"value,omitempty" bson:"value"`
}
