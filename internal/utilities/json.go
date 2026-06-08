package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RespondJson(w http.ResponseWriter, v any, statusCode int) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func ReadJson(r *http.Request, v any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Unexpected Content-Type %s", r.Header.Get("Content-Type"))
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
