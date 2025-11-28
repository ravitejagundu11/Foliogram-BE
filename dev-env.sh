#!/bin/bash

# Foliogram Development Environment Manager
# This script manages the Foliogram backend and MySQL database

set -e

# Check if output is to a terminal and enable colors accordingly
if [ -t 1 ]; then
    # Colors for output (terminal detected)
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m' # No Color
else
    # No colors (output redirected/piped)
    RED=''
    GREEN=''
    YELLOW=''
    BLUE=''
    NC=''
fi

# Configuration
PROJECT_NAME="Foliogram"
COMPOSE_FILE="docker-compose.yml"
ENV_FILE=".env"

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

# Check if Docker is running
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        print_error "Docker is not running. Please start Docker and try again."
        exit 1
    fi
}

# Check if .env file exists, create from example if not
check_env_file() {
    if [ ! -f "$ENV_FILE" ]; then
        if [ -f ".env.example" ]; then
            print_warning ".env file not found. Creating from .env.example..."
            cp .env.example .env
            print_info "Please update .env with your configuration"
        else
            print_warning ".env file not found. Using default values."
        fi
    fi
}

# Start all services
start() {
    print_info "Starting $PROJECT_NAME development environment..."
    check_docker
    check_env_file
    
    docker-compose up -d
    
    print_success "Services started successfully!"
    print_info "Backend API: http://localhost:8080"
    print_info "phpMyAdmin: http://localhost:8081"
    print_info ""
    print_info "Run './dev-env.sh logs' to view logs"
    print_info "Run './dev-env.sh status' to check service status"
}

# Start specific service
start_service() {
    if [ -z "$2" ]; then
        print_error "Please specify a service: backend, mysql, or phpmyadmin"
        exit 1
    fi
    
    check_docker
    check_env_file
    
    print_info "Starting $2..."
    docker-compose up  "$2"
    print_success "$2 started successfully!"
}

# Stop all services
stop() {
    print_info "Stopping $PROJECT_NAME development environment..."
    docker-compose stop
    print_success "Services stopped successfully!"
}

# Stop specific service
stop_service() {
    if [ -z "$2" ]; then
        print_error "Please specify a service: backend, mysql, or phpmyadmin"
        exit 1
    fi
    
    print_info "Stopping $2..."
    docker-compose stop "$2"
    print_success "$2 stopped successfully!"
}

# Restart all services
restart() {
    print_info "Restarting $PROJECT_NAME development environment..."
    docker-compose restart
    print_success "Services restarted successfully!"
}

# Restart specific service
restart_service() {
    if [ -z "$2" ]; then
        print_error "Please specify a service: backend, mysql, or phpmyadmin"
        exit 1
    fi
    
    print_info "Restarting $2..."
    docker-compose restart "$2"
    print_success "$2 restarted successfully!"
}

# Stop and remove all containers, networks
down() {
    print_warning "Stopping and removing $PROJECT_NAME containers and networks..."
    docker-compose down
    print_success "Services removed successfully!"
}

# Stop and remove all containers, networks, and volumes (DELETES DATA!)
destroy() {
    read -p "$(echo -e ${RED}WARNING: This will delete all data including database. Continue? [y/N]${NC} )" -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_warning "Destroying $PROJECT_NAME environment and data..."
        docker-compose down -v
        print_success "Environment destroyed successfully!"
    else
        print_info "Operation cancelled."
    fi
}

# View logs
logs() {
    if [ -z "$2" ]; then
        docker-compose logs -f
    else
        docker-compose logs -f "$2"
    fi
}

# Show status of services
status() {
    print_info "$PROJECT_NAME service status:"
    docker-compose ps
}

# Rebuild and restart services
rebuild() {
    print_info "Rebuilding $PROJECT_NAME services..."
    docker-compose build --no-cache
    docker-compose up -d
    print_success "Services rebuilt and restarted successfully!"
}

# Execute command in backend container
exec_backend() {
    if [ -z "$2" ]; then
        docker-compose exec backend sh
    else
        shift
        docker-compose exec backend "$@"
    fi
}

# Execute command in MySQL container
exec_mysql() {
    docker-compose exec mysql mysql -u foliogram_user -pfoliogram_pass foliogram_db
}

