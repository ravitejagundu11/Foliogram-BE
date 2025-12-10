package routers

import (
	"Foliogram-BE/controllers"
	"Foliogram-BE/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Apply CORS middleware globally
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.CORSMiddleware)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})

	// API namespace with middleware (skips auth for login/register in middleware)
	apiNs := beego.NewNamespace("/api/foliogram",
		beego.NSBefore(middlewares.AuthMiddleware),
		beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "post:Logout"),
		beego.NSRouter("/change-password", &controllers.AuthController{}, "put:ChangePassword"),
		beego.NSRouter("/change-email", &controllers.AuthController{}, "put:ChangeEmail"),
		beego.NSRouter("/profile", &controllers.AuthController{}, "put:UpdateProfile"),

		// Template routes
		beego.NSRouter("/templates", &controllers.TemplateController{}, "get:GetAll"),
		beego.NSRouter("/templates/categories", &controllers.TemplateController{}, "get:GetCategories"),
		beego.NSRouter("/templates/:id", &controllers.TemplateController{}, "get:GetByID"),

		// Portfolio routes
		beego.NSRouter("/portfolios", &controllers.PortfolioController{}, "post:CreatePortfolio"),
		beego.NSRouter("/portfolios/my-portfolios", &controllers.PortfolioController{}, "get:GetMyPortfolios"),
		beego.NSRouter("/portfolios/:id", &controllers.PortfolioController{}, "get:GetPortfolioByID;put:UpdatePortfolio;delete:DeletePortfolio"),
		beego.NSRouter("/portfolios/:id/publish", &controllers.PortfolioController{}, "put:PublishPortfolio"),
		beego.NSRouter("/portfolios/:id/unpublish", &controllers.PortfolioController{}, "put:UnpublishPortfolio"),

		// Project routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/projects", &controllers.ProjectController{}, "post:CreateProject;get:GetProjectsByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/projects/:projectId", &controllers.ProjectController{}, "get:GetProjectByID;put:UpdateProject;delete:DeleteProject"),

		// Skill routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/skills", &controllers.SkillController{}, "post:CreateSkill;get:GetSkillsByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/skills/:skillId", &controllers.SkillController{}, "get:GetSkillByID;put:UpdateSkill;delete:DeleteSkill"),

		// Testimonial routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/testimonials", &controllers.TestimonialController{}, "post:CreateTestimonial;get:GetTestimonialsByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/testimonials/:testimonialId", &controllers.TestimonialController{}, "get:GetTestimonialByID;put:UpdateTestimonial;delete:DeleteTestimonial"),

		// Education routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/education", &controllers.EducationController{}, "post:CreateEducation;get:GetEducationByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/education/:educationId", &controllers.EducationController{}, "get:GetEducationByID;put:UpdateEducation;delete:DeleteEducation"),

		// Experience routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/experience", &controllers.ExperienceController{}, "post:CreateExperience;get:GetExperienceByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/experience/:experienceId", &controllers.ExperienceController{}, "get:GetExperienceByID;put:UpdateExperience;delete:DeleteExperience"),

		// Publication routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/publications", &controllers.PublicationController{}, "post:CreatePublication;get:GetPublicationsByPortfolio"),
		beego.NSRouter("/portfolios/:portfolioId/publications/:publicationId", &controllers.PublicationController{}, "get:GetPublicationByID;put:UpdatePublication;delete:DeletePublication"),

		// Contact routes (nested under portfolios)
		beego.NSRouter("/portfolios/:portfolioId/contact", &controllers.ContactController{}, "post:CreateOrUpdateContact;get:GetContactByPortfolio;delete:DeleteContact"),

		// Upload routes
		beego.NSRouter("/upload", &controllers.UploadController{}, "post:UploadImage"),
		beego.NSRouter("/upload/image", &controllers.UploadController{}, "post:UploadImage"),
		beego.NSRouter("/upload/multiple", &controllers.UploadController{}, "post:UploadMultiple"),

		// Add more protected routes here
		// beego.NSRouter("/profile", &controllers.UserController{}, "get:GetProfile"),
	)
	beego.AddNamespace(apiNs)

	// Public routes (no authentication required)
	publicNs := beego.NewNamespace("/api/foliogram/public",
		beego.NSRouter("/portfolios/slug/:slug", &controllers.PortfolioController{}, "get:GetPublicPortfolioBySlug"),
		beego.NSRouter("/portfolios/:id", &controllers.PortfolioController{}, "get:GetPublicPortfolio"),
	)
	beego.AddNamespace(publicNs)

	// Admin routes - require admin role
	adminNs := beego.NewNamespace("/api/foliogram/admin",
		beego.NSBefore(middlewares.AdminMiddleware),
		// Add admin routes here
		// beego.NSRouter("/users", &controllers.AdminController{}, "get:GetAllUsers"),
	)
	beego.AddNamespace(adminNs)
}
