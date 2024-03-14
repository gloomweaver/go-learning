run:
	@npm ci
	@npm run generate
	@templ generate
	@go build cmd/main.go
