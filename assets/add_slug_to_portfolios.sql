-- Add slug column to portfolios table
ALTER TABLE portfolios 
ADD COLUMN slug VARCHAR(255) AFTER user_id,
ADD UNIQUE INDEX idx_portfolios_slug (slug);
