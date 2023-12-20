FROM golang:1.21.5-alpine AS builder
RUN mkdir /build
ADD go.mod main.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./main"]