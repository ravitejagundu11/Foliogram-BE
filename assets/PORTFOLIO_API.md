# Portfolio Builder API Documentation

## Overview
This document describes the portfolio builder functionality that allows users to create and customize their portfolios using pre-defined templates.

## Database Setup

### 1. Run Schema Migration
The database schema includes tables for templates, portfolios, projects, skills, and testimonials with proper foreign key relationships.

```bash
# The schema will be auto-migrated when the server starts
# Or manually run the SQL schema:
mysql -u foliogram_user -p foliogram_db < assets/schema_mysql.sql
```

### 2. Seed Templates
Populate the database with initial templates:

```bash
mysql -u foliogram_user -p foliogram_db < assets/seed_templates.sql
```

## API Endpoints

### Templates

#### Get All Templates
```http
GET /api/foliogram/templates
Query Parameters:
  - category: Filter by category (developer, designer, photographer, business, creator)
  - premium: Filter premium templates (true/false)
```

#### Get Template by ID
```http
GET /api/foliogram/templates/:id
```

#### Get Template Categories
```http
GET /api/foliogram/templates/categories
```

### Portfolios

#### Create Portfolio
```http
POST /api/foliogram/portfolios
Authorization: Bearer <token>

Body:
{
  "template_id": "uuid",
  "title": "My Portfolio",
  "description": "Portfolio description"
}
```

#### Get User's Portfolios
```http
GET /api/foliogram/portfolios/my
Authorization: Bearer <token>
```

#### Get Portfolio by ID
```http
GET /api/foliogram/portfolios/:id
Authorization: Bearer <token> (optional for published portfolios)
```

#### Get Portfolio by Slug
```http
GET /api/foliogram/portfolios/slug/:slug
```

#### Update Portfolio
```http
PUT /api/foliogram/portfolios/:id
Authorization: Bearer <token>

Body:
{
  "title": "Updated Title",
  "description": "Updated description",
  "config": {
    "primaryColor": "#3B82F6",
    "secondaryColor": "#10B981",
    "fontFamily": "Inter",
    "customCSS": "..."
  },
  "is_published": true,
  "meta_title": "SEO Title",
  "meta_description": "SEO Description",
  "meta_keywords": "keyword1, keyword2"
}
```

#### Delete Portfolio
```http
DELETE /api/foliogram/portfolios/:id
Authorization: Bearer <token>
```

### File Uploads

#### Upload Single Image
```http
POST /api/foliogram/upload/image
Authorization: Bearer <token>
Content-Type: multipart/form-data

Form Data:
  - file: <image file>

Allowed formats: jpg, jpeg, png, gif, webp, svg
Max size: 10MB
```

#### Upload Multiple Images
```http
POST /api/foliogram/upload/multiple
Authorization: Bearer <token>
Content-Type: multipart/form-data

Form Data:
  - files: <image files>

Max files: 10
```

#### Delete Uploaded File
```http
DELETE /api/foliogram/upload/:filename
Authorization: Bearer <token>
```

## Models

### Portfolio Model
- ID (UUID)
- UserID (FK to users)
- TemplateID (FK to templates)
- Title
- Slug (unique, URL-friendly)
- Description
- IsPublished
- Config (JSON - colors, fonts, layout preferences)
- SEO fields (meta_title, meta_description, meta_keywords)
- ViewCount
- Associations: Projects, Skills, Testimonials

### Template Model
- ID (UUID)
- Name
- Description
- Category (developer, designer, photographer, business, creator)
- PreviewURL
- ThumbnailURL
- IsActive
- IsPremium
- Features (JSON array)
- DefaultConfig (JSON object)

### Project Model
- ID (UUID)
- PortfolioID (FK)
- Title
- Description
- ImageURL
- GalleryURLs (JSON array)
- DemoURL
- GithubURL
- Tags (JSON array)
- Technologies (JSON array)
- StartDate, EndDate
- IsFeatured
- SortOrder

### Skill Model
- ID (UUID)
- PortfolioID (FK)
- Name
- Category (Frontend, Backend, Tools, etc.)
- Level (Beginner, Intermediate, Expert)
- Percentage (0-100)
- IconURL
- SortOrder

### Testimonial Model
- ID (UUID)
- PortfolioID (FK)
- Name
- Position
- Company
- Content
- Rating (1-5)
- AvatarURL
- LinkedInURL
- IsVisible
- SortOrder

## GORM Associations

The models use GORM's association features:
- `Portfolio` has many `Projects`, `Skills`, and `Testimonials`
- `Portfolio` belongs to `Template` and `User`
- Cascade delete is enabled (deleting a portfolio deletes all associated data)

## File Upload Configuration

Uploaded files are stored in `./uploads/images/` directory and served via the `/uploads` route.

Configure static file serving in `conf/app.conf`:
```ini
StaticDir = uploads:uploads
```

## Usage Example

1. **Select a Template**
   ```bash
   GET /api/foliogram/templates?category=developer
   ```

2. **Create Portfolio**
   ```bash
   POST /api/foliogram/portfolios
   {
     "template_id": "abc-123",
     "title": "John's Portfolio"
   }
   ```

3. **Upload Project Images**
   ```bash
   POST /api/foliogram/upload/image
   (multipart form with image file)
   ```

4. **Customize Portfolio**
   ```bash
   PUT /api/foliogram/portfolios/xyz-789
   {
     "config": {
       "primaryColor": "#FF6B6B",
       "fontFamily": "Roboto"
     }
   }
   ```

5. **Publish Portfolio**
   ```bash
   PUT /api/foliogram/portfolios/xyz-789
   {
     "is_published": true
   }
   ```

## Security Notes

- All portfolio modification endpoints require authentication
- File uploads are restricted to authenticated users
- File size limited to 10MB
- Only image formats are allowed for uploads
- Published portfolios are publicly accessible
- Unpublished portfolios are only accessible by the owner
