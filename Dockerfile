FROM golang:1.19-alpine as builder

WORKDIR /app

COPY main.go ./main.go

RUN CGO_ENABLED=0 go build -o main -ldflags="-s -w" main.go

FROM gcr.io/distroless/static-debian11 as runtime

COPY --from=builder /app/main /

EXPOSE 8080
ENTRYPOINT ["/main"]