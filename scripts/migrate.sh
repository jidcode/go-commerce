#!/bin/bash

set -e

source .env

MIGRATIONS_DIR=./internal/db/migrations

goose -dir $MIGRATIONS_DIR postgres "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up