.PHONY: install update create_migration migrate

include .env
schema=sql/schema.prisma

install:
	@go install github.com/steebchen/prisma-client-go@latest github.com/sqlc-dev/sqlc/cmd/sqlc@latest
update:
	@go get -u ./...
create_migration:
	@prisma-client-go migrate dev --schema=$(schema)
migrate:
	@prisma-client-go migrate deploy --schema=$(schema)
generate:
	@go generate ./...
