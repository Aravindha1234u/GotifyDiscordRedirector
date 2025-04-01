# Project configuration
BINARY_NAME=GotifyDiscordRedirector
DOCKER_IMAGE=aravindha1234u/gotifydiscordredirector
DOCKER_TAG=latest

# Go build settings
GOOS_LIST=linux darwin windows
GOARCH_LIST=amd64 arm64

# Docker build settings
DOCKERFILE=Dockerfile

# Go build targets
build-go:
	@echo "Building Go binaries..."
	@for os in $(GOOS_LIST); do \
		for arch in $(GOARCH_LIST); do \
			echo "Building $(BINARY_NAME)-$$os-$$arch..."; \
			if [ "$$os" = "windows" ]; then \
				GOOS=$$os GOARCH=$$arch go build -o bin/$(BINARY_NAME)-$$os-$$arch.exe . ; \
			else \
				GOOS=$$os GOARCH=$$arch go build -o bin/$(BINARY_NAME)-$$os-$$arch . ; \
			fi; \
			echo "Completed"; \
		done; \
	done

# Docker build target
build-docker:
	@echo "Building Docker image..."
	sudo docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) -f $(DOCKERFILE) .

# Docker push target
push-docker:
	@echo "Pushing Docker image..."
	sudo docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

# Clean target
clean:
	@echo "Cleaning up..."
	rm -rf bin/

	sudo docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)

# All target
all: clean build-go build-docker push-docker

.PHONY: build-go build-docker push-docker clean all all-push