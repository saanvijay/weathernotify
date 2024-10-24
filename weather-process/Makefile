BINARY := weather-notify
INPUT_FILES := $(wildcard *.go)

all: ${BINARY}

${BINARY}: ${INPUT_FILES}
	@echo '--------- Building $@.... ------'
	go build -o $@ *.go
	@echo '=============== DONE ==========='

clean:
	rm -rf ${BINARY}