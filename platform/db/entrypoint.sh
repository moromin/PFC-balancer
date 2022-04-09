#!/bin/sh

/wait
/migrate \
    -path $MIGRATIONS_DIR \
    -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
    $@
