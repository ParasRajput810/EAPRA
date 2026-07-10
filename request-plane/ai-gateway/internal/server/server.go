package server

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Server struct{}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("POST /v1/chat/completions", s.handleChat)
	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleChat(w http.ResponseWriter, r *http.Request) {
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errBody("invalid_request", "malformed JSON body"))
		return
	}
	if req.Model == "" || len(req.Messages) == 0 {
		writeJSON(w, http.StatusBadRequest, errBody("invalid_request", "model and messages are required"))
		return
	}

	// The model, hardcoded. The gateway KNOWS what it is. It cannot be swapped.
	last := req.Messages[len(req.Messages)-1].Content
	writeJSON(w, http.StatusOK, map[string]any{
		"model":   req.Model,
		"content": "you said: " + last,
	})
}

func errBody(kind, msg string) map[string]any {
	return map[string]any{"error": map[string]string{"type": kind, "message": msg}}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}