# Backup database
backup() {
    BACKUP_DIR="./backups"
    mkdir -p "$BACKUP_DIR"
    BACKUP_FILE="$BACKUP_DIR/foliogram_backup_$(date +%Y%m%d_%H%M%S).sql"
    
    print_info "Creating database backup..."
    docker-compose exec -T mysql mysqldump -u root -prootpassword foliogram_db > "$BACKUP_FILE"
    print_success "Database backed up to: $BACKUP_FILE"
}

# Restore database
restore() {
    if [ -z "$2" ]; then
        print_error "Please provide backup file path"
        print_info "Usage: ./dev-env.sh restore <backup-file.sql>"
        exit 1
    fi
    
    if [ ! -f "$2" ]; then
        print_error "Backup file not found: $2"
        exit 1
    fi
    
    print_warning "Restoring database from: $2"
    docker-compose exec -T mysql mysql -u root -prootpassword foliogram_db < "$2"
    print_success "Database restored successfully!"
}

# Show database info
db_info() {
    print_info "Database Connection Information:"
    echo "Host: localhost"
    echo "Port: 3306"
    echo "Database: foliogram_db"
    echo "Username: foliogram_user"
    echo "Password: foliogram_pass"
    echo ""
    print_info "Root Password: rootpassword"
    echo ""
    print_info "Connect with MySQL client:"
    echo "mysql -h localhost -P 3306 -u foliogram_user -pfoliogram_pass foliogram_db"
}

# Initialize or reinitialize database with schema
init_db() {
    if [ "$2" = "force" ]; then
        print_warning "Force initializing database..."
        ./init-db.sh force
    else
        print_info "Initializing database..."
        ./init-db.sh init
    fi
}

# Show database tables and info
db_tables() {
    print_info "Database tables:"
    docker-compose exec -T mysql mysql -u root -prootpassword foliogram_db -e "SHOW TABLES;"
    echo ""
    print_info "Table row counts:"
    docker-compose exec -T mysql mysql -u root -prootpassword foliogram_db -e "
        SELECT 'users' as table_name, COUNT(*) as row_count FROM users
        UNION SELECT 'portfolios', COUNT(*) FROM portfolios
        UNION SELECT 'projects', COUNT(*) FROM projects
        UNION SELECT 'skills', COUNT(*) FROM skills
        UNION SELECT 'testimonials', COUNT(*) FROM testimonials
        UNION SELECT 'appointments', COUNT(*) FROM appointments
        UNION SELECT 'blogs', COUNT(*) FROM blogs
        UNION SELECT 'blog_subscribers', COUNT(*) FROM blog_subscribers
        UNION SELECT 'blog_likes', COUNT(*) FROM blog_likes
        UNION SELECT 'blog_comments', COUNT(*) FROM blog_comments
        UNION SELECT 'analytics_views', COUNT(*) FROM analytics_views
        UNION SELECT 'notifications', COUNT(*) FROM notifications;
    "
}

# Run database migrations (placeholder for future use)
migrate() {
    print_info "Running database migrations..."
    print_warning "Migration functionality not yet implemented"
    # Add migration logic here when available
}

# Clean up Docker resources
clean() {
    print_warning "Cleaning up Docker resources..."
    docker system prune -f
    print_success "Docker cleanup completed!"
}

