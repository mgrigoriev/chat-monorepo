run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker image build --tag users .
d-run:
	 docker container run --rm --name users-container -p 8082:8080 users:latest
d-sh:
	docker run --rm -it users sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=users")
