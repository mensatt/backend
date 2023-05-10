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
