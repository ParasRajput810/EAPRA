GATEWAY := request-plane/ai-gateway
ADDR    := localhost:8080

.DEFAULT_GOAL := help
.PHONY: run test 
run: 
	cd $(GATEWAY) && go run ./cmd/gateway

test:
	cd $(GATEWAY) && go test ./...

