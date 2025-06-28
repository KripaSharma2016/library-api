# -------- VARIABLES --------
APP_NAME = library-app
DOCKER_IMAGE = <user>/$(APP_NAME)
K8S_DIR = k8s

# -------- BUILD COMMANDS --------

.PHONY: all
all: swag build

swag:
	@echo "Generating Swagger docs..."
	swag init -g main.go --output docs

build:
	@echo "Building Go binary..."
	go build -o main .

run:
	@echo "Running app locally on :8080"
	go run main.go

test:
	@echo "Running tests..."
	go test ./... -v

# -------- DOCKER COMMANDS --------

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):latest .

docker-push:
	@echo "Pushing Docker image to Docker Hub..."
	docker push $(DOCKER_IMAGE):latest

docker-run:
	@echo "Running Docker container locally..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE):latest

# -------- KUBERNETES COMMANDS --------

k8s-deploy:
	@echo "Deploying to Kubernetes..."
	kubectl apply -f $(K8S_DIR)

k8s-delete:
	@echo "Removing Kubernetes resources..."
	kubectl delete -f $(K8S_DIR)

k8s-logs:
	kubectl logs -l app=library-api

k8s-status:
	kubectl get all

# -------- CLEAN --------

clean:
	@echo "Cleaning build artifacts..."
	rm -f main