FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cbwebs ./main.go

FROM scratch
COPY --from=builder /app/cbwebs /usr/bin/cbwebs

ENTRYPOINT ["/usr/bin/cbwebs"]