FROM golang:1.13 AS build

WORKDIR /home/app

COPY . .

RUN go build -o server main.go

ENTRYPOINT [ "/home/app/server" ]

EXPOSE 8080
