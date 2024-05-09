.PHONY: deploy

run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
deploy:
	$(MAKE) d-build
	$(MAKE) d-push
	kubectl --context minikube --namespace main apply -f ./deploy
	kubectl --context minikube rollout restart -n main deployment chatmessages
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker build --tag mgrigoriev/chatmessages .
d-push:
	docker push mgrigoriev/chatmessages
d-run:
	 docker container run --rm --name chatmessages-container -p 8081:8080 chatmessages:latest
d-sh:
	docker run --rm -it chatmessages sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=chatmessages")
