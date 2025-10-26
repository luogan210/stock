#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"

cleanup() {
  echo "\nShutting down..."
  [[ -n "${SERVER_PID:-}" ]] && kill ${SERVER_PID} >/dev/null 2>&1 || true
  [[ -n "${WEB_PID:-}" ]] && kill ${WEB_PID} >/dev/null 2>&1 || true
}
trap cleanup EXIT INT TERM

echo "Starting backend (Go) on :8080..."
( cd "$ROOT_DIR/server" && go run main.go ) &
SERVER_PID=$!

sleep 1

echo "Starting frontend (Vite) on :3000..."
( cd "$ROOT_DIR/web" && npm run dev ) &
WEB_PID=$!

echo "Backend PID: $SERVER_PID"
echo "Frontend PID: $WEB_PID"
echo "Press Ctrl+C to stop."

wait


