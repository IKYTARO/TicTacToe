package web

import (
	"TicTacToe/internal/domain/model"
	"TicTacToe/internal/domain/service"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type GameHandler struct {
	service service.GameService
}

func NewGameHandler(service service.GameService) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

func (handler *GameHandler) ServeUI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/web/static/index.html")
}

func (handler *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "only POST method is allowed")
		return
	}

	game := handler.service.CreateGame()
	response := DomainToResponse(game, model.InProgress)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (handler *GameHandler) HandleGame(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/game/" && r.Method == http.MethodPost {
		handler.CreateGame(w, r)
		return
	}

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "only POST method is allowed")
		return
	}

	uuidString := extractUUID(r.URL.Path)
	if uuidString == "" {
		writeError(w, http.StatusBadRequest, "missing game UUID in URL")
		return
	}

	var request GameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	current, err := RequestToDomain(uuidString, &request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := handler.service.ProcessGame(current)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := DomainToResponse(current, result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func extractUUID(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 2 && parts[0] == "game" {
		return parts[1]
	}
	return ""
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
