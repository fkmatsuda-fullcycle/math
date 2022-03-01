FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o math


FROM alpine:3.15

COPY --from=builder /app /app

CMD ["/app/math"]