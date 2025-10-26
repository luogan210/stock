@echo off
setlocal enableextensions enabledelayedexpansion

set ROOT_DIR=%~dp0

echo Starting backend (Go) on :8080...
start "server" cmd /c "cd /d %ROOT_DIR%server && go run main.go"

echo Starting frontend (Vite) on :3000...
start "web" cmd /c "cd /d %ROOT_DIR%web && npm run dev"

echo Both services started. Close windows to stop.


