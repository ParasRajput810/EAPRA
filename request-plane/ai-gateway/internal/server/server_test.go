package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eapra/eapra/request-plane/ai-gateway/internal/provider"
	"github.com/eapra/eapra/request-plane/ai-gateway/internal/provider/stub"
)

type shouty struct{}

func (shouty) Name() string { return "shouty" }
func (shouty) Complete(_ context.Context, req provider.Request) (provider.Response, error) {
	last := req.Messages[len(req.Messages)-1].Content
	return provider.Response{Model: req.Model, Content: strings.ToUpper(last) + "!!!"}, nil
}

func call(t *testing.T, p provider.Provider, body string) (*httptest.ResponseRecorder, map[string]any) {
	t.Helper()
	r := httptest.NewRequest(http.MethodPost, "/v1/chat/completions", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	(&Server{Provider: p}).Routes().ServeHTTP(w, r)
	var got map[string]any
	json.Unmarshal(w.Body.Bytes(), &got)
	return w, got
}

const okBody = `{"model":"demo","messages":[{"role":"user","content":"hello there"}]}`

func TestHandlerIsProviderAgnostic(t *testing.T) {
	_, a := call(t, stub.New(), okBody)
	if a["provider"] != "stub" {
		t.Errorf("stub: %v", a)
	}

	_, b := call(t, shouty{}, okBody)
	if b["provider"] != "shouty" || b["content"] != "HELLO THERE!!!" {
		t.Errorf("shouty: %v", b)
	}
}

func TestMissingMessagesIsRejected(t *testing.T) {
	if w, _ := call(t, stub.New(), `{"model":"demo"}`); w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}

var _ provider.Provider = (*stub.Provider)(nil)
var _ provider.Provider = shouty{}