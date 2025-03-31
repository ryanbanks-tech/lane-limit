package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadImage handles image uploads
func UploadImage(c *gin.Context) {
	// Restrict the maximum upload size if desired, e.g. c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}

	// Create an uploads folder if not exists
	uploadPath := "./uploads"
	os.MkdirAll(uploadPath, os.ModePerm)

	// Generate a unique filename (for example, using timestamp)
	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, filename)

	// Save the uploaded file to disk
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
	})
}
