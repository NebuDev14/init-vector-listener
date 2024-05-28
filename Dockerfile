FROM golang:1.19

RUN     mkdir /app
WORKDIR /app
ADD     go.mod main.go /app/
RUN     go build

CMD     ./init-vector-listener