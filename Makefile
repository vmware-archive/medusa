build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/medusa cmd/medusa/main.go
