-- Foliogram Database Schema (MySQL)
-- Created: November 26, 2025
-- Description: MySQL schema for Foliogram portfolio management platform

-- ============================================================================
-- USERS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    mobile VARCHAR(20),
    location VARCHAR(255),
    profile_picture TEXT,
    role ENUM('user', 'admin') NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_users_email (email),
    INDEX idx_users_role (role)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- PORTFOLIOS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS portfolios (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36) NOT NULL,
    headline VARCHAR(255),
    description TEXT,
    template_id INT,
    primary_color VARCHAR(7),
    secondary_color VARCHAR(7),
    heading_font VARCHAR(100),
    body_font VARCHAR(100),
    layout_style ENUM('minimal', 'classic', 'modern') DEFAULT 'minimal',
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_portfolios_user_id (user_id),
    INDEX idx_portfolios_is_published (is_published),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- PROJECTS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS projects (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    portfolio_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    technologies JSON,
    media_urls JSON,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_projects_portfolio_id (portfolio_id),
    FOREIGN KEY (portfolio_id) REFERENCES portfolios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- SKILLS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS skills (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    portfolio_id CHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_skills_portfolio_id (portfolio_id),
    FOREIGN KEY (portfolio_id) REFERENCES portfolios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- TESTIMONIALS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS testimonials (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    portfolio_id CHAR(36) NOT NULL,
    text TEXT NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    author_role VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_testimonials_portfolio_id (portfolio_id),
    FOREIGN KEY (portfolio_id) REFERENCES portfolios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- APPOINTMENTS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS appointments (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    owner_id CHAR(36) NOT NULL,
    attendee_id CHAR(36) NOT NULL,
    date DATE NOT NULL,
    time TIME NOT NULL,
    status ENUM('booked', 'confirmed', 'cancelled') NOT NULL DEFAULT 'booked',
    platform ENUM('zoom', 'skype'),
    meeting_link TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_appointments_owner_id (owner_id),
    INDEX idx_appointments_attendee_id (attendee_id),
    INDEX idx_appointments_date (date),
    INDEX idx_appointments_status (status),
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (attendee_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- BLOGS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS blogs (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    subscriber_count INT DEFAULT 0,
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_blogs_user_id (user_id),
    INDEX idx_blogs_is_published (is_published),
    INDEX idx_blogs_created_at (created_at DESC),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- BLOG_SUBSCRIBERS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS blog_subscribers (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    blog_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    subscribed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_blog_subscribers_blog_id (blog_id),
    INDEX idx_blog_subscribers_user_id (user_id),
    UNIQUE KEY unique_blog_subscriber (blog_id, user_id),
    FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- BLOG_LIKES TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS blog_likes (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    blog_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    liked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_blog_likes_blog_id (blog_id),
    INDEX idx_blog_likes_user_id (user_id),
    UNIQUE KEY unique_blog_like (blog_id, user_id),
    FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- BLOG_COMMENTS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS blog_comments (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    blog_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_blog_comments_blog_id (blog_id),
    INDEX idx_blog_comments_user_id (user_id),
    FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- ANALYTICS_VIEWS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS analytics_views (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    portfolio_id CHAR(36) NOT NULL,
    viewer_ip VARCHAR(45),
    session_id VARCHAR(255),
    viewed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_analytics_views_portfolio_id (portfolio_id),
    INDEX idx_analytics_views_viewed_at (viewed_at DESC),
    FOREIGN KEY (portfolio_id) REFERENCES portfolios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- NOTIFICATIONS TABLE
-- ============================================================================
CREATE TABLE IF NOT EXISTS notifications (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36) NOT NULL,
    type ENUM('appointment', 'blog_update') NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_notifications_user_id (user_id),
    INDEX idx_notifications_is_read (is_read),
    INDEX idx_notifications_created_at (created_at DESC),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
