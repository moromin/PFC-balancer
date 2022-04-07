MIGRATION_PATH	= db/migration
FOOD_DB_PATH = db/food

migrateup:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:password@localhost:5432/balance-app?sslmode=disable" -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:password@localhost:5432/balance-app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

food-db:
	go1.18 run db/food/main.go

.PHONY: migrateup migratedown sqlc
