
DOCKER_CMD=docker
DOCKER_IMAGE=objectdb
DOCKER_IMAGE_VERSION=1.0.0
DOCKERFILE=Dockerfile
# Set the Docker build command
DOCKER_BUILD_CMD=$(DOCKER_CMD) build -t $(DOCKER_IMAGE):$(DOCKER_IMAGE_VERSION) -f $(DOCKERFILE) .


REDIS_CONTAINER_NAME=my-redis-container
REDIS_PORT=6379
REDIS_PASSWORD=pass@123
REDIS_LAUNCH_CMD=$(DOCKER_CMD) run --name $(REDIS_CONTAINER_NAME) -p $(REDIS_PORT):$(REDIS_PORT) -d redis redis-server --requirepass $(REDIS_PASSWORD)


APP_CONTAINER=my-app-container
APP_PORT=8080
LAUNCH_APP_CMD=$(DOCKER_CMD) run -d --name $(APP_CONTAINER) -p $(APP_PORT):$(APP_PORT) --link $(REDIS_CONTAINER_NAME):redis $(DOCKER_IMAGE):$(DOCKER_IMAGE_VERSION)

build:
	@echo "Building Docker image..."
	@$(DOCKER_BUILD_CMD)

redis:
	@echo "Launching redis container..."
	@$(REDIS_LAUNCH_CMD)

deploy:
	@echo "Launching myapp container..."
	@$(LAUNCH_APP_CMD)


all: redis build deploy

clean:
	@echo "Killing containers..."
	@$(DOCKER_CMD) kill $(APP_CONTAINER) $(REDIS_CONTAINER_NAME) || true

	@echo "Removing containers..."
	@$(DOCKER_CMD) rm $(APP_CONTAINER) $(REDIS_CONTAINER_NAME) || true

	@echo "Removing myapp image..."
	@$(DOCKER_CMD) rmi $(DOCKER_IMAGE):$(DOCKER_IMAGE_VERSION) || true


