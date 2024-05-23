.PHONY: deploy

LOCAL_BIN := $(CURDIR)/bin
MIGRATION_DIR :=./migrations

run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
deploy:
	$(MAKE) d-build
	$(MAKE) d-push
	kubectl --context minikube --namespace main apply -f ./deploy
	kubectl --context minikube rollout restart -n main deployment chatservers
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker build --tag mgrigoriev/chatservers .
d-push:
	docker push mgrigoriev/chatservers
d-run:
	 docker container run --rm --name chatservers-container -p 8082:8080 chatservers:latest
d-sh:
	docker run --rm -it chatservers sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=chatservers")

# Migrations
.install-migrate: export GOBIN := $(LOCAL_BIN)
.install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create-migration: .install-migrate
create-migration:
	PATH=$(LOCAL_BIN) migrate create -ext sql -dir "${MIGRATION_DIR}" -seq migration

DB_DSN:=postgresql://mikhail@0.0.0.0:5432/chatservers?sslmode=disable

migrate:
	PATH=$(LOCAL_BIN) migrate -path ${MIGRATION_DIR} -database "${DB_DSN}" -verbose up

migrate-rollback:
	PATH=$(LOCAL_BIN) migrate -path ${MIGRATION_DIR} -database "${DB_DSN}" -verbose down
