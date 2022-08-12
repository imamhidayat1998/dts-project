package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}
