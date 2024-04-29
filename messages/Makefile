run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker image build --tag messages .
d-run:
	 docker container run --rm --name messages-container -p 8080:8080 messages:latest
d-sh:
	docker run --rm -it messages sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=messages")
