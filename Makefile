.PHONY: run

run:
	go run cmd/main.go --env=local

migrate-create:
	@read -p  "Migration name (eg:create_users, alter_entities, ...): " NAME; \
	sql-migrate new -config=internal/config/dbconfig.yml -env=local $$NAME

migrate-up:
	sql-migrate up -config=internal/config/dbconfig.yml -env=local

migrate-down:
	sql-migrate down -config=internal/config/dbconfig.yml -env=local

dependencies:
	docker compose -f ./build/dev/docker-compose.yml up -d
