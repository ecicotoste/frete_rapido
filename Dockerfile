FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /main

ENTRYPOINT [ "/main" ]