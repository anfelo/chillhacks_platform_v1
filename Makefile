.PHONY: postgres adminer migrate platform

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

platform:
	export PGUSER="postgres" && \
	export PGHOST="localhost" && \
	export PGDATABASE="postgres" && \
	export PGPASSWORD="secret" && \
	export JWTSECRET="jwtsecret" && \
	reflex -s go run main.go

migrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down