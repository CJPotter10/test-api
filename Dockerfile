FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /test-api

EXPOSE 8080

CMD [ "/test-api" ]