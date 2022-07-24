# Mensatt Backend

[![Build & Test](https://github.com/mensatt/backend/actions/workflows/go.yml/badge.svg)](https://github.com/mensatt/backend/actions/workflows/go.yml)
[![Development Deployment](https://github.com/mensatt/backend/actions/workflows/deploy-main-in-dev-env.yml/badge.svg)](https://github.com/mensatt/backend/actions/workflows/deploy-main-in-dev-env.yml)

The backend repository for Mensatt written in go.

## Requirements
- Docker & Docker Compose
- GCC (build-essentials)
- Go 1.18

## Building

Generate all: ```make generate```  
Run only gqlgen: ```make gqlgen```  
Run only sqlc: ```make sqlc```  

## Running

Start dev server: ```make up```  
Stop the server: ```make down```  
