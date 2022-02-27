# Mensatt Backend

The backend repository for Mensatt (neosigfood)

## Dev Tools

Run gqlgen: `go generate ./internal/graphql/...`  
Run sqlc: `go generate ./internal/db/...`

Run everything: `go generate ./...`

## Dev server

### Database

Install: docker & docker-compose
`curl -sSL https://get.docker.com/ | CHANNEL=stable bash`

Run: `docker-compose up -d`  
Stop: `docker-compose down`

### Database Migrations

Install dbmate: https://github.com/amacneil/dbmate

Upgrade: `dbmate up`
Downgrade: `dbmate down`

### Mensatt

Install air: `https://github.com/cosmtrek/air`

Run: `air`
