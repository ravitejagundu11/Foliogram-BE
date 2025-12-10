-- ============================================================================
-- SEED DATA FOR TEMPLATES
-- ============================================================================

-- Developer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Modern Developer', 'Clean and modern portfolio template perfect for developers and engineers', 'developer', 'https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=400&h=300&fit=crop', TRUE, FALSE, 
'["Responsive Design", "Dark Mode", "Project Showcase", "Skills Section", "Contact Form"]',
'{"primaryColor": "#3B82F6", "secondaryColor": "#10B981", "fontFamily": "Inter", "layout": "modern"}'),

(UUID(), 'Minimal Tech', 'Minimalist template focused on code and projects', 'developer', 'https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=400&h=300&fit=crop', TRUE, FALSE,
'["Clean Layout", "GitHub Integration", "Code Snippets", "Timeline", "Blog"]',
'{"primaryColor": "#000000", "secondaryColor": "#FFFFFF", "fontFamily": "Roboto Mono", "layout": "minimal"}'),

(UUID(), 'Full Stack Pro', 'Comprehensive template for full-stack developers', 'developer', 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=400&h=300&fit=crop', TRUE, TRUE,
'["Multi-page", "Tech Stack Showcase", "Testimonials", "Case Studies", "Analytics Dashboard"]',
'{"primaryColor": "#8B5CF6", "secondaryColor": "#EC4899", "fontFamily": "Poppins", "layout": "professional"}');

-- Designer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Creative Studio', 'Bold and creative template for designers and artists', 'designer', 'https://images.unsplash.com/photo-1561070791-2526d30994b5?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1561070791-2526d30994b5?w=400&h=300&fit=crop', TRUE, FALSE,
'["Gallery Grid", "Animation Effects", "Color Palette", "Process Timeline", "Client Logos"]',
'{"primaryColor": "#F59E0B", "secondaryColor": "#EF4444", "fontFamily": "Playfair Display", "layout": "creative"}'),

(UUID(), 'UI/UX Portfolio', 'Showcase your design process and case studies', 'designer', 'https://images.unsplash.com/photo-1559028012-481c04fa702d?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1559028012-481c04fa702d?w=400&h=300&fit=crop', TRUE, TRUE,
'["Case Study Pages", "Figma Integration", "Before/After", "Design System", "Testimonials"]',
'{"primaryColor": "#6366F1", "secondaryColor": "#8B5CF6", "fontFamily": "Inter", "layout": "showcase"}'),

(UUID(), 'Artistic Vision', 'Elegant template for visual artists and illustrators', 'designer', 'https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=400&h=300&fit=crop', TRUE, FALSE,
'["Full-screen Gallery", "Lightbox", "Categories", "About Artist", "Shop Integration"]',
'{"primaryColor": "#EC4899", "secondaryColor": "#F472B6", "fontFamily": "Lora", "layout": "gallery"}');

-- Photographer Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Photo Gallery', 'Image-focused template for photographers', 'photographer', 'https://images.unsplash.com/photo-1452587925148-ce544e77e70d?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1452587925148-ce544e77e70d?w=400&h=300&fit=crop', TRUE, FALSE,
'["Masonry Grid", "Full-screen Slideshow", "EXIF Data", "Collections", "Booking"]',
'{"primaryColor": "#0F172A", "secondaryColor": "#64748B", "fontFamily": "Montserrat", "layout": "gallery"}'),

(UUID(), 'Wedding Photography', 'Romantic template for wedding photographers', 'photographer', 'https://images.unsplash.com/photo-1519741497674-611481863552?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1519741497674-611481863552?w=400&h=300&fit=crop', TRUE, TRUE,
'["Elegant Galleries", "Client Portal", "Package Pricing", "Testimonials", "Contact Form"]',
'{"primaryColor": "#FBBF24", "secondaryColor": "#FCA5A5", "fontFamily": "Cormorant Garamond", "layout": "elegant"}');

-- Business Templates
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Professional Business', 'Corporate template for consultants and agencies', 'business', 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?w=400&h=300&fit=crop', TRUE, FALSE,
'["Services Section", "Team Members", "Client Logos", "Statistics", "Newsletter"]',
'{"primaryColor": "#1E40AF", "secondaryColor": "#3B82F6", "fontFamily": "Open Sans", "layout": "corporate"}'),

(UUID(), 'Entrepreneur', 'Dynamic template for entrepreneurs and startups', 'business', 'https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=400&h=300&fit=crop', TRUE, TRUE,
'["Product Showcase", "Metrics Dashboard", "Press Mentions", "Video Integration", "CTA Sections"]',
'{"primaryColor": "#059669", "secondaryColor": "#10B981", "fontFamily": "Space Grotesk", "layout": "modern"}');

-- Content Creator Templates  
INSERT INTO templates (id, name, description, category, preview_url, thumbnail_url, is_active, is_premium, features, default_config) VALUES
(UUID(), 'Content Creator', 'Vibrant template for YouTubers and influencers', 'creator', 'https://images.unsplash.com/photo-1611162617474-5b21e879e113?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1611162617474-5b21e879e113?w=400&h=300&fit=crop', TRUE, FALSE,
'["Video Embeds", "Social Media Links", "Latest Content", "Subscriber Count", "Merch Shop"]',
'{"primaryColor": "#DC2626", "secondaryColor": "#F59E0B", "fontFamily": "Nunito", "layout": "dynamic"}'),

(UUID(), 'Blogger', 'Clean template for writers and bloggers', 'creator', 'https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=800&h=600&fit=crop', 'https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=400&h=300&fit=crop', TRUE, FALSE,
'["Blog Posts", "Categories", "Search", "Newsletter", "Social Sharing"]',
'{"primaryColor": "#4338CA", "secondaryColor": "#818CF8", "fontFamily": "Merriweather", "layout": "blog"}');

