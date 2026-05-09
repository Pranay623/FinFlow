# FinFlow — A Mutual Fund / SIP Order Processing System

## Architecture
```
[Client / cURL]
      |
   [API Gateway Service - Go/Gin]
      |
   [Kafka Topics]
  /           \
[Order Service]  [Notification Service]
      |
[Portfolio Service - gRPC]
      |
[PostgreSQL + Redis Cache]
```

## Tech Stack
- **Language**: Go (Gin/Echo)
- **Message Queue**: Apache Kafka
- **Cache**: Redis
- **Database**: PostgreSQL
- **Inter-service**: gRPC
- **Containers**: Docker + Kubernetes
- **Observability**: Prometheus + Grafana

## Setup Instructions
1. Run `docker-compose up -d` to start up infrastructure (Kafka, Postgres, Redis, Prometheus, Grafana).
2. Go to each service directory (`api-gateway`, `order-service`, `portfolio-service`, `notification-service`) and run `go run main.go`.
