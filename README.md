# Kubernetes deployment for a Go REST APIs</h2>

## <a name="tech-stack">⚙️ Tech Stack</a>

- Go
- PostgreSQL
- Gin
- Docker
- Kubernetes

## <a name="quick-start">🤸 Quick Start</a>

Follow these steps to set up the project locally on your machine.

**Prerequisites**

Make sure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [Gin](https://gin-gonic.com/docs/quickstart/)
- [Docker](https://www.docker.com/)

**Cloning the Repository**

```bash
git clone https://github.com/Ademayowa/kubernetes-job-board.git
```

## Setup

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Create .env file

```bash
echo "POSTGRES_PASSWORD=mysecurepassword123" > .env
```

### 3. Running Docker compose to start all services

```bash
docker-compose up -d
```

### 3. Test the endpoint

```bash
curl -X POST http://localhost:8080/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Software Engineer",
    "description": "Develop software applications"
  }'
```

Open [http://localhost:8080/jobs](http://localhost:8080/jobs) in your browser to view all jobs.

### Create separate .env.local for Go backend development with Docker compose

```bash
cat > .env.local << EOF
DATABASE_URL=postgres://jobboard:mysecurepassword123@localhost:5432/jobboard_dev?sslmode=disable
GIN_MODE=debug
EOF
```

### Then Load it

```bash
export $(cat .env.local | xargs)
go run cmd/main.go
```

### Fixes Go backend in Local dev

```bash
# Stop everything and remove volumes
docker-compose down -v

# Rebuild with NO cache to get your current simple code
docker-compose build --no-cache app

docker-compose up -d
```
