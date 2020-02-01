dev:
	godotenv -f ./.env go run ./cmd/rest-api-starter/rest-api.go

test:
	godotenv -f ./.env.test go test ./pkg/* -v

build:
	go build -o bin/rest-api ./cmd/rest-api-starter/rest-api.go