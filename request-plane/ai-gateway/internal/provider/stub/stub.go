package stub

import (
	"context"
	"fmt"

	"github.com/eapra/eapra/request-plane/ai-gateway/internal/provider"
)

type Provider struct{}

func New() *Provider { return &Provider{} }

func (p *Provider) Name() string { return "stub" }

func (p *Provider) Complete(_ context.Context, req provider.Request) (provider.Response, error) {
	last := req.Messages[len(req.Messages)-1].Content
	return provider.Response{
		Model:   req.Model,
		Content: fmt.Sprintf("[stub] you said: %q", last),
	}, nil
}