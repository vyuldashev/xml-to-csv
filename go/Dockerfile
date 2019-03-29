FROM golang:1.12 AS builder

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main

FROM scratch

COPY --from=builder /app/main /app/
WORKDIR /app

ENTRYPOINT ["./main"]
