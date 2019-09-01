

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=registry.ng.bluemix.net/${REGISTRY_NAMESPACE}/${APP}:${RELEASE}
GOOS?=linux
GOARCH?=amd64

clean:
		rm -f ${APP}

build: clean
		CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
			-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
			-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
			-o ${APP}

run: container
		docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
		docker run --name ${APP} -p ${PORT}:${PORT} --rm \
			-e "PORT=${PORT}" \
			$(APP):$(RELEASE)

push: container
		docker push $(CONTAINER_IMAGE)
