package routers

import (
	"Foliogram-BE/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})
	
	// API routes
	apiNs := beego.NewNamespace("/api/foliogram",
		beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
	)
	beego.AddNamespace(apiNs)
}
