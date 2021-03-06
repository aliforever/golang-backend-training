package main

import (
	"encoding/json"

	"github.com/aliforever/golang-backend-training/chapter7/section7.1/srv/logger"

	_ "github.com/reiver/go-simplehttp/driver/json"
)

type Fruit struct {
	Apple  bool   `json:"apple"`
	Banana string `json:"banana"`
	Cherry int64  `json:"cherry"`
}

func main() {
	log := logger.Begin()
	defer log.End()

	var data map[string]interface{} = map[string]interface{}{
		"apple":  "ONE",
		"banana": "TWO",
		"cherry": "THREE",
	}

	p, err := json.Marshal(data)
	if err != nil {
		log.Error("Error marshaling JSON", err)
		return
	}

	log.Log("First Json:", string(p))

	newFruit := Fruit{
		Apple:  true,
		Banana: "TWO",
		Cherry: 3,
	}

	p, err = json.Marshal(newFruit)
	if err != nil {
		log.Error("Error marshaling JSON", err)
		return
	}

	log.Log("Second Json:", string(p))
}
