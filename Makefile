build:
	dep ensure
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/linux/medusa cmd/medusa/main.go
	#env GOARCH=amd64 GOOS=windows go build -ldflags="-s -w" -o bin/win/medusa.exe cmd/medusa/main.go
	#env GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o bin/osx/medusa cmd/medusa/main.go
