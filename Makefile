migration-up:
	./scripts/migrate.sh up

migration-down:
	./scripts/migrate.sh down 1

migration-version:
	./scripts/migrate.sh version

migration-force:
	./scripts/migrate.sh force 1