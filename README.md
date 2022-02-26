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

todo: run first database seeding

### Mensatt

Install: `https://github.com/cosmtrek/air`

Run: `air`
