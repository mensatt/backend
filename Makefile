.PHONY: generate tidy up down start stop sqlc gqlgen dbmate-new dbmate-migrate dbmate-up dbmate-down dbmate-drop dbmate-status

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

#go code generation
generate:
	$(EXEC_MENSATT) go generate ./...
sqlc:
	$(EXEC_MENSATT) go generate ./internal/db/...
gqlgen:
	$(EXEC_MENSATT) go generate ./internal/graphql/...

# go
tidy:
	$(EXEC_MENSATT) go mod tidy

# dbmate
dbmate-new: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) new
dbmate-migrate: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) migrate
dbmate-up: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) up
dbmate-down: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) down
dbmate-drop: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) drop
dbmate-status: 
	$(EXEC_MENSATT) dbmate $(DBMATE_OPTS) status
