package responses

import (
	beego "github.com/beego/beego/v2/server/web"
)

// BaseResponse represents the standard API response structure
type BaseResponse struct {
	Status int                    `json:"status"`
	Data   interface{}            `json:"data,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

// SendSuccess sends a success response with 200 OK
func SendSuccess(c *beego.Controller, data interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = BaseResponse{
		Status: 200,
		Data:   data,
	}
	c.ServeJSON()
}

// SendCreated sends a success response with 201 Created
func SendCreated(c *beego.Controller, data interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(201)
	c.Data["json"] = BaseResponse{
		Status: 201,
		Data:   data,
	}
	c.ServeJSON()
}

// SendBadRequest sends an error response with 400 Bad Request
func SendBadRequest(c *beego.Controller, errors map[string]interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.Data["json"] = BaseResponse{
		Status: 400,
		Errors: errors,
	}
	c.ServeJSON()
}

// SendBadRequestWithMessage sends a bad request with a single error message
func SendBadRequestWithMessage(c *beego.Controller, message string) {
	SendBadRequest(c, map[string]interface{}{
		"message": message,
	})
}

// SendUnauthorized sends an error response with 401 Unauthorized
func SendUnauthorized(c *beego.Controller, message string) {
	c.Ctx.ResponseWriter.WriteHeader(401)
	c.Data["json"] = BaseResponse{
		Status: 401,
		Errors: map[string]interface{}{
			"message": message,
		},
	}
	c.ServeJSON()
}

// SendForbidden sends an error response with 403 Forbidden
func SendForbidden(c *beego.Controller, message string) {
	c.Ctx.ResponseWriter.WriteHeader(403)
	c.Data["json"] = BaseResponse{
		Status: 403,
		Errors: map[string]interface{}{
			"message": message,
		},
	}
	c.ServeJSON()
}

// SendNotFound sends an error response with 404 Not Found
func SendNotFound(c *beego.Controller, message string) {
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.Data["json"] = BaseResponse{
		Status: 404,
		Errors: map[string]interface{}{
			"message": message,
		},
	}
	c.ServeJSON()
}

// SendConflict sends an error response with 409 Conflict
func SendConflict(c *beego.Controller, message string) {
	c.Ctx.ResponseWriter.WriteHeader(409)
	c.Data["json"] = BaseResponse{
		Status: 409,
		Errors: map[string]interface{}{
			"message": message,
		},
	}
	c.ServeJSON()
}

// SendValidationError sends validation errors with 422 Unprocessable Entity
func SendValidationError(c *beego.Controller, errors map[string]interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(422)
	c.Data["json"] = BaseResponse{
		Status: 422,
		Errors: errors,
	}
	c.ServeJSON()
}

// SendInternalServerError sends an error response with 500 Internal Server Error
func SendInternalServerError(c *beego.Controller, message string) {
	c.Ctx.ResponseWriter.WriteHeader(500)
	c.Data["json"] = BaseResponse{
		Status: 500,
		Errors: map[string]interface{}{
			"message": message,
		},
	}
	c.ServeJSON()
}
