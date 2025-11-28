package middlewares

import (
	"Foliogram-BE/utils"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

// AuthMiddleware validates JWT token and sets user context
func AuthMiddleware(ctx *context.Context) {
	// Skip authentication for public routes
	path := ctx.Request.URL.Path
	if path == "/api/foliogram/register" || path == "/api/foliogram/login" {
		return
	}

	// Get token from Authorization header
	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Output.JSON(map[string]interface{}{
			"status": 401,
			"errors": map[string]interface{}{
				"message": "Authorization header required",
			},
		}, false, false)
		return
	}

	// Extract token (format: "Bearer <token>")
	var token string
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = authHeader[7:]
	} else {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Output.JSON(map[string]interface{}{
			"status": 401,
			"errors": map[string]interface{}{
				"message": "Invalid authorization header format",
			},
		}, false, false)
		return
	}

	// Validate token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Output.JSON(map[string]interface{}{
			"status": 401,
			"errors": map[string]interface{}{
				"message": "Invalid or expired token",
			},
		}, false, false)
		return
	}

	// Check if user session is valid
	if !utils.IsSessionValid(claims.UserID, claims.ExpiresAt.Time) {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Output.JSON(map[string]interface{}{
			"status": 401,
			"errors": map[string]interface{}{
				"message": "Session has expired or been invalidated",
			},
		}, false, false)
		return
	}

	// Set user context
	ctx.Input.SetData("user_id", claims.UserID)
	ctx.Input.SetData("user_email", claims.Email)
	ctx.Input.SetData("user_role", claims.Role)
}

// AdminMiddleware validates JWT token and checks for admin role
func AdminMiddleware(ctx *context.Context) {
	// First validate the token
	AuthMiddleware(ctx)

	// Check if previous middleware returned an error
	if ctx.ResponseWriter.Status != 0 && ctx.ResponseWriter.Status != 200 {
		return
	}

	// Check if user has admin role
	role := ctx.Input.GetData("user_role")
	if role != "admin" {
		ctx.ResponseWriter.WriteHeader(403)
		ctx.Output.JSON(map[string]interface{}{
			"status": 403,
			"errors": map[string]interface{}{
				"message": "Admin access required",
			},
		}, false, false)
		return
	}
}
