# URL Shortener Service

A modern URL shortening service built with Go, Redis, and Docker.

## Features

- Shorten long URLs to manageable links
- Track click statistics for shortened URLs
- Redis-based storage for fast access
- Docker containerization for easy deployment
- RESTful API endpoints

## Prerequisites

- Docker and Docker Compose
- Go 1.21 or later (for local development)

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/Sharani12/Url-Shortener.git
cd Url-Shortener
```

2. Run with Docker Compose:
```bash
docker compose up --build
```

The service will be available at `http://localhost:8080`

## API Endpoints

### Create Short URL
```
POST /api/v1/shorten
Content-Type: application/json

{
    "url": "https://example.com/very/long/url"
}
```

### Get URL Statistics
```
GET /api/v1/stats/{shortCode}
```

### Redirect to Original URL
```
GET /{shortCode}
```

## Local Development

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

## Project Structure

```
.
├── main.go           # Application entry point
├── handlers.go       # HTTP request handlers
├── handlers_test.go  # Unit tests
├── Dockerfile        # Docker configuration
├── docker-compose.yml # Docker Compose configuration
└── go.mod           # Go module definition
```

## Environment Variables

- `REDIS_ADDR`: Redis server address (default: redis:6379)
- `REDIS_PASSWORD`: Redis password (optional)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 