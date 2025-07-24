FROM golang:1.24

WORKDIR /app

COPY . /app/

RUN go build -o main ./cmd/e-commerce/main.go

EXPOSE 8082

ENV CONFIG_PATH=./config/local.yaml

ENTRYPOINT ["./main"]

