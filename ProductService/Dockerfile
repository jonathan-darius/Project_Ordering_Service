# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /ProductService

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server /ProductService/ProductService/*.go

EXPOSE 50052

CMD [ "./server" ]
