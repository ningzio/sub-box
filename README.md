# Sub-Box

A JSON Schema based configuration editor with validation and generation capabilities.

## Project Structure

```
/sub-box
├── frontend/               # React frontend
│   ├── src/               # Source code
│   ├── public/            # Static files
│   └── package.json       # Dependencies and scripts
└── backend/               # Go backend
    ├── cmd/               # Command-line applications
    │   └── server/        # Main server application
    ├── internal/          # Private application code
    │   ├── api/          # HTTP handlers
    │   ├── service/      # Business logic
    │   └── config/       # Configuration
    └── pkg/              # Public libraries
```

## Development

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

## API Endpoints

- `GET /health` - Health check
- `POST /api/config/validate` - Validate configuration
- `POST /api/config/generate` - Generate configuration
