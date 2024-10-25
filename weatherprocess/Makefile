BINARY := weatherproducer
INPUT_FILES := $(wildcard *.go)

all: ${BINARY}

${BINARY}: ${INPUT_FILES}
	@echo '--------- Building $@.... ------'
	go mod download
	go build -o $@ *.go
	@echo '=============== DONE ==========='

docker: ${BINARY}
	@echo '--------- Building $@.... ------'
	docker build --platform linux/amd64 --no-cache -t weatherproducer:latest .


clean:
	rm -rf ${BINARY}