# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /UserServices

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server /UserServices/UserService/*.go

EXPOSE 50051

CMD [ "./server" ]
