FROM golang:1.18-alpine as builder
WORKDIR /app/
COPY . .
RUN go build -o /app/kumparan-assesment cmd/server/main.go

FROM alpine:3.6
WORKDIR /app/
COPY --from=builder ["/app/kumparan-assesment", "/app/.env", "/app/configs/","./"]
CMD /app/kumparan-assesment