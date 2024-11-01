# weathernotify (WIP)
Notify the bad weather to the subscribed users. It could be rain, it could be snow or it could be a heat wave.

## Pre-requisites
1. Docker
```
Mac : https://docs.docker.com/desktop/install/mac-install/
Linux: https://docs.docker.com/desktop/install/linux/
```

2. confluent-kafka-go
```
go get github.com/confluentinc/confluent-kafka-go/kafka
Mac: brew install librdkafka 
Linux: sudo apt-get install librdkafka-dev
```

## Build
```
cd weatherprocess
make docker

cd weatherconsumer
make docker
//docker run -p 8080:8080 weathernotify:latest
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes

cd weathernotify
docker-compose up -d

```

## Test
```
//curl localhost:8080/getlocation
//curl localhost:8080/getcurrentlocationforecast
//curl localhost:8080/getforecast/42/83
```