package middlewares

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")

	// Allow specific origins or all origins in development
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:5173",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:5173",
	}

	// Check if origin is allowed
	allowed := false
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			allowed = true
			break
		}
	}

	// Set CORS headers
	if allowed || origin == "" {
		if origin != "" {
			ctx.Output.Header("Access-Control-Allow-Origin", origin)
		} else {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
		}
	} else {
		// In development, allow all origins
		runmode, _ := beego.AppConfig.String("runmode")
		if runmode == "dev" {
			ctx.Output.Header("Access-Control-Allow-Origin", origin)
		}
	}

	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Max-Age", "86400")

	// Handle preflight OPTIONS request
	if ctx.Input.Method() == "OPTIONS" {
		ctx.ResponseWriter.WriteHeader(204)
		ctx.ResponseWriter.Write([]byte(""))
		return
	}
}
