run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker image build --tag servers .
d-run:
	 docker container run --rm --name servers-container -p 8081:8080 servers:latest
d-sh:
	docker run --rm -it servers sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=servers")
