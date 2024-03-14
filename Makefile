run:
	@npm ci
	@npm run generate
	@templ generate
	@go run cmd/main.go
