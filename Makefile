.PHONY: generate tidy up down start stop sqlc gqlgen

DC = docker compose -f docker-compose.yml
EXEC_MENSATT = $(DC) exec mensatt

up:
	$(DC) up -d

down:
	$(DC) down

start:
	$(DC) start

stop:
	$(DC) stop

generate:
	$(EXEC_MENSATT) go generate ./...

sqlc:
	$(EXEC_MENSATT) go generate ./internal/db/...

gqlgen:
	$(EXEC_MENSATT) go generate ./internal/graphql/...

tidy:
	$(EXEC_MENSATT) go mod tidy
