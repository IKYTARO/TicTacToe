run:
	go run cmd/server/main.go

test:
	go test ./internal/... -v

demo:
	@echo "Starting server in background..."
	@echo
	@go run cmd/server/main.go &
	@sleep 2
	@echo
	@bash examples/demo.sh
	@kill $$(lsof -t -i :8080) 2>/dev/null || true

.PHONY: run test demo
