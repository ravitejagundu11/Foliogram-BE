#!/bin/bash

# Seed Templates Script
# This script populates the templates table with pre-defined portfolio templates

echo "🌱 Seeding templates into database..."

# Get database credentials from environment or use defaults
DB_HOST=${DB_HOST:-"localhost"}
DB_PORT=${DB_PORT:-"3306"}
DB_NAME=${DB_NAME:-"foliogram_db"}
DB_USER=${DB_USER:-"foliogram_user"}
DB_PASSWORD=${DB_PASSWORD:-"foliogram_pass"}

# Check if running with Docker
if docker ps | grep -q foliogram_mysql; then
    echo "📦 Detected Docker MySQL container"
    docker exec -i foliogram_mysql mysql -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < assets/seed_templates.sql
else
    echo "💻 Using local MySQL"
    mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < assets/seed_templates.sql
fi

if [ $? -eq 0 ]; then
    echo "✅ Templates seeded successfully!"
else
    echo "❌ Failed to seed templates"
    exit 1
fi
