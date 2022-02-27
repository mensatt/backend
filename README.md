# mensatt-backend

The backend-repo for "new" sigfood

## Dev Tools

Run gqlgen: `go generate ./internal/graphql/...`  
Run sqlc: `go generate ./internal/db/...`

Run everything: `go generate ./...`

## Dev server

### Database

Install: docker-compose

Run: `docker-compose up -d`  
Stop: `docker-compose down`

### Database Migrations

Install dbmate: https://github.com/amacneil/dbmate

Upgrade: `dbmate up`
Downgrade: `dbmate down`

### Mensatt

Install: `https://github.com/cosmtrek/air`

Run: `air`
