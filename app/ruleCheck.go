package app

import (
	"go.mongodb.org/mongo-driver/bson"
)

func Action(ruleName string) (bson.M, int) {
	if ruleName == "default" {
		code := 204
		return bson.M{"SMTP_ACTION": ruleName}, code
	}
	code := 200
	return bson.M{"SMTP_ACTION": ruleName}, code
}
