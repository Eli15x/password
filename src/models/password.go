package models

type PasswordSetting struct {
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	Rules 	  []Rule `json:"rules,omitempty" bson:"rules"`
}
