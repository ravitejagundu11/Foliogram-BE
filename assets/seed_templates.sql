-- ============================================================================
-- SEED DATA FOR TEMPLATES
-- ============================================================================

-- Developer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Modern Developer', 'Clean and modern portfolio template perfect for developers and engineers', 'developer', '/templates/modern-dev/preview.png', '/templates/modern-dev/thumb.png', TRUE, FALSE, 
'["Responsive Design", "Dark Mode", "Project Showcase", "Skills Section", "Contact Form"]',
'{"primaryColor": "#3B82F6", "secondaryColor": "#10B981", "fontFamily": "Inter", "layout": "modern"}'),

(UUID(), 'Minimal Tech', 'Minimalist template focused on code and projects', 'developer', '/templates/minimal-tech/preview.png', '/templates/minimal-tech/thumb.png', TRUE, FALSE,
'["Clean Layout", "GitHub Integration", "Code Snippets", "Timeline", "Blog"]',
'{"primaryColor": "#000000", "secondaryColor": "#FFFFFF", "fontFamily": "Roboto Mono", "layout": "minimal"}'),

(UUID(), 'Full Stack Pro', 'Comprehensive template for full-stack developers', 'developer', '/templates/fullstack-pro/preview.png', '/templates/fullstack-pro/thumb.png', TRUE, TRUE,
'["Multi-page", "Tech Stack Showcase", "Testimonials", "Case Studies", "Analytics Dashboard"]',
'{"primaryColor": "#8B5CF6", "secondaryColor": "#EC4899", "fontFamily": "Poppins", "layout": "professional"}');

-- Designer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Creative Studio', 'Bold and creative template for designers and artists', 'designer', '/templates/creative-studio/preview.png', '/templates/creative-studio/thumb.png', TRUE, FALSE,
'["Gallery Grid", "Animation Effects", "Color Palette", "Process Timeline", "Client Logos"]',
'{"primaryColor": "#F59E0B", "secondaryColor": "#EF4444", "fontFamily": "Playfair Display", "layout": "creative"}'),

(UUID(), 'UI/UX Portfolio', 'Showcase your design process and case studies', 'designer', '/templates/uiux-portfolio/preview.png', '/templates/uiux-portfolio/thumb.png', TRUE, TRUE,
'["Case Study Pages", "Figma Integration", "Before/After", "Design System", "Testimonials"]',
'{"primaryColor": "#6366F1", "secondaryColor": "#8B5CF6", "fontFamily": "Inter", "layout": "showcase"}'),

(UUID(), 'Artistic Vision', 'Elegant template for visual artists and illustrators', 'designer', '/templates/artistic-vision/preview.png', '/templates/artistic-vision/thumb.png', TRUE, FALSE,
'["Full-screen Gallery", "Lightbox", "Categories", "About Artist", "Shop Integration"]',
'{"primaryColor": "#EC4899", "secondaryColor": "#F472B6", "fontFamily": "Lora", "layout": "gallery"}');

-- Photographer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Photo Gallery', 'Image-focused template for photographers', 'photographer', '/templates/photo-gallery/preview.png', '/templates/photo-gallery/thumb.png', TRUE, FALSE,
'["Masonry Grid", "Full-screen Slideshow", "EXIF Data", "Collections", "Booking"]',
'{"primaryColor": "#0F172A", "secondaryColor": "#64748B", "fontFamily": "Montserrat", "layout": "gallery"}'),

(UUID(), 'Wedding Photography', 'Romantic template for wedding photographers', 'photographer', '/templates/wedding-photo/preview.png', '/templates/wedding-photo/thumb.png', TRUE, TRUE,
'["Elegant Galleries", "Client Portal", "Package Pricing", "Testimonials", "Contact Form"]',
'{"primaryColor": "#FBBF24", "secondaryColor": "#FCA5A5", "fontFamily": "Cormorant Garamond", "layout": "elegant"}');

-- Business Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Professional Business', 'Corporate template for consultants and agencies', 'business', '/templates/pro-business/preview.png', '/templates/pro-business/thumb.png', TRUE, FALSE,
'["Services Section", "Team Members", "Client Logos", "Statistics", "Newsletter"]',
'{"primaryColor": "#1E40AF", "secondaryColor": "#3B82F6", "fontFamily": "Open Sans", "layout": "corporate"}'),

(UUID(), 'Entrepreneur', 'Dynamic template for entrepreneurs and startups', 'business', '/templates/entrepreneur/preview.png', '/templates/entrepreneur/thumb.png', TRUE, TRUE,
'["Product Showcase", "Metrics Dashboard", "Press Mentions", "Video Integration", "CTA Sections"]',
'{"primaryColor": "#059669", "secondaryColor": "#10B981", "fontFamily": "Space Grotesk", "layout": "modern"}');

-- Content Creator Templates  
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Content Creator', 'Vibrant template for YouTubers and influencers', 'creator', '/templates/content-creator/preview.png', '/templates/content-creator/thumb.png', TRUE, FALSE,
'["Video Embeds", "Social Media Links", "Latest Content", "Subscriber Count", "Merch Shop"]',
'{"primaryColor": "#DC2626", "secondaryColor": "#F59E0B", "fontFamily": "Nunito", "layout": "dynamic"}'),

(UUID(), 'Blogger', 'Clean template for writers and bloggers', 'creator', '/templates/blogger/preview.png', '/templates/blogger/thumb.png', TRUE, FALSE,
'["Blog Posts", "Categories", "Search", "Newsletter", "Social Sharing"]',
'{"primaryColor": "#4338CA", "secondaryColor": "#818CF8", "fontFamily": "Merriweather", "layout": "blog"}');
