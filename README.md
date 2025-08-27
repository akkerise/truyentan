# Truyen Reader

[![CI](https://github.com/OWNER/truyentan/actions/workflows/ci.yml/badge.svg)](https://github.com/OWNER/truyentan/actions/workflows/ci.yml)

This repository contains the initial skeleton for a Wails v2 project with separate backend and frontend directories.

## Structure

- `backend/` - Go backend.
  - `cmd/api/main.go` - entry point.
  - `cmd/seed/main.go` - database seed command.
- `web/` - React frontend.
  - `src/main.tsx`
  - `src/pages/`, `src/components/`, `src/context/`

## Getting Started

### Prerequisites
- Go 1.24+
- Node.js 20+
- Docker & docker compose
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Setup
1. Copy `.env.example` to `.env` and adjust values.
2. Start supporting services:
   ```bash
   docker compose up -d
   ```
3. Generate Swagger docs:
   ```bash
   make swagger
   ```
4. Seed the database:
   ```bash
   make seed
   ```

### Development
- Run the app in development mode:
  ```bash
  make dev
  ```
- Start only the frontend:
  ```bash
  cd web && npm run dev
  ```
- Start only the backend:
  ```bash
  go run backend/cmd/api/main.go
  ```

### Build
Build the desktop application:
```bash
make build
```

### Linting
Run Go and frontend linters:
```bash
make lint
```

## Configuration
See `wails.json` for project configuration.

