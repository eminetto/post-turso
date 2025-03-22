tests: generate-mocks
	@go test ./...
generate-mocks:
	@go tool mockery --output book/mocks --dir book --all