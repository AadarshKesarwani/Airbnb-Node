.PHONY: migrate-create migrate-up migrate-down migrate-reset migrate-status

MIGRATIONS_DIR=db/migrations
DB_URL=root:shivam99@tcp(127.0.0.1:3306)/auth_dev

migrate-up:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
