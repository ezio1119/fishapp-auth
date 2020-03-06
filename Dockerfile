FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src .
RUN go build -o main .

FROM alpine
WORKDIR /app
CMD ["./main"]
COPY --from=builder /src/main .
COPY --from=builder /src/conf/conf.yml /app/conf/conf.yml