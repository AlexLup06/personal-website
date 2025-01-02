package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware to serve .gz files
func ServeGzippedFiles(mode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the client accepts gzip
		acceptEncoding := c.GetHeader("Accept-Encoding")
		c.Writer.Header().Set("Cache-Control", "public, max-age=31536000")
		if strings.Contains(acceptEncoding, "gzip") {
			// Check if a .gz version of the file exists
			requestedFile := c.Request.URL.Path
			gzippedFile := requestedFile + ".gz"
			if _, err := http.Dir("./frontend/src").Open(gzippedFile); err == nil && mode == "production" {
				// Serve the .gz file
				c.Writer.Header().Set("Content-Encoding", "gzip")
				c.Writer.Header().Set("Content-Type", resolveContentType(requestedFile))
				c.File("./frontend/src" + gzippedFile)
				c.Abort()
				return
			}

		}

		// Continue to the next handler if no .gz file is found
		c.Next()
	}
}

func ServeStaticFiles(basePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := c.Param("filepath")
		c.File(basePath + filePath)
	}
}

// Helper to resolve Content-Type based on file extension
func resolveContentType(filePath string) string {
	if strings.HasSuffix(filePath, ".css") {
		return "text/css"
	} else if strings.HasSuffix(filePath, ".js") {
		return "application/javascript"
	}
	return "application/octet-stream"
}
