# weathernotify
Notify the bad weather to the subscribed users. It could be rain, it could be snow or it could be a heat wave.

## Build
```
make docker
docker run -p 8080:8080 weathernotify:latest
```

## Test
```
curl localhost:8080/getlocation
curl localhost:8080/getcurrentlocationforecast
curl localhost:8080/getforecast/42/83
```