FROM golang:1.20 AS builder

RUN apk add --no-cache git ca-certificates

ARG GOPROXY=https://proxy.golang.org,direct

ENV GO111MODULE=on \
  GOPROXY=${GOPROXY} \
  GOPRIVATE=${GOPRIVATE}

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./app"]
