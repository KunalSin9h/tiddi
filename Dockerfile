# Builder stage
FROM golang:1.21.0-alpine3.17 as BUILDER
# install gcc and musl-dev for compiling go code
RUN apk add gcc musl-dev

WORKDIR /tiddi
aman
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build ./cmd/main.go

# Final Stage
FROM alpine:3.17

WORKDIR /tiddi

# Copy sample frontend
COPY --from=BUILDER /tiddi/cmd/frontend ./cmd/frontend

COPY --from=BUILDER /tiddi/main .

EXPOSE 5656

ENTRYPOINT ["./main"]
