# Mensatt Backend

[![Build & Test](https://github.com/mensatt/backend/actions/workflows/go.yml/badge.svg)](https://github.com/mensatt/backend/actions/workflows/go.yml)
[![Development Deployment](https://github.com/mensatt/backend/actions/workflows/deploy-main-in-dev-env.yml/badge.svg)](https://github.com/mensatt/backend/actions/workflows/deploy-main-in-dev-env.yml)

The backend repository for [Mensatt](https://www.mensatt.de) written in go.

## Requirements

- Docker & Docker Compose
- Go 1.19
- libvips-dev

## Building

- Generate all: ```make generate```
- Run only gqlgen: ```make gqlgen```
- Run only entc: ```make entc```

## Running

- Start dev server: ```make up```
- Stop the server: ```make down```

## Migrations

Migrations require Atlas to be installed. See [here](https://entgo.io/docs/versioned-migrations/#quick-guide) for installation instructions.

Create a new migration:
```bash
atlas migrate diff <migration_name> \
 --dir "file://internal/database/migrations" \
 --to "ent://internal/database/schema" \
 --dev-url "postgres://mensatt:mensatt@localhost:5432/mensatt?search_path=public&sslmode=disable" \
 --format '{{ sql . "	" }}'
```
The `--format` flag uses a tab as indentation. If you want to manually edit this command you can use
`ctrl + v + tab` to insert a tab character in your shell.

Applying the migration:
```bash
atlas migrate apply \
 --dir file://internal/database/migrations \
 --url "postgres://mensatt:mensatt@localhost:5432/mensatt?search_path=public&sslmode=disable"
```
