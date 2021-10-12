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
	Action     string `json:"action"`
	Type       string `json:"type"`
	CadStringA string `json:"cad_string_a,omitempty"`
	CadStringB string `json:"cad_string_b,omitempty"`
}

func APIHandler(writer http.ResponseWriter, request *http.Request) {
	logger := log.Begin()
	defer logger.End()
	logger.Trace("new request incoming", request)

	var cr *CalculateRequest
	err := httpjson.ParseRequest(writer, request, &cr)
	if err != nil {
		logger.Error("Error parsing request")
		return
	}

	if cr.Action == "calculate" {
		/*
			{
				"action": "calculate",
				"type": "ADD",
				"cad_string_a": "$1.29",
				"cad_string_b": "$2.30"
			}
		*/
		if cr.CadStringA == "" || cr.CadStringB == "" {
			httpjson.BadRequest(writer, "empty_cad_strings")
			logger.Error("empty_cad_strings")
			return
		}

		cadA, err := cad.ParseCAD(cr.CadStringA)
		if err != nil {
			httpjson.BadRequest(writer, "invalid_cad_a")
			logger.Error("invalid_cad", err)
			return
		}

		cadB, err := cad.ParseCAD(cr.CadStringB)
		if err != nil {
			httpjson.BadRequest(writer, "invalid_cad_b")
			logger.Error("invalid_cad", err)
			return
		}

		if cr.Type == "ADD" {
			httpjson.Ok(writer, cadA.Add(cadB).AsCents())
			return
		}
	}
	httpjson.BadRequest(writer, "invalid_action_or_type")
}
