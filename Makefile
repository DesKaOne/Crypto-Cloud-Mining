.PHONY: help backend-run backend-test infra-up infra-down infra-ps

help:
	@echo "Targets: backend-run backend-test infra-up infra-down infra-ps"

backend-run:
	cd backend && go run ./api

backend-test:
	cd backend && go test ./...

infra-up:
	docker compose -f infrastructure/docker/compose.yaml up -d --build

infra-down:
	docker compose -f infrastructure/docker/compose.yaml down

infra-ps:
	docker compose -f infrastructure/docker/compose.yaml ps
