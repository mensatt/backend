.PHONY: gen generate tidy log logs up down start stop sqlc gqlgen jwt jwt-keys

DC = docker compose -f docker-compose.yml
EXEC_MENSATT = $(DC) exec mensatt
DBMATE_OPTS = --migrations-dir ./internal/db/sql/migrations --schema-file ./internal/db/sql/schema.sql

# docker compose stuff
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
	$(DC) logs mensatt

#go code generation
gen: generate
generate:
	go generate ./...
sqlc:
	go generate ./internal/db/...
gqlgen:
	go generate ./internal/graphql/...

# go
tidy:
	go mod tidy

jwt: jwt-keys
jwt-keys:
	openssl genrsa -out ./.secrets/jwt_private_key.pem 2048
	openssl rsa -in ./.secrets/jwt_private_key.pem -pubout > ./.secrets/jwt_public_key.pem
