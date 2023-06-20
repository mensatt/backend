dc := "docker compose -f docker-compose.yml"

build:
    {{dc}} build

up:
    {{dc}} up -d

down:
    {{dc}} down

start:
    {{dc}} start

stop:
    {{dc}} stop

alias log := logs
logs:
    {{dc}} logs -f mensatt

# go code generation
alias gen := generate
generate:
    go generate ./...

entc:
    go generate ./internal/database/...

gqlgen:
	go generate ./internal/graphql/...

# go format
alias fmt := format
format:
	go fmt ./...

# go
tidy:
	go mod tidy

alias jwt := jwt-keys
jwt-keys:
	openssl genrsa -out ./.secrets/jwt_private_key.pem 2048
	openssl rsa -in ./.secrets/jwt_private_key.pem -pubout > ./.secrets/jwt_public_key.pem
