MIGRATION_PATH	= db/migration

migrateup:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:password@localhost:5432/balance-app?sslmode=disable" -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:password@localhost:5432/balance-app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc
