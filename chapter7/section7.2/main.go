package main

import (
	"encoding/json"
	"strings"

	"github.com/aliforever/golang-backend-training/chapter7/section7.2/srv/logger"
)

type Info struct {
	GivenName     string `json:"given_name"`
	AdditionNames string `json:"additional_names"`
	FamilyName    string `json:"family_name"`
}

func main() {
	logger.DefaultLogger = logger.DefaultLogger.Begin()

	var jsonString strings.Builder

	info := Info{
		GivenName:  "Joe",
		FamilyName: "Blow",
	}

	err := json.NewEncoder(&jsonString).Encode(info)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return
	}

	logger.DefaultLogger.LogF("json string: %s", jsonString.String())
}
