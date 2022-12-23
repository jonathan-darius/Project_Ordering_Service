# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /WebService

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o webservice /WebService/*.go

EXPOSE 8080

CMD [ "./webservice" ]
