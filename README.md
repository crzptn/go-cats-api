# cats

A minimal HTTP API for managing your cats, built with Go.

## Stack

- **Go 1.25** – standard `net/http` (no framework)
- **sqlc** – generates type-safe Go code from SQL queries
- **goose** – database migrations
- **SQLite** – via `modernc.org/sqlite` (pure Go, no CGo)


## Getting started

```bash
# install tools (sqlc, goose)
make tools

# run migrations
make goose-up

# generate Go code from SQL queries
make sqlc

# build the binary
make build

# run the server (default port 6969)
./cats-app
```

The server starts on `127.0.0.1:6969` by default. Set a custom port with `-port`:

```bash
./cats-app -port 8080
```

