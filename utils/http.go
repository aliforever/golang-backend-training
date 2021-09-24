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
	writer.Header().Set("Content-Type", "application/json")

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
