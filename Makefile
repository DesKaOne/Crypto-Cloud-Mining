.PHONY: help backend-run backend-test infra-up infra-down

help:
	@echo "Targets: backend-run backend-test infra-up infra-down"

backend-run:
	cd backend && go run ./api

backend-test:
	cd backend && go test ./...

infra-up:
	docker compose -f infrastructure/docker/compose.yaml up -d

infra-down:
	docker compose -f infrastructure/docker/compose.yaml down
