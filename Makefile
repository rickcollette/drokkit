# Go variables
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_FMT=$(GO_CMD) fmt
DISTDIR="dist/"
GO_LINT=golint 

# Project variables
SERVER_NAME=drokkit
GEN_NAME=drokkgen
PKG=./...
DOCKER_IMAGE_NAME=drokkit-server
DOCKER_CONTAINER_NAME=drokkit-container

# Build the application
build: fmt lint
	@echo "Building the application..."
	@$(GO_BUILD) -o $(DISTDIR)/$(SERVER_NAME) main.go
	@$(GO_BUILD) -o $(DISTDIR)/$(GEN_NAME) $(GEN_NAME)/main.go
	@cp runner.sh $(DISTDIR)
	@chmod +x $(DISTDIR)/runner.sh
	@echo "Build complete!"

# Format the code
fmt:
	@echo "Formatting code..."
	@$(GO_FMT) $(PKG)

# Run linter
lint:
	@echo "Running linter..."
	@$(GO_LINT) $(PKG) || echo "Linting issues found. Please review and fix."

# Clean generated files
clean:
	@echo "Cleaning up..."
	@rm -fr $(DISTDIR)/*
	@echo "Cleanup complete!"

# Docker build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME) .
	@echo "Docker image $(DOCKER_IMAGE_NAME) built successfully!"

# Docker run
docker-run:
	@echo "Running Docker container..."
	@docker run -d --name $(DOCKER_CONTAINER_NAME) -p 80:80 $(DOCKER_IMAGE_NAME)
	@echo "Docker container $(DOCKER_CONTAINER_NAME) is running."

# Docker stop and remove container
docker-clean:
	@echo "Stopping and removing Docker container if running..."
	@docker rm -f $(DOCKER_CONTAINER_NAME) || echo "No existing container to remove."
	@echo "Docker cleanup complete!"
