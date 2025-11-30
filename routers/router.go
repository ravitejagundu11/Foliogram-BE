package routers

import (
	"Foliogram-BE/controllers"
	"Foliogram-BE/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})

	// API namespace with middleware (skips auth for login/register in middleware)
	apiNs := beego.NewNamespace("/api/foliogram",
		beego.NSBefore(middlewares.AuthMiddleware),
		beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "post:Logout"),
		// Add more protected routes here
		// beego.NSRouter("/profile", &controllers.UserController{}, "get:GetProfile"),
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
