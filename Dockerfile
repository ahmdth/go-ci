FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]
