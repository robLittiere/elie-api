#!/bin/sh

if [ ! -f "/app/.initialized" ]; then
  go mod download
  touch /app/.initialized
fi

go run /app/main.go

