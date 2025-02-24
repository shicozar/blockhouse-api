# blockhouse-api
# Simple REST API in Golang

## Overview
We built a simple REST API in **Golang** using the **Gin** framework. The API allows users to submit trade order details and retrieve all submitted orders. The data is stored in an **SQLite** database.

## Features Implemented
- **POST /orders**: Accepts trade order details such as symbol, price, quantity, and order type.
- **GET /orders**: Returns a list of all submitted orders.
- **Database**: Uses **SQLite** for lightweight storage.

## Technologies Used
- **Golang**: Chosen for its performance, efficiency, and strong concurrency handling.
- **Gin**: A lightweight and fast HTTP web framework for building REST APIs.
- **SQLite**: Chosen for simplicity and ease of setup for a small-scale application.

---

## DevOps & Deployment (AWS EC2 + Docker)

### Containerization
To package our application efficiently, we used **Docker**. This ensures that our application runs consistently across different environments.

#### Dockerfile:
```dockerfile
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite-dev
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod tidy
COPY . . 
RUN go build -o main .

FROM alpine:latest
RUN apk add --no-cache sqlite
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```
## Github Actions

- Automates testing and deployment.
- Ensures every change is validated before merging.
- Provides seamless integration with Docker and AWS EC2.


#### CI/CD 
```ci-cd.yml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.24
      - name: Run tests
        run: go test ./...

  build-and-deploy:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
      - name: Login to Docker Hub
        run: echo "${{ secrets.DOCKER_PASSWORD }}" | docker login -u "${{ secrets.DOCKER_USERNAME }}" --password-stdin
      - name: Build and push Docker image
        run: |
          docker build -t ${{ secrets.DOCKER_USERNAME }}/blockhouse-api:latest .
          docker push ${{ secrets.DOCKER_USERNAME }}/blockhouse-api:latest
      - name: Deploy to EC2
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.EC2_HOST }}
          username: ubuntu
          key: ${{ secrets.EC2_SSH_KEY }}
          script: |
            docker pull ${{ secrets.DOCKER_USERNAME }}/blockhouse-api:latest
            docker stop blockhouse-api || true
            docker rm blockhouse-api || true
            docker run -d -p 8080:8080 --name blockhouse-api ${{ secrets.DOCKER_USERNAME }}/blockhouse-api:latest
```

## Key Takeaways
- ✅ REST API in Golang (Gin) → For high-performance backend development.
- ✅ SQLite → A simple, embedded database for lightweight storage.
- ✅ Docker → Ensures consistent deployment across environments.
- ✅ AWS EC2 → Provides a cloud-hosted server for our API.
- ✅ GitHub Actions → Automates testing and deployment to EC2.
- ✅ Swagger → Documents API endpoints for easier development & integration.
