TIMESTAMP=$(shell date +%Y%m%d%H%M%S)
MIGRATION_PATH=./internal/sql/migrations

.PHONY: new-migration

new-migration: ## Generate new migration file. example: `make new-migration name=create_foobar_table`
	@touch $(MIGRATION_PATH)/$(TIMESTAMP)_$(name).tx.up.sql
