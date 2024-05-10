.PHONY: deploy

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/api

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# генерация .go файлов с помощью protoc
.protoc-generate:
	mkdir -p $(PKG_PROTO_PATH)
	$(PROTOC) --proto_path=$(CURDIR) \
	          --go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	          --go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .protoc-generate .tidy

# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/client
	go build -o $(LOCAL_BIN) ./cmd/server

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
