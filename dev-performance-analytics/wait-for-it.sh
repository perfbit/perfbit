#!/usr/bin/env bash

# Wait for the PostgreSQL server to be available
while ! pg_isready -h "$1" -p 5432; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Postgres is up - executing command"
shift
exec "$@"
