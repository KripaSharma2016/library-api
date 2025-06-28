FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o library-app .

EXPOSE 8080

CMD [ "./library-app" ]
