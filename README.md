# URL Shortener Service

A RESTful URL shortener service built with Go, using Redis for storage and including Prometheus metrics.

## Features

- Shorten long URLs to short, unique codes
- Redirect short URLs to original URLs
- Rate limiting to prevent abuse
- Prometheus metrics for monitoring
- Docker support for easy deployment

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- Redis (included in Docker Compose setup)

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd url-shortener
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application using Docker Compose:
```bash
docker-compose up --build
```

The service will be available at `http://localhost:8080`

## API Endpoints

### Shorten URL
```http
POST /shorten
Content-Type: application/json

{
    "url": "https://www.example.com"
}
```

Response:
```json
{
    "short_url": "http://localhost:8080/abc123"
}
```

### Redirect
```http
GET /:shortCode
```

Redirects to the original URL.

### Metrics
```http
GET /metrics
```

Prometheus metrics endpoint.

## Testing

Run the tests:
```bash
go test -v
```

## Environment Variables

- `PORT`: Server port (default: 8080)
- `REDIS_ADDR`: Redis address (default: localhost:6379)
- `REDIS_PASSWORD`: Redis password (default: empty)

## License

MIT 