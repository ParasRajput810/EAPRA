package provider

import "context"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string
	Messages []Message
}

type Response struct {
	Model   string
	Content string
}

type Provider interface {
	Name() string
	Complete(ctx context.Context, req Request) (Response, error)
}