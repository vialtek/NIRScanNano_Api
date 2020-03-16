fmt: ## Format the code
	go fmt ./...

single_scan: 
	go run cmd/scan-cli/main.go
