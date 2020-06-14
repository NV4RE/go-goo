PROJECT_NAME=go-goo
VERSION=latest
GOCMD=go
BUILD_DIR=./dist
BINARY_NAME=$(PROJECT_NAME)

.PHONY: run

build:
	$(GOCMD) build -o $(BUILD_DIR)/$(BINARY_NAME) -tags netgo
test:
	$(GOCMD) test -v ./...
clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)
run:
	make build
	./$(BUILD_DIR)/$(BINARY_NAME)

build-docker:
	docker build -t $(PROJECT_NAME):$(VERSION) .
push-docker:
	docker push $(PROJECT_NAME):$(VERSION)
run-docker:
	docker run -p $(EXPOSED_PORT):80 $(REPO):$(VERSION)
