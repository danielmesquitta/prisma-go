.PHONY:

include .env
schema=sql/schema.prisma


default: rest_dev

install:
	@go install -u github.com/steebchen/prisma-client-go@latest
update:
	@go get -u ./...
create_migration:
	@prisma-client-go migrate dev --schema=$(schema)
migrate:
	@prisma-client-go migrate deploy --schema=$(schema)
