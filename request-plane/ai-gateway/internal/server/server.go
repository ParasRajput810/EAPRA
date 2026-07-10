package server

import (
	"encoding/json"
	"net/http"

	"github.com/eapra/eapra/request-plane/ai-gateway/internal/provider"
)

type chatRequest struct {
	Model    string             `json:"model"`
	Messages []provider.Message `json:"messages"`
}

type Server struct {
	Provider provider.Provider
}

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

	resp, err := s.Provider.Complete(r.Context(), provider.Request{Model: req.Model, Messages: req.Messages})
	if err != nil {
		writeJSON(w, http.StatusBadGateway, errBody("provider_error", err.Error()))
		return
	}

	total := resp.InputTokens + resp.OutputTokens

	writeJSON(w, http.StatusOK, map[string]any{
		"provider": s.Provider.Name(),
		"model":    resp.Model,
		"content":  resp.Content,
		"usage": map[string]int{
			"input_tokens":  resp.InputTokens,
			"output_tokens": resp.OutputTokens,
			"total_tokens":  total,
		},
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