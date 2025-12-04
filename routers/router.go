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

	// Public API namespace (no authentication required)
	publicNs := beego.NewNamespace("/api/foliogram/public",
		// Public portfolio routes
		beego.NSRouter("/portfolios", &controllers.PortfolioController{}, "get:GetPublicPortfolios"),
		beego.NSRouter("/portfolios/:id", &controllers.PortfolioController{}, "get:GetPublicByID"),
		beego.NSRouter("/portfolios/slug/:slug", &controllers.PortfolioController{}, "get:GetPublicBySlug"),
	)
	beego.AddNamespace(publicNs)

	// API namespace with middleware (skips auth for login/register in middleware)
	apiNs := beego.NewNamespace("/api/foliogram",
		beego.NSBefore(middlewares.AuthMiddleware),
		beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "post:Logout"),

		// Template routes (public)
		beego.NSRouter("/templates", &controllers.TemplateController{}, "get:GetAll"),
		beego.NSRouter("/templates/categories", &controllers.TemplateController{}, "get:GetCategories"),
		beego.NSRouter("/templates/:id", &controllers.TemplateController{}, "get:GetByID"),

		// Portfolio routes
		beego.NSRouter("/portfolios", &controllers.PortfolioController{}, "post:Create"),
		beego.NSRouter("/portfolios/my", &controllers.PortfolioController{}, "get:GetUserPortfolios"),
		beego.NSRouter("/portfolios/:id", &controllers.PortfolioController{}, "get:GetByID;put:Update;delete:Delete"),
		beego.NSRouter("/portfolios/slug/:slug", &controllers.PortfolioController{}, "get:GetBySlug"),

		// Upload routes (protected)
		beego.NSRouter("/upload/image", &controllers.UploadController{}, "post:UploadImage"),
		beego.NSRouter("/upload/multiple", &controllers.UploadController{}, "post:UploadMultiple"),
		beego.NSRouter("/upload/:filename", &controllers.UploadController{}, "delete:DeleteFile"),
	)
	beego.AddNamespace(apiNs)

	// Admin routes - require admin role
	adminNs := beego.NewNamespace("/api/foliogram/admin",
		beego.NSBefore(middlewares.AdminMiddleware),
		// Add admin routes here
		// beego.NSRouter("/users", &controllers.AdminController{}, "get:GetAllUsers"),
	)
	beego.AddNamespace(adminNs)
}
