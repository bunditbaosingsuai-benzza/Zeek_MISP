FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
# เอา .env มาใช้ด้วย
COPY .env .  
RUN apt-get update && apt-get install -y tzdata ca-certificates && rm -rf /var/lib/apt/lists/*
EXPOSE 5050
ENTRYPOINT ["./app"]
