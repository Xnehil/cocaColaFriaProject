FROM golang:1.21.5-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum *.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY static/ /app/static
COPY templates/ /app/templates
copy scripts/ /app/scripts
copy tests/ /app/tests
copy configs/ /app/configs
copy docs/ /app/docs
WORKDIR /app
CMD ["./main"]