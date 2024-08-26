package api

import (
	"encoding/json"
	"net/http"
)

// GreetResponse representa la estructura de la respuesta JSON.
type GreetResponse struct {
	Message string `json:"message"`
}

// HelloHandler maneja las solicitudes a la ruta "/".
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := GreetResponse{Message: "Hello, Candy Blog!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
