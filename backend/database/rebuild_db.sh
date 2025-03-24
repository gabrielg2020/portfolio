#!/bin/bash

database_file=database/portfolio.db

# Run SQL script to rebuild schema
schema_file=database/schema.sql
echo "Rebuilding database schema, using $schema_file"
sqlite3 "$database_file" < $schema_file

echo ""

# Run database patches to insert data
for dir in database/data/*/; do
  echo "Processing patch version: $dir"
  for sql_file in "$dir"*.sql; do
    if [ -f "$sql_file" ]; then
      echo "  - Running SQL file: $sql_file"
      sqlite3 "$database_file" < "$sql_file"
    fi
  done
done

echo " "
echo "Success!"