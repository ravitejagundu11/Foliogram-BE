package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	beego "github.com/beego/beego/v2/server/web"
)

// UploadController handles file uploads
type UploadController struct {
	beego.Controller
}

// allowedImageExtensions defines allowed image file extensions
var allowedImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".svg":  true,
}

// maxFileSize defines maximum file size (10MB)
const maxFileSize = 10 * 1024 * 1024

// UploadImage handles POST /api/foliogram/upload/image
func (c *UploadController) UploadImage() {
	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	// Get uploaded file
	file, header, err := c.GetFile("file")
	if err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "No file uploaded",
			"error":   err.Error(),
		})
		return
	}
	defer file.Close()

	// Validate file size
	if header.Size > maxFileSize {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": fmt.Sprintf("File size exceeds maximum limit of %d MB", maxFileSize/(1024*1024)),
		})
		return
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedImageExtensions[ext] {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid file type. Allowed types: jpg, jpeg, png, gif, webp, svg",
		})
		return
	}

	// Create upload directory if it doesn't exist
	uploadDir := "./uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filepath := filepath.Join(uploadDir, filename)

	// Save file
	if err := c.SaveToFile("file", filepath); err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to save file",
			"error":   err.Error(),
		})
		return
	}

	// Return file URL (adjust based on your server configuration)
	fileURL := fmt.Sprintf("/uploads/images/%s", filename)

	SendCreated(&c.Controller, map[string]interface{}{
		"url":      fileURL,
		"filename": filename,
		"size":     header.Size,
	})
}

// UploadMultiple handles POST /api/foliogram/upload/multiple
func (c *UploadController) UploadMultiple() {
	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	// Get all uploaded files
	files := c.Ctx.Request.MultipartForm.File["files"]
	if len(files) == 0 {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "No files uploaded",
		})
		return
	}

	// Limit number of files
	maxFiles := 10
	if len(files) > maxFiles {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": fmt.Sprintf("Maximum %d files allowed", maxFiles),
		})
		return
	}

	// Create upload directory if it doesn't exist
	uploadDir := "./uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
		return
	}

	uploadedFiles := []map[string]interface{}{}
	errors := []string{}

	for _, fileHeader := range files {
		// Validate file size
		if fileHeader.Size > maxFileSize {
			errors = append(errors, fmt.Sprintf("%s: file size exceeds limit", fileHeader.Filename))
			continue
		}

		// Validate file extension
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !allowedImageExtensions[ext] {
			errors = append(errors, fmt.Sprintf("%s: invalid file type", fileHeader.Filename))
			continue
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: failed to open file", fileHeader.Filename))
			continue
		}
		defer file.Close()

		// Generate unique filename
		filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
		savePath := filepath.Join(uploadDir, filename)

		// Save file
		if err := c.SaveToFile("files", savePath); err != nil {
			errors = append(errors, fmt.Sprintf("%s: failed to save file", fileHeader.Filename))
			continue
		}

		// Return file URL
		fileURL := fmt.Sprintf("/uploads/images/%s", filename)
		uploadedFiles = append(uploadedFiles, map[string]interface{}{
			"url":      fileURL,
			"filename": filename,
			"size":     fileHeader.Size,
		})
	}

	result := map[string]interface{}{
		"uploaded": uploadedFiles,
		"count":    len(uploadedFiles),
	}

	if len(errors) > 0 {
		result["errors"] = errors
	}

	SendCreated(&c.Controller, result)
}

// DeleteFile handles DELETE /api/foliogram/upload/:filename
func (c *UploadController) DeleteFile() {
	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	filename := c.Ctx.Input.Param(":filename")
	if filename == "" {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Filename is required",
		})
		return
	}

	// Validate filename (prevent directory traversal)
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid filename",
		})
		return
	}

	// Delete file
	filepath := filepath.Join("./uploads/images", filename)
	if err := os.Remove(filepath); err != nil {
		if os.IsNotExist(err) {
			SendNotFound(&c.Controller, map[string]interface{}{
				"message": "File not found",
			})
			return
		}
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to delete file",
			"error":   err.Error(),
		})
		return
	}

	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "File deleted successfully",
	})
}
