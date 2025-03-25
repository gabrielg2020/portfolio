#!/bin/sh
echo ""

RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

database_directory_path=database

# Parse options
while getopts ":p:" opt; do
  case $opt in
    p) database_directory_path="$OPTARG" ;;
    \?) echo "${RED}Invalid option -$OPTARG" >&2; echo "Use -p </path/to/database>${NC}"; exit 1 ;;
  esac
done

# Shift the parsed options
shift $((OPTIND-1))

database_file=$database_directory_path/portfolio.db
schema_file=$database_directory_path/schema.sql
data_patch_directory=$database_directory_path/data

# Check if portfolio.db exisits
if [ ! -f "$database_file" ]; then
  echo "${RED}'portfolio.db' doesn't exist in $database_directory_path" >&2
  exit 1
fi

# Check if schema.sql exisits
if [ ! -f "$schema_file" ]; then
  echo "${RED}'schema.sql' doesn't exist in $database_directory_path" >&2
  exit 1
fi

# Check if data directory exists
if [ ! -d "$data_patch_directory" ]; then
  echo "${RED}'/data' doesn't exist in $database_directory_path" >&2
  exit 1
fi

# Run SQL script to rebuild schema
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
for dir in $data_patch_directory/*/; do
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