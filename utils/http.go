package utils

import (
	"encoding/json"
	"net/http"
)

func HttpOkJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusOK, data, true)
	return
}

func HttpBadRequestJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusBadRequest, data, true)
	return
}

func HttpOkRaw(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusOK, data, false)
	return
}

func HttpBadRequestRaw(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusBadRequest, data, false)
	return
}

func writeResponseJSON(writer http.ResponseWriter, httpStatus int, data interface{}, isJson bool) (err error) {
	if isJson {
		writer.Header().Set("Content-Type", "application/json")
	}

	var b []byte

	b, err = json.Marshal(map[string]interface{}{
		"msg": data,
	})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}

	writer.WriteHeader(httpStatus)
	writer.Write(b)
	return
}
