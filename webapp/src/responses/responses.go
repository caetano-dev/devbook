package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON returns a JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}
