#!/bin/bash

DB_URL=$(go run ./cmd/migrate)

migrate \
-path migrations \
-database "$DB_URL" \
"$@"