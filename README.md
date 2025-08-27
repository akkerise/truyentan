# Truyen Reader

This repository contains the initial skeleton for a Wails v2 project with separate backend and frontend directories.

## Structure

- `backend/` - Go backend stub.
  - `cmd/api/main.go` - entry point.
- `web/` - React frontend.
  - `src/main.tsx`
  - `src/pages/`, `src/components/`, `src/context/`

## Development

Use the Wails CLI to run the application in development mode:

```bash
wails dev
```

This will start the frontend and the Go backend. The backend Gin server listens on `127.0.0.1:API_PORT` where `API_PORT` is read from the environment (default `8080`).

## Building

To produce a desktop binary run:

```bash
wails build
```

The build targets Windows and Linux as configured in `wails.json`.

## Configuration

See `wails.json` for project configuration.
