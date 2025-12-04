#!/bin/bash

# Foliogram Development Environment Manager
set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
print_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Start all services
start_all() {
    print_info "Starting services..."
    docker-compose up -d
    print_info "Backend: http://localhost:8080 | phpMyAdmin: http://localhost:8081"
}

# Start backend only
start() {
    print_info "Starting backend..."
    docker-compose up  backend
    print_info "Backend: http://localhost:8080"
}

# Stop all services
stop() {
    print_info "Stopping services..."
    docker-compose stop
}

# Restart all services
restart() {
    print_info "Restarting services..."
    docker-compose restart
}

# Stop and remove containers
down() {
    print_info "Removing containers..."
    docker-compose down
}

# Build backend with latest code
build() {
    print_info "Building backend with latest code..."
    docker-compose build --no-cache backend
    print_info "Build complete!"
}

# View logs
logs() {
    docker-compose logs -f ${2:-}
}

# Show status
status() {
    docker-compose ps
}

# Show help
help() {
    echo "Foliogram Dev Environment"
    echo ""
    echo "Usage: ./dev-env.sh <command>"
    echo ""
    echo "Commands:"
    echo "  start      - Start backend only"
    echo "  start-all  - Start all services"
    echo "  stop       - Stop all services"
    echo "  restart    - Restart all services"
    echo "  down       - Stop and remove containers"
    echo "  build      - Build backend with latest code"
    echo "  logs       - View logs (add service name for specific logs)"
    echo "  status   - Show service status"
    echo "  help     - Show this help"
}

# Main
case "$1" in
    start) start ;;
    start-all) start_all ;;
    stop) stop ;;
    restart) restart ;;
    down) down ;;
    build) build ;;
    logs) logs "$@" ;;
    status) status ;;
    help|--help|-h|"") help ;;
    *)
        print_error "Unknown command: $1"
        help
        exit 1
        ;;
esac
