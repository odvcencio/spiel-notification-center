CURRENT_HEAD = $$(git rev-parse HEAD)
REGISTRY_URL = registry.tryspiel.com
PROJECT_NAME = spiel/notification-center
KUBE_DEPLOYMENT = deployments/notification-center-deployment
KUBE_CONTAINER_NAME = spiel-notification-center
STAGE = dev

binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/notification-center

docker-build:
	@docker build --force-rm \
	              --tag ${PROJECT_NAME}:${CURRENT_HEAD} \
	              --file ./Dockerfile .

docker-tag:
	@docker tag ${PROJECT_NAME}:${CURRENT_HEAD} ${PROJECT_NAME}:${STAGE}
	@docker tag ${PROJECT_NAME}:${STAGE} ${REGISTRY_URL}/${PROJECT_NAME}:${STAGE}
	@docker tag ${PROJECT_NAME}:${STAGE} ${REGISTRY_URL}/${PROJECT_NAME}:${CURRENT_HEAD}

docker-push:
	@docker push ${REGISTRY_URL}/${PROJECT_NAME}:${CURRENT_HEAD}
	@docker push ${REGISTRY_URL}/${PROJECT_NAME}:${STAGE}

docker-deploy:
	@kubectl set image ${KUBE_DEPLOYMENT} ${KUBE_CONTAINER_NAME}=${REGISTRY_URL}/${PROJECT_NAME}:${CURRENT_HEAD}

docker-all:
	@make docker-build
	@make docker-tag
	@make docker-push
	@make docker-deploy
