ifneq (,$(wildcard ./.env))
include .env
export
endif

DATABASE_URL="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"


.DEFAULT_GOAL := run

run:
	go run cmd/main.go

migration-create:
ifndef NAME
	$(error NAME não definido. Ex: make migration-create NAME=create_users_table)
endif
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up

migrate-down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down
