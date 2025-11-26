# Foliogram Database Setup Guide

## MySQL Docker Setup

This guide will help you set up the MySQL database for Foliogram using Docker.

## Prerequisites

- Docker installed on your machine
- Docker Compose installed

## Quick Start

### 1. Start the Database

```bash
# Start MySQL and phpMyAdmin containers
docker-compose up -d

# Check if containers are running
docker-compose ps
```

### 2. Verify Connection

The database will be automatically initialized with the schema. You can verify:

**Using phpMyAdmin:**
- Open your browser and navigate to: `http://localhost:8080`
- Login with:
  - Username: `foliogram_user`
  - Password: `foliogram_pass`

**Using MySQL CLI:**
```bash
docker exec -it foliogram_mysql mysql -u foliogram_user -pfoliogram_pass foliogram_db
```

### 3. Environment Configuration

Copy the example environment file:
```bash
cp .env.example .env
```

Update the `.env` file with your configuration if needed.

## Database Credentials

- **Database Name:** `foliogram_db`
- **Username:** `foliogram_user`
- **Password:** `foliogram_pass`
- **Root Password:** `rootpassword`
- **Port:** `3306`
- **Host:** `localhost` (or `mysql` from within Docker network)

## Useful Commands

### Stop the Database
```bash
docker-compose down
```

### Stop and Remove Data (CAUTION: This will delete all data)
```bash
docker-compose down -v
```

### View Logs
```bash
docker-compose logs -f mysql
```

### Backup Database
```bash
docker exec foliogram_mysql mysqldump -u root -prootpassword foliogram_db > backup.sql
```

### Restore Database
```bash
docker exec -i foliogram_mysql mysql -u root -prootpassword foliogram_db < backup.sql
```

### Access MySQL Shell
```bash
docker exec -it foliogram_mysql mysql -u foliogram_user -pfoliogram_pass foliogram_db
```

## Database Schema

The database schema is automatically created when the container starts. The schema includes:

1. **users** - User accounts and authentication
2. **portfolios** - User portfolio configurations
3. **projects** - Portfolio projects with media
4. **skills** - Portfolio skills
5. **testimonials** - Client testimonials
6. **appointments** - Virtual meeting bookings
7. **blogs** - Blog posts
8. **blog_subscribers** - Blog subscriptions
9. **blog_likes** - Blog likes
10. **blog_comments** - Blog comments
11. **analytics_views** - Portfolio view tracking
12. **notifications** - User notifications

## phpMyAdmin

phpMyAdmin is included for easy database management:

- **URL:** `http://localhost:8080`
- **Server:** `mysql`
- **Username:** `foliogram_user`
- **Password:** `foliogram_pass`

## Troubleshooting

### Container won't start
```bash
# Check container logs
docker-compose logs mysql

# Restart containers
docker-compose restart
```

### Connection refused
- Ensure the container is running: `docker-compose ps`
- Check if port 3306 is already in use: `lsof -i :3306`
- Wait for healthcheck to pass (may take 10-30 seconds)

### Reset everything
```bash
# Stop and remove containers, volumes, and networks
docker-compose down -v

# Start fresh
docker-compose up -d
```

## Production Considerations

Before deploying to production:

1. **Change all passwords** in `docker-compose.yml`
2. **Update `.env`** with production values
3. **Enable SSL/TLS** for MySQL connections
4. **Remove phpMyAdmin** or restrict access
5. **Set up regular backups**
6. **Configure proper firewall rules**
7. **Use Docker secrets** for sensitive data
8. **Enable MySQL slow query log** for performance monitoring

## Connection from Go Application

Update your Beego `app.conf`:

```ini
[database]
driver = mysql
connection = foliogram_user:foliogram_pass@tcp(localhost:3306)/foliogram_db?charset=utf8mb4&parseTime=True&loc=Local
```

Or use environment variables in your application.
