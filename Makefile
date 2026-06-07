BINARY := cats-app
DB := ./database.db
MIGRATIONS := ./sql/migrations

build:
	go build -trimpath -ldflags="-s -w" -o $(BINARY)

install-sqlc:
	@command -v sqlc >/dev/null 2>&1 || { \
		echo "Installing sqlc..."; \
		go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
	}

install-goose:
	@command -v goose >/dev/null 2>&1 || { \
		echo "Installing goose..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	}

tools: install-sqlc install-goose

sqlc: install-sqlc
	sqlc generate

goose-up: install-goose
	goose -dir $(MIGRATIONS) sqlite $(DB) up

goose-down: install-goose
	goose -dir $(MIGRATIONS) sqlite $(DB) down
