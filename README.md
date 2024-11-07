# weathernotify (WIP)
Notify the bad weather to the subscribed users. It could be rain, it could be snow or it could be a heat wave.

## Pre-requisites
1. Docker
```
Mac : https://docs.docker.com/desktop/install/mac-install/
Linux: https://docs.docker.com/desktop/install/linux/
```

2. confluent-kafka-go (only for local build)
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

cd weathernotify
docker-compose up -d

```

## Test
```
git:(main*) docker logs weathernotify-producer-1
Weather message queued for delivery
Produced event to topic weather: key = forcast    value = [
 {
  "name": "This Afternoon",
  "temperature": 81,
  "windSpeed": "8 mph",
  "shortForecast": "Slight Chance Rain Showers"
 },
 {
  "name": "Tonight",
  "temperature": 58,
  "windSpeed": "2 to 6 mph",
  "shortForecast": "Partly Cloudy then Slight Chance Rain Showers"
 },
 {
  "name": "Thursday",
  "temperature": 73,
  "windSpeed": "2 to 8 mph",
  "shortForecast": "Chance Rain Showers then Partly Sunny"
 },
 {
  "name": "Thursday Night",
  "temperature": 41,
  "windSpeed": "2 to 6 mph",
  "shortForecast": "Mostly Clear"
 },
 {
  "name": "Friday",
  "temperature": 68,
  "windSpeed": "2 to 10 mph",
  "shortForecast": "Sunny"
 },
 {
  "name": "Friday Night",
  "temperature": 36,
  "windSpeed": "3 to 9 mph",
  "shortForecast": "Mostly Clear"
 },
 {
  "name": "Saturday",
  "temperature": 61,
  "windSpeed": "5 mph",
  "shortForecast": "Sunny"
 },
 {
  "name": "Saturday Night",
  "temperature": 38,
  "windSpeed": "3 mph",
  "shortForecast": "Mostly Cloudy"
 },
 {
  "name": "Sunday",
  "temperature": 60,
  "windSpeed": "3 to 7 mph",
  "shortForecast": "Chance Rain Showers"
 },
 {
  "name": "Sunday Night",
  "temperature": 49,
  "windSpeed": "6 mph",
  "shortForecast": "Rain Showers Likely"
 },
 {
  "name": "Veterans Day",
  "temperature": 69,
  "windSpeed": "6 to 10 mph",
  "shortForecast": "Rain Showers Likely"
 },
 {
  "name": "Monday Night",
  "temperature": 43,
  "windSpeed": "3 to 7 mph",
  "shortForecast": "Partly Cloudy"
 },
 {
  "name": "Tuesday",
  "temperature": 64,
  "windSpeed": "7 mph",
  "shortForecast": "Mostly Sunny"
 },
 {
  "name": "Tuesday Night",
  "temperature": 37,
  "windSpeed": "1 to 5 mph",
  "shortForecast": "Mostly Clear"
 }
]


➜  weathernotify
git:(main*) docker logs weathernotify-consumer-1
Consumed event from topic weather: key = forcast    value = [
 {
  "name": "This Afternoon",
  "temperature": 81,
  "windSpeed": "8 mph",
  "shortForecast": "Slight Chance Rain Showers"
 },
 {
  "name": "Tonight",
  "temperature": 58,
  "windSpeed": "2 to 6 mph",
  "shortForecast": "Partly Cloudy then Slight Chance Rain Showers"
 },
 {
  "name": "Thursday",
  "temperature": 73,
  "windSpeed": "2 to 8 mph",
  "shortForecast": "Chance Rain Showers then Partly Sunny"
 },
 {
  "name": "Thursday Night",
  "temperature": 41,
  "windSpeed": "2 to 6 mph",
  "shortForecast": "Mostly Clear"
 },
 {
  "name": "Friday",
  "temperature": 68,
  "windSpeed": "2 to 10 mph",
  "shortForecast": "Sunny"
 },
 {
  "name": "Friday Night",
  "temperature": 36,
  "windSpeed": "3 to 9 mph",
  "shortForecast": "Mostly Clear"
 },
 {
  "name": "Saturday",
  "temperature": 61,
  "windSpeed": "5 mph",
  "shortForecast": "Sunny"
 },
 {
  "name": "Saturday Night",
  "temperature": 38,
  "windSpeed": "3 mph",
  "shortForecast": "Mostly Cloudy"
 },
 {
  "name": "Sunday",
  "temperature": 60,
  "windSpeed": "3 to 7 mph",
  "shortForecast": "Chance Rain Showers"
 },
 {
  "name": "Sunday Night",
  "temperature": 49,
  "windSpeed": "6 mph",
  "shortForecast": "Rain Showers Likely"
 },
 {
  "name": "Veterans Day",
  "temperature": 69,
  "windSpeed": "6 to 10 mph",
  "shortForecast": "Rain Showers Likely"
 },
 {
  "name": "Monday Night",
  "temperature": 43,
  "windSpeed": "3 to 7 mph",
  "shortForecast": "Partly Cloudy"
 },
 {
  "name": "Tuesday",
  "temperature": 64,
  "windSpeed": "7 mph",
  "shortForecast": "Mostly Sunny"
 },
 {
  "name": "Tuesday Night",
  "temperature": 37,
  "windSpeed": "1 to 5 mph",
  "shortForecast": "Mostly Clear"
 }
]
➜  weathernotify
```