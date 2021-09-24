package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HttpOkJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponse(writer, http.StatusOK, data, true)
	return
}

func HttpBadRequestJSON(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeResponse(writer, http.StatusBadRequest, data, true)
	return
}

func HttpOkRaw(writer http.ResponseWriter, data string) (err error) {
	err = writeResponse(writer, http.StatusOK, data, false)
	return
}

func HttpBadRequestRaw(writer http.ResponseWriter, data string) (err error) {
	err = writeResponse(writer, http.StatusBadRequest, data, false)
	return
}

func writeResponse(writer http.ResponseWriter, httpStatus int, data interface{}, isJson bool) (err error) {
	if isJson {
		writer.Header().Set("Content-Type", "application/json")
	}

	var b []byte

	if isJson {
		b, err = json.Marshal(map[string]interface{}{
			"msg": data,
		})
	} else {
		switch data.(type) {
		case string:
			b = []byte(data.(string))
			break
		default:
			err = errors.New("invalid_type_passed_for_data")
		}
	}

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}

	writer.WriteHeader(httpStatus)
	writer.Write(b)
	return
}
