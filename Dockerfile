FROM golang:1.14 AS base

WORKDIR /usr/src
COPY . .

RUN go build -o /helloweb

FROM ubuntu

COPY --from=base /helloweb /helloweb

CMD ["/helloweb"]
