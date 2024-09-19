# GoGateway

GoGateway is a microservice API Gateway built using Golang and Go-chi, designed to integrate and manage API services with security, scalability, and cloud integration. It provides rate limiting, health checks, authentication, logging, and cloud integration with AWS. The gateway is containerized with Docker and orchestrated using Kubernetes, making it ideal for production environments.

## Key Features

- **API Rate Limiting**: Prevent abuse and manage traffic using a Golang rate limiter.
- **JWT Authentication**: Stateless authentication for securing API endpoints.
- **Role-Based Access Control (RBAC)**: Manage access to routes based on user roles.
- **Health Checks**: Built-in health endpoints for monitoring service health and uptime.
- **PostgreSQL Database Integration**: Used for managing persistent data.
- **AWS Cloud Integration**: Interact with cloud services like S3 and IAM via AWS SDK.
- **Logging & Monitoring**: Structured logging and Prometheus integration for monitoring.
- **Containerization & Orchestration**: Docker for containerization and Kubernetes for deployment.

## Technology Stack

- **Language**: Golang
- **Router**: Go-Chi (HTTP Router)
- **Database**: PostgreSQL
- **Cloud**: AWS (S3, IAM)
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Rate Limiting**: Golang Limiter package
- **Logging & Monitoring**: Structured logging, Prometheus for monitoring

## Prerequisites

Before you can run the project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (1.16+)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/) (optional for orchestration)
- [PostgreSQL](https://www.postgresql.org/) (for database)
- AWS credentials configured if you're using cloud integration

## Getting Started

### 1. Clone the Repository
```bash
git clone 
cd GoGateway
```

### 2. install dependencies
```bash
go mod tidy
```

### 3. Setup Environment Variable
Create a .env file in the project root with the following variable ()
```bash
DATABASE_URL=postgres://username:password@localhost:5432/gatewaydb
JWT_SECRET=your_jwt_secret_key
AWS_ACCESS_KEY=your_aws_access_key
AWS_SECRET_KEY=your_aws_secret_key

```

### 4. Database setup

### 5. Run the application


## Api Endpoints
### Health check
### Protected Route
### Rate limited Endpoint

## Docker setup

## Kubernetes

## testing

## Cloud integration (AWS)

## logging and monitoring

## Conclusion

GoGateway is a production-ready API Gateway that implements essential features for microservice architecture, such as rate limiting, authentication, health checks, logging, and cloud integration. This project is containerized using Docker and Kubernetes for easy deployment and scaling, and integrates with AWS for cloud functionality.

This project showcases real-world knowledge of developing and deploying scalable microservices, making it an excellent candidate for anyone looking to gain experience in Golang, cloud computing, and infrastructure management.