FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY src/go.mod .
COPY src/go.sum .
RUN src/go mod download

COPY src .
RUN go build -o main .

FROM alpine
WORKDIR /app
RUN apk add --no-cache tzdata
COPY --from=builder /src/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /src/conf/conf.yml /app/conf/conf.yml

CMD ["./main"]