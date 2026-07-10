package stub

import (
	"context"
	"fmt"
	"strings"

	"github.com/eapra/eapra/request-plane/ai-gateway/internal/provider"
)

type Provider struct{}

func New() *Provider { return &Provider{} }

func (p *Provider) Name() string { return "stub" }

func (p *Provider) Complete(_ context.Context, req provider.Request) (provider.Response, error) {
	last := req.Messages[len(req.Messages)-1].Content
	content := fmt.Sprintf("[stub] you said: %q", last)

	
	var prompt strings.Builder
	for _, m := range req.Messages {
		prompt.WriteString(m.Role)
		prompt.WriteString(m.Content)
	}

	return provider.Response{
		Model:        req.Model,
		Content:      content,
		InputTokens:  ApproxTokens(prompt.String()),
		OutputTokens: ApproxTokens(content),
	}, nil
}

func ApproxTokens(s string) int {
	if s == "" {
		return 0
	}
	if n := len(s) / 4; n > 0 {
		return n
	}
	return 1
}