.PHONY: run

run:
	go run cmd/main.go --env=local

local:
	docker compose -f ./build/dev/docker-compose.yml up -d
	go run cmd/main.go --env=local

migrate-create:
	@read -p  "Migration name (eg:create_users, alter_entities, ...): " NAME; \
	sql-migrate new -config=internal/config/dbconfig.yml -env=local $$NAME
