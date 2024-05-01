.PHONY: deploy

run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
deploy:
	$(MAKE) d-build
	$(MAKE) d-push
	kubectl --context minikube --namespace main apply -f ./deploy
	kubectl --context minikube rollout restart -n main deployment servers
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker build --tag mgrigoriev/servers .
d-push:
	docker push mgrigoriev/servers
d-run:
	 docker container run --rm --name servers-container -p 8082:8080 servers:latest
d-sh:
	docker run --rm -it servers sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=servers")
