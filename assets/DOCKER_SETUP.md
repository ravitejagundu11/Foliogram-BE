# Foliogram Backend - Development Environment

## Quick Start

### Using the dev-env script (Recommended)

```bash
# Make the script executable (first time only)
chmod +x dev-env.sh

# Start all services
./dev-env.sh start

# View logs
./dev-env.sh logs

# Stop services
./dev-env.sh stop
```

### Manual Docker Compose

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## Dev Environment Script Commands

### Service Management
```bash
./dev-env.sh start          # Start all services
./dev-env.sh stop           # Stop all services
./dev-env.sh restart        # Restart all services
./dev-env.sh down           # Stop and remove containers
./dev-env.sh destroy        # Remove everything including data
./dev-env.sh status         # Check service status
./dev-env.sh rebuild        # Rebuild and restart services
```

### Logs
```bash
./dev-env.sh logs           # View all logs
./dev-env.sh logs backend   # View backend logs only
./dev-env.sh logs mysql     # View MySQL logs only
```

### Backend Container
```bash
./dev-env.sh backend        # Open shell in backend container
./dev-env.sh backend ls -la # Run command in backend container
```

### Database Management
```bash
./dev-env.sh mysql          # Open MySQL client
./dev-env.sh backup         # Backup database
./dev-env.sh restore <file> # Restore from backup
./dev-env.sh db-info        # Show DB connection info
```

### Utilities
```bash
./dev-env.sh clean          # Clean up Docker resources
./dev-env.sh help           # Show help message
```

## Service URLs

- **Backend API**: http://localhost:8080
- **phpMyAdmin**: http://localhost:8081
- **MySQL**: localhost:3306

## Environment Configuration

Copy `.env.example` to `.env` and customize:

```bash
cp .env.example .env
```

## Project Structure

```
Foliogram-BE/
├── conf/               # Configuration files
├── controllers/        # API controllers
├── docs/              # Documentation & SQL schema
├── models/            # Data models
├── routers/           # Route definitions
├── static/            # Static assets
├── tests/             # Test files
├── uploads/           # File uploads (volume mounted)
├── views/             # Templates
├── docker-compose.yml # Docker services configuration
├── Dockerfile         # Backend Docker image
├── dev-env.sh         # Development environment manager
├── .env.example       # Environment variables template
└── main.go            # Application entry point
```

## Development Workflow

### 1. Initial Setup
```bash
# Clone repository
git clone <repository-url>
cd Foliogram-BE

# Copy environment file
cp .env.example .env

# Start services
./dev-env.sh start
```

### 2. Making Changes
```bash
# Make code changes
# ...

# Rebuild backend only
docker-compose up -d --build backend

# Or rebuild everything
./dev-env.sh rebuild
```

### 3. Debugging
```bash
# View backend logs
./dev-env.sh logs backend

# Access backend container
./dev-env.sh backend

# Check database
./dev-env.sh mysql
```

### 4. Database Operations
```bash
# Backup before major changes
./dev-env.sh backup

# Restore if needed
./dev-env.sh restore ./backups/backup_file.sql
```

## Docker Services

### Backend (Go/Beego)
- **Container**: foliogram_backend
- **Port**: 8080
- **Image**: Built from Dockerfile
- **Volumes**: 
  - `./uploads:/app/uploads` (file uploads)
  - `./conf:/app/conf` (configuration)

### MySQL
- **Container**: foliogram_mysql
- **Port**: 3306
- **Image**: mysql:8.0
- **Volume**: `mysql_data` (persistent storage)
- **Auto-init**: Schema loaded from `docs/schema_mysql.sql`

### phpMyAdmin
- **Container**: foliogram_phpmyadmin
- **Port**: 8081
- **Image**: phpmyadmin:latest

## Troubleshooting

### Services won't start
```bash
# Check Docker is running
docker info

# Check logs
./dev-env.sh logs

# Reset everything
./dev-env.sh destroy
./dev-env.sh start
```

### Port already in use
Edit `.env` file to change ports:
```env
APP_PORT=8090
DB_PORT=3307
```

### Backend can't connect to MySQL
```bash
# Ensure MySQL is healthy
./dev-env.sh status

# Check backend logs
./dev-env.sh logs backend

# Restart services
./dev-env.sh restart
```

### Database connection from host
```bash
# Connection string
mysql -h localhost -P 3306 -u foliogram_user -pfoliogram_pass foliogram_db

# Or use the script
./dev-env.sh mysql
```

## Production Deployment

Before deploying to production:

1. ✅ Update all passwords in `.env`
2. ✅ Set `APP_ENV=production`
3. ✅ Configure proper JWT secret
4. ✅ Set up SSL/TLS
5. ✅ Remove phpMyAdmin or restrict access
6. ✅ Configure backup strategy
7. ✅ Set up monitoring and logging
8. ✅ Use Docker secrets for sensitive data
9. ✅ Configure reverse proxy (nginx/traefik)
10. ✅ Set up CI/CD pipeline

## Additional Resources

- [Beego Documentation](https://beego.me/docs)
- [Docker Documentation](https://docs.docker.com)
- [MySQL Documentation](https://dev.mysql.com/doc)
- [Go Documentation](https://golang.org/doc)
