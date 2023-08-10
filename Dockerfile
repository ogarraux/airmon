FROM golang:1.19-alpine

ADD . /app
WORKDIR /app
RUN go build -o /air-monitor .

ENTRYPOINT ["/air-monitor"]
