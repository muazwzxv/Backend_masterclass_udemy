#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE DATABASE go_masterclass;

  CREATE USER bank_user WITH ENCRYPTED PASSWORD 'bank_pass';

  GRANT CONNECT ON DATABASE go_masterclass to bank_user;
EOSQL

