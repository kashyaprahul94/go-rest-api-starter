#---------------------#

# To run the project for local dev environment
dev:
	godotenv -f ./.env go run ./cmd/rest-api-starter/rest-api.go

#---------------------#

# To run all the tests
test:
	godotenv -f ./.env.test go test ./pkg/* -v

#---------------------#

# To install all dependencies
install:
	go mod download

#---------------------#

# To compile the go binary
compile:
	go build -o bin/rest-api ./cmd/rest-api-starter/rest-api.go

#---------------------#

# To removed unused dependencies
tidy:
	go mod tidy

#---------------------#

# To build the project, remove any unused deps, and compile the project
build: install tidy compile

#---------------------#


#---------------------#
# .PHONY
.phony: dev test compile tidy build
#---------------------#