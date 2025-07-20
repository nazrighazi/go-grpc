#!/bin/bash
set -e

# Read database configuration from config.json
CONFIG_FILE="./products/config.json"

DB_HOST=$(jq -r '.db.host' "$CONFIG_FILE")
DB_PORT=$(jq -r '.db.port' "$CONFIG_FILE")
DB_USER=$(jq -r '.db.user' "$CONFIG_FILE")
DB_PASSWORD=$(jq -r '.db.password' "$CONFIG_FILE")
DB_NAME=$(jq -r '.db.dbname' "$CONFIG_FILE")
DB_SSLMODE=$(jq -r '.db.sslmode' "$CONFIG_FILE")

# Construct the database URL
DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"

echo "Waiting for PostgreSQL to be ready..."
until pg_isready -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}"; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

echo "PostgreSQL is up - executing migrations"

# Run database migrations
echo "Running database migrations..."
migrate -path ./products/migrations -database "${DATABASE_URL}" up

if [ $? -eq 0 ]; then
    echo "Migrations completed successfully"
else
    echo "Migration failed"
    exit 1
fi

echo "Starting products service..."
exec "$@"