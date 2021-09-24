package utils

import (
	"encoding/json"
	"net/http"
)

func HttpOkJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusOK, data)
	return
}

func HttpBadRequestJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponseJSON(writer, http.StatusBadRequest, data)
	return
}

func writeResponseJSON(writer http.ResponseWriter, httpStatus int, data interface{}) (err error) {
	var b []byte
	b, err = json.Marshal(map[string]interface{}{
		"msg": data,
	})
	writer.WriteHeader(httpStatus)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(b)
	return
}
