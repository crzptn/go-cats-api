package utilities

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, v any, statusCode int) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
