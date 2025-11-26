#!/bin/bash

# Foliogram Database Initialization Script
# This script initializes or reinitializes the database with the schema

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

SCHEMA_FILE="./assets/schema_mysql.sql"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_NAME="${DB_NAME:-foliogram_db}"
DB_USER="${DB_USER:-root}"
DB_PASSWORD="${DB_PASSWORD:-rootpassword}"

# Print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if schema file exists
if [ ! -f "$SCHEMA_FILE" ]; then
    print_error "Schema file not found: $SCHEMA_FILE"
    exit 1
fi

# Check if running in Docker or host
if docker ps | grep -q foliogram_mysql; then
    print_info "Detected Docker environment. Using Docker MySQL container..."
    USE_DOCKER=true
else
    print_info "Using host MySQL connection..."
    USE_DOCKER=false
fi

# Initialize database
init_database() {
    print_info "Initializing database: $DB_NAME"
    
    if [ "$USE_DOCKER" = true ]; then
        # Using Docker container
        print_info "Dropping existing database if exists..."
        docker exec -i foliogram_mysql mysql -u root -p"$DB_PASSWORD" -e "DROP DATABASE IF EXISTS $DB_NAME;"
        
        print_info "Creating database..."
        docker exec -i foliogram_mysql mysql -u root -p"$DB_PASSWORD" -e "CREATE DATABASE $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
        
        print_info "Loading schema..."
        docker exec -i foliogram_mysql mysql -u root -p"$DB_PASSWORD" "$DB_NAME" < "$SCHEMA_FILE"
    else
        # Using host MySQL
        print_info "Dropping existing database if exists..."
        mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" -e "DROP DATABASE IF EXISTS $DB_NAME;"
        
        print_info "Creating database..."
        mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" -e "CREATE DATABASE $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
        
        print_info "Loading schema..."
        mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < "$SCHEMA_FILE"
    fi
    
    print_success "Database initialized successfully!"
}

# Show database info
show_info() {
    print_info "Database Information:"
    echo "  Host: $DB_HOST"
    echo "  Port: $DB_PORT"
    echo "  Database: $DB_NAME"
    echo "  Schema File: $SCHEMA_FILE"
    echo ""
    
    if [ "$USE_DOCKER" = true ]; then
        print_info "Tables in database:"
        docker exec foliogram_mysql mysql -u root -p"$DB_PASSWORD" "$DB_NAME" -e "SHOW TABLES;"
    else
        print_info "Tables in database:"
        mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" -e "SHOW TABLES;"
    fi
}

# Main execution
case "${1:-init}" in
    init)
        print_warning "This will DROP and recreate the database '$DB_NAME'."
        read -p "Are you sure you want to continue? [y/N] " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            init_database
            show_info
        else
            print_info "Operation cancelled."
            exit 0
        fi
        ;;
    force)
        print_warning "Force mode: Dropping and recreating database without confirmation..."
        init_database
        show_info
        ;;
    info)
        show_info
        ;;
    *)
        echo "Usage: $0 [init|force|info]"
        echo "  init  - Initialize database (with confirmation)"
        echo "  force - Initialize database (without confirmation)"
        echo "  info  - Show database information"
        exit 1
        ;;
esac
