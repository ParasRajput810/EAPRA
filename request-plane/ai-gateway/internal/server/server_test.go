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
	c := strings.ToUpper(last) + "!!!"
	return provider.Response{Model: req.Model, Content: c, InputTokens: 1, OutputTokens: 1}, nil
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

func TestResponseReportsUsage(t *testing.T) {
	_, got := call(t, stub.New(), okBody)
	usage, ok := got["usage"].(map[string]any)
	if !ok {
		t.Fatal("every response must report usage")
	}
	in := usage["input_tokens"].(float64)
	out := usage["output_tokens"].(float64)
	tot := usage["total_tokens"].(float64)
	if in == 0 || out == 0 {
		t.Error("tokens must be counted")
	}
	if tot != in+out {
		t.Errorf("total %v != %v + %v", tot, in, out)
	}
}

func TestLongerPromptsCostMore(t *testing.T) {
	p := stub.New()
	short, _ := p.Complete(context.Background(), provider.Request{
		Model: "m", Messages: []provider.Message{{Role: "user", Content: "hi"}}})
	long, _ := p.Complete(context.Background(), provider.Request{
		Model: "m", Messages: []provider.Message{{Role: "user",
			Content: "this is a considerably longer prompt that costs a great deal more to process"}}})

	if long.InputTokens <= short.InputTokens {
		t.Fatalf("long prompt (%d tokens) must cost more than short (%d) -- yet both are ONE request",
			long.InputTokens, short.InputTokens)
	}
}