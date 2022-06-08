FROM golang:ubuntu as builder

RUN apt update && apt upgrade && \
    apt add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment

# Run container
FROM ubuntu:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]
