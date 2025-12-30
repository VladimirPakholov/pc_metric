ENV_FILE := .env
run:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "'.env' file not found, creating..."; \
		echo "DEFAULT_TIME_WORK=30s" > .env; \
		echo "DEFAULT_TIME_GET_METRIC=1s" >> .env; \
		echo "DB_NAME=test" >> .env; \
		echo "DB_USER=postgres" >> .env; \
		echo "DB_PASSWORD=postgres" >> .env; \
		echo "DB_PORT=5432" >> .env; \
		echo "DB_SSLMODE=disable" >> .env; \
		echo "DB_HOST=localhost" >> .env; \
	fi
	@set -a; \
	. ./$(ENV_FILE); \
	go run ./cmd/cli/main.go
