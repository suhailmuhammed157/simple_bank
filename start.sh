#!/bin/sh
set -e

echo "run database migration"
/app/migrate -path /app/migrations -database "$DATA_SOURCE" -verbose up

echo "start the app"
exec "$@"