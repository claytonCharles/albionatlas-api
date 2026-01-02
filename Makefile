ifneq (,$(wildcard ./.env))
include .env
export
endif

DATABASE_URL="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"

migration-create:
ifndef NAME
	$(error NAME não definido. Ex: make migration-create NAME=create_users_table)
endif
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(POSTGRES_URL) up
