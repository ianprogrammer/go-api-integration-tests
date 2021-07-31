#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER product WITH PASSWORD 'product';
    CREATE DATABASE product;
    GRANT ALL PRIVILEGES ON DATABASE product TO product;
EOSQL