package main

import (
	"net/http"

	"github.com/aliforever/go-cad"

	"github.com/aliforever/go-httpjson"
	"github.com/aliforever/golang-backend-training/chapter10/section10.3/srv/log"
)

func main() {
	logger := log.Begin()
	defer logger.End()
	mux := http.NewServeMux()
	mux.HandleFunc("/api", APIHandler)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		logger.Error("Error listening on port 80", err)
		return
	}
}

type CalculateRequest struct {
	Type       string `json:"type"`
	CadStringA string `json:"cad_string_a,omitempty"`
	CadStringB string `json:"cad_string_b,omitempty"`
}

func APIHandler(writer http.ResponseWriter, request *http.Request) {
	logger := log.Begin()
	defer logger.End()
	logger.Trace("new request incoming", request)

	r, err := httpjson.ParseRequest(writer, request)
	if err != nil {
		logger.Error("Error parsing request")
		return
	}

	if r.Action() == "calculate" {
		/*
			{
				"action": "calculate",
				"data": {
					"type": "ADD",
					"cad_string_a": "$1.29",
					"cad_string_b": "$2.30"
				}
			}
		*/
		var data *CalculateRequest
		err = r.UnmarshalData(&data)
		if err != nil {
			httpjson.BadRequest(writer, "invalid_json")
			logger.Error("invalid_json_request_data", err)
			return
		}
		if data.CadStringA == "" || data.CadStringB == "" {
			httpjson.BadRequest(writer, "empty_cad_strings")
			logger.Error("empty_cad_strings")
			return
		}

		cadA, err := cad.ParseCAD(data.CadStringA)
		if err != nil {
			httpjson.BadRequest(writer, "invalid_cad_a")
			logger.Error("invalid_cad", err)
			return
		}

		cadB, err := cad.ParseCAD(data.CadStringB)
		if err != nil {
			httpjson.BadRequest(writer, "invalid_cad_b")
			logger.Error("invalid_cad", err)
			return
		}

		if data.Type == "ADD" {
			httpjson.Ok(writer, cadA.Add(cadB).AsCents())
			return
		}
	}
}
