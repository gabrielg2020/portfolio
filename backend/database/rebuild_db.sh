#!/bin/sh
echo ""

RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

database_file=database/portfolio.db

# Run SQL script to rebuild schema
schema_file=database/schema.sql
echo "${CYAN}Rebuilding database schema${NC}"
echo "  - Running schema: $schema_file"
if ! output=$(sqlite3 "$database_file" < "$schema_file" 2>&1); then
  echo "    - Error in $schema_file"
  echo ""
  echo "\033[0;31m$output\033[0m"
  exit 1
fi

echo ""

# Run database patches to insert data
for dir in database/data/*/; do
  echo "${CYAN}Processing patch version: $dir${NC}"
  for sql_file in "$dir"*.sql; do
    if [ -f "$sql_file" ]; then
      echo "  - Running SQL file: $sql_file"
      if ! output=$(sqlite3 "$database_file" < "$sql_file" 2>&1); then
        echo "    - Error in $sql_file"
        echo ""
        echo "${RED}$output${NC}"
        exit 1
      fi
    fi
  done
done

echo ""
echo "${GREEN}Success!${NC}"
echo ""