# Builder stage
FROM golang:1.19.5-alpine3.17 as BUILDER
# install gcc and musl-dev for compiling go code
RUN apk add gcc musl-dev

WORKDIR /tiddi

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY src ./src

RUN go build ./src/main.go

# Final Stage
FROM alpine:3.17

WORKDIR /tiddi

# Copy sample frontend
COPY --from=BUILDER /tiddi/src/frontend ./src/frontend

COPY --from=BUILDER /tiddi/main .

EXPOSE 5656

CMD ["./main"]