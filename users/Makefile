.PHONY: deploy

run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
deploy:
	$(MAKE) d-build
	$(MAKE) d-push
	kubectl --context minikube --namespace main apply -f ./deploy
	kubectl --context minikube rollout restart -n main deployment users
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker build --tag mgrigoriev/users .
d-push:
	docker push mgrigoriev/users
d-run:
	 docker container run --rm --name users-container -p 8083:8080 users:latest
d-sh:
	docker run --rm -it users sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=users")