# Show help
help() {
    printf "${GREEN}Foliogram Development Environment Manager${NC}\n\n"
    printf "${YELLOW}Usage:${NC}\n"
    printf "    ./dev-env.sh <command>\n\n"
    
    printf "${YELLOW}Commands:${NC}\n"
    printf "    ${GREEN}start${NC}              Start all services (backend + MySQL + Redis + phpMyAdmin)\n"
    printf "    ${GREEN}start-backend${NC}      Start backend service only\n"
    printf "    ${GREEN}start-mysql${NC}        Start MySQL service only\n"
    printf "    ${GREEN}start-redis${NC}        Start Redis service only\n"
    printf "    ${GREEN}start-phpmyadmin${NC}   Start phpMyAdmin service only\n"
    printf "\n"
    printf "    ${GREEN}stop${NC}               Stop all services\n"
    printf "    ${GREEN}stop-backend${NC}       Stop backend service only\n"
    printf "    ${GREEN}stop-mysql${NC}         Stop MySQL service only\n"
    printf "    ${GREEN}stop-redis${NC}         Stop Redis service only\n"
    printf "    ${GREEN}stop-phpmyadmin${NC}    Stop phpMyAdmin service only\n"
    printf "\n"
    printf "    ${GREEN}restart${NC}            Restart all services\n"
    printf "    ${GREEN}restart-backend${NC}    Restart backend service only\n"
    printf "    ${GREEN}restart-mysql${NC}      Restart MySQL service only\n"
    printf "    ${GREEN}restart-redis${NC}      Restart Redis service only\n"
    printf "    ${GREEN}restart-phpmyadmin${NC} Restart phpMyAdmin service only\n"
    printf "\n"
    printf "    ${GREEN}down${NC}               Stop and remove containers and networks\n"
    printf "    ${GREEN}destroy${NC}            Stop and remove containers, networks, and volumes (deletes data!)\n"
    printf "    ${GREEN}logs${NC} [service]     View logs (all services or specific service)\n"
    printf "    ${GREEN}status${NC}             Show status of all services\n"
    printf "    ${GREEN}rebuild${NC}            Rebuild and restart services\n"
    printf "\n"
    printf "    ${BLUE}backend${NC} [cmd]   Execute command in backend container (default: shell)\n"
    printf "    ${BLUE}mysql${NC}           Open MySQL client\n"
    printf "    ${BLUE}init-db${NC} [force] Initialize/reinitialize database with schema\n"
    printf "    ${BLUE}db-tables${NC}       Show all database tables and row counts\n"
    printf "    ${BLUE}backup${NC}          Backup database to ./backups/\n"
    printf "    ${BLUE}restore${NC} <file>  Restore database from backup file\n"
    printf "    ${BLUE}db-info${NC}         Show database connection information\n"
    printf "    ${BLUE}migrate${NC}         Run database migrations\n"
    printf "\n"
    printf "    ${YELLOW}clean${NC}           Clean up Docker resources\n"
    printf "    ${YELLOW}help${NC}            Show this help message\n"
    printf "\n"
    
    printf "${YELLOW}Examples:${NC}\n"
    printf "    ./dev-env.sh start                      # Start all services\n"
    printf "    ./dev-env.sh start-backend              # Start only backend\n"
    printf "    ./dev-env.sh restart-mysql              # Restart only MySQL\n"
    printf "    ./dev-env.sh logs backend               # View backend logs\n"
    printf "    ./dev-env.sh backend ls -la             # Run command in backend\n"
    printf "    ./dev-env.sh backup                     # Backup database\n"
    printf "    ./dev-env.sh init-db                    # Initialize database\n"
    printf "    ./dev-env.sh restore ./backups/backup.sql\n"
    printf "\n"
    
    printf "${YELLOW}Service URLs:${NC}\n"
    printf "    Backend API:  http://localhost:8000\n"
    printf "    phpMyAdmin:   http://localhost:8081\n"
    printf "\n"
}

# Main script logic
case "$1" in
    start)
        start
        ;;
    start-backend)
        start_service "$@" backend
        ;;
    start-mysql)
        start_service "$@" mysql
        ;;
    start-redis)
        start_service "$@" redis
        ;;
    start-phpmyadmin)
        start_service "$@" phpmyadmin
        ;;
    stop)
        stop
        ;;
    stop-backend)
        stop_service "$@" backend
        ;;
    stop-mysql)
        stop_service "$@" mysql
        ;;
    stop-redis)
        stop_service "$@" redis
        ;;
    stop-phpmyadmin)
        stop_service "$@" phpmyadmin
        ;;
    restart)
        restart
        ;;
    restart-backend)
        restart_service "$@" backend
        ;;
    restart-mysql)
        restart_service "$@" mysql
        ;;
    restart-redis)
        restart_service "$@" redis
        ;;
    restart-phpmyadmin)
        restart_service "$@" phpmyadmin
        ;;
    down)
        down
        ;;
    destroy)
        destroy
        ;;
    logs)
        logs "$@"
        ;;
    status)
        status
        ;;
    rebuild)
        rebuild
        ;;
    backend)
        exec_backend "$@"
        ;;
    mysql)
        exec_mysql
        ;;
    init-db)
        init_db "$@"
        ;;
    db-tables)
        db_tables
        ;;
    backup)
        backup
        ;;
    restore)
        restore "$@"
        ;;
    db-info)
        db_info
        ;;
    migrate)
        migrate
        ;;
    clean)
        clean
        ;;
    help|--help|-h)
        help
        ;;
    *)
        print_error "Unknown command: $1"
        echo ""
        help
        exit 1
        ;;
esac
