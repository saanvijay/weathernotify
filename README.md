# weathernotify
Notify the weather forcast to the email users which are listed in .env file. The forcast data will be fetched for every 12 hour. The location is your current location.

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

## Create App password for gmail
In https://myaccount.google.com/security, do you see 2-step verification set to ON? If yes, then visiting https://myaccount.google.com/apppasswords should allow you to set up application specific passwords. 

## Build
```
Edit .env and add email_ids and app password

FROM_EMAIL_ID="from_email_id"
FROM_EMAIL_APP_PASS="from_email_id_passwd"
TO_EMAIL_ID="recipient_email_id"
EMAIL_NOTIFICATION_IN_MINUTES="720" // 720 minutes -> 12 hour

cd weatherprocess
make docker

cd weatherconsumer
make docker

cd weathernotify
docker-compose up -d

```
## Email Notification
The notification of the weather forcast will be sent to the recipient email which is in .env

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

## Enhancements
1. Register the emails in Database
2. Send the notification for the registered email addressess
3. Add authentication for REST server
4. Kafka authentication

## License
MIT License

Copyright (c) 2024 Vijaya Prakash Masilamani

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
