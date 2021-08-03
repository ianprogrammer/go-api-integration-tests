#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER productuser WITH PASSWORD 'product';
    CREATE DATABASE product;
    GRANT ALL PRIVILEGES ON DATABASE product TO productuser;
    \c product   
    create extension if not exists "uuid-ossp";
    create extension if not exists citext;
EOSQL