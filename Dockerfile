FROM golang:1.19.5-alpine3.17

RUN apk add gcc musl-dev

WORKDIR /tiddi

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY src ./src

RUN go build -o main ./src/main.go

EXPOSE 5656

CMD ["./main"]

# We can do better by using a multi-stage build
# cause we don't need the go compiler in the final image
# and we can also use a smaller base image