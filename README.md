# Mensatt Backend

The backend repository for Mensatt written in go.

## Requirements
- Docker & Docker Compose
- GCC (build-essentials)
- Go 1.17

## Building

Generate all: ```make generate```  
Run only gqlgen: ```make gqlgen```  
Run only sqlc: ```make sqlc```  

## Running

Start dev server: ```make up```  
Stop the server: ```make down```  
