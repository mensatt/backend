.PHONY: gen generate tidy log logs up down start stop sqlc gqlgen fmt format jwt jwt-keys

DC = docker compose -f docker-compose.yml

# docker compose stuff
build:
	$(DC) build
up:
	$(DC) up -d
down:
	$(DC) down
start:
	$(DC) start
stop:
	$(DC) stop
log: logs
logs:
	$(DC) logs -f mensatt

# go code generation
gen: generate
generate:
	go generate ./...
sqlc:
	go generate ./internal/db/...
gqlgen:
	go generate ./internal/graphql/...

# go format
fmt: format
format:
	go fmt ./...

# go
tidy:
	go mod tidy

jwt: jwt-keys
jwt-keys:
	openssl genrsa -out ./.secrets/jwt_private_key.pem 2048
	openssl rsa -in ./.secrets/jwt_private_key.pem -pubout > ./.secrets/jwt_public_key.pem
