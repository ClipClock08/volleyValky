MIG_DIR := internal/storage/migrations

.PHONY: goose-status

goose-status:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=postgres password=postgres sslmode=disable" status

.PHONY: goose-up

goose-up:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=postgres password=postgres sslmode=disable" up

.PHONY: goose-down

goose-down:
	@cd $(MIG_DIR) && goose postgres "user=postgres dbname=postgres password=postgres sslmode=disable" down
