
FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git


COPY go.mod ./

RUN go mod download || true

COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/api


FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Bangkok
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 3000
ENTRYPOINT ["./server"]
