# Instagram Microservices

This is a sample project to demonstrate how to build a project using microservices architecture. In this project i use Golang as the main programming language and PostgreSQL as the main database. I also use Docker and Docker Compose to run the services. Redis-stream is used as the message broker. Nginx is used as the API Gateway and Reverse Proxy. SMTP is used to send email for user verification.

## Architecture

![Architecture](

## Technologies

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Nginx](https://www.nginx.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)

## Services

- Nginx: API Gateway and Reverse Proxy
- auth_service: handles user authentication and authorization
- post_service: handles posts
- account_service: handles creating account for users
- comment_service: handles comments for posts
- like_service: handles likes for posts & comments
- user-profile_service: handles user profile
- followers_service: handles followers for users

## How to run

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Golang](https://golang.org/)
- [go-redis](https://github.com/redis/go-redis/v9)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Nginx](https://www.nginx.com/)
- [SMTP](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)

### Steps

1. Clone the repository

```bash
git clone 
```

2. Run the following command to start the services

```bash
docker-compose up -d
```

3. Run the following command to stop the services

```bash
docker-compose down
```