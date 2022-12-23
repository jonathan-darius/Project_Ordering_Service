# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /OrderingService

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server /OrderingService/Service/*.go

EXPOSE 50052

CMD [ "./server" ]
