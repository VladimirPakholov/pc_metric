ENV_FILE := .env
run:
	@if [ ! -f $(ENV_FILE) ]; then \
	echo "'.env' file not found, creating..."; \
	echo "DEFAULT_TIME_WORK=30s" > .env; \
	echo "DEFAULT_TIME_GET_METRIC=1s" >> .env; \
	fi
	@set -a; \
	. ./$(ENV_FILE); \
	go run ./cmd/cli/main.go
