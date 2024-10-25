BINARY := weather-consumer
INPUT_FILES := $(wildcard *.go)

all: ${BINARY}

${BINARY}: ${INPUT_FILES}
	@echo '--------- Building $@.... ------'
	go mod download
	go build -o $@ *.go
	@echo '=============== DONE ==========='

docker: ${BINARY}
	@echo '--------- Building $@.... ------'
	docker build -t weatherconsumer:latest .


clean:
	rm -rf ${BINARY}