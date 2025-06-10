# =================================================================
# setup
# =================================================================
.PHONY: setup-env install-tools
setup-env:
	cp -p .env_sample .env

install-tools: setup-env
	go install github.com/air-verse/air@latest


# =================================================================
# local
# =================================================================
.PHONY: lint deps build run
lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o build/shop_api cmd/main.go

run: build
	./build/shop_api


# =================================================================
# container
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d shop-api

down:
	docker compose down shop-api

logs:
	docker compose logs -f shop-api


# =================================================================
# api request
# =================================================================
.PHONY: healthcheck get-users-me
healthcheck:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/healthcheck' \
		-H 'accept: application/json' | jq .

get-me:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/users/me' \
		-H 'accept: application/json' | jq .
