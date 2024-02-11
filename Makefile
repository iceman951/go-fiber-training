include .env
export

# init:
# 	git config core.hooksPath .githooks

hello:
	echo "Hello"

update:
	go mod tidy

air: update ## HOT RELOAD MODE ##
	air

test:
	go test ./pkg/... -coverprofile coverage.unit.out -run 'Test([^I]|I[^n]|In[^t]|Int[^e]|Inte[^g]|Integ[^r]|Integr[^a]|Integra[^t]|Integrat[^i]|Integrati[^o]|Integratio[^n]).*'

build:
	go build -o ./bin/${PROJECT_NAME} main.go

TAG := $(shell jq -r '.Version' ${CONFIGPATH})
ENV := $(shell jq -r '.ENV' ${CONFIGPATH})
prepare-env:
	echo TAG=${TAG} >> .env
	echo ENV=${ENV} >> .env

build-docker-image:
	echo "docker build --build-arg APP_VERSION=${TAG} -f Dockerfile . -t ${DOCKER_REPO}/${WORKSPACE_NAME}/${PROJECT_NAME}:${TAG}"

push-docker-image:
	docker login ${DOCKER_REPO} -u ${DOCKER_REPO_USER} -p ${DOCKER_REPO_PASS}
	docker push ${DOCKER_REPO}/${WORKSPACE_NAME}/${PROJECT_NAME}:${TAG}

restart-container:
	docker-compose stop ${PROJECT_NAME}
	docker-compose rm -f ${PROJECT_NAME}
	docker-compose up -d ${PROJECT_NAME}
