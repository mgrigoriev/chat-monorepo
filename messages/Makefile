.PHONY: deploy

run:
	cd ./cmd && go run .
test:
	go test -v -count=1 ./...
deploy:
	$(MAKE) d-build
	$(MAKE) d-push
	kubectl --context minikube --namespace main apply -f ./deploy
	kubectl --context minikube rollout restart -n main deployment messages
d:
	$(MAKE) d-build
	$(MAKE) d-run
d-build:
	docker build --tag mgrigoriev/messages .
d-push:
	docker push mgrigoriev/messages
d-run:
	 docker container run --rm --name messages-container -p 8081:8080 messages:latest
d-sh:
	docker run --rm -it messages sh
d-log:
	docker logs -f $(shell docker ps -qf "ancestor=messages")
