.PHONY: deploy

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
