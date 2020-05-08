### Docker for developers

This is the code I wrote live during the [Go North East April 2020](https://www.meetup.com/Golang-North-East/events/qtpnmqybcgblc/) meeting.

To run:

```
docker-compose up
```

The Docker CLI commands we used during the talk:

```
docker run -it -p 3010:3010 -e PORT=3010 -v `pwd`:/usr/src -w /usr/src golang:1.14 go run main.go
docker run -it -p 3010:3010 -e PORT=3010 -v `pwd`:/usr/src -w /usr/src golang:1.14 go test ./...
```
