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
1. Run `docker compose up --build` from the repo root.
2. This starts the infrastructure, all Go backend services, and the frontend in one command.
3. Open the frontend at http://localhost:5173 and the API gateway at http://localhost:8080.
