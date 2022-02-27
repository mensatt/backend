.PHONY: generate tidy up

DC = docker-compose -f docker-compose.yml

up:
	$(DC) up -d

generate:
	$(DC) exec -it go generate ./...

tidy:
	$(DC) exec -it go mod tidy
