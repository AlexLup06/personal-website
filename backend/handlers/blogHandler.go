package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"alexlupatsiy.com/personal-website/backend/helpers"
	"alexlupatsiy.com/personal-website/frontend/src/views/blog"
	"github.com/gin-gonic/gin"
)

const pathToRoot = "./"

func BlogHandler(ctx *gin.Context, blogMetadata map[string]blog.BlogType) {
	slug := ctx.Param("blog")

	// Check if blog exists
	blogData, exists := blogMetadata[slug]
	if !exists {
		ctx.String(http.StatusNotFound, "Blog not found")
		return
	}

	// Read blog content from HTML file
	contentPath := filepath.Join(pathToRoot+"frontend/src/views/blog/blog-html", slug+".html")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error reading blog content")
		return
	}

	// Render blog page with Templ
	helpers.Render(ctx, 200, blog.Blog(blogData.Title, blogData.Date, blogData.Intro, string(content)))
}

func LoadBlogMetadata() (map[string]blog.BlogType, error) {
	metadataPath := filepath.Join(pathToRoot+"blog-generator", "blogMetadata.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		fmt.Println("Error reading blog metadata:", err)
		return map[string]blog.BlogType{}, err
	}

	var blogs struct {
		Blogs []blog.BlogType `json:"blogs"`
	}
	if err := json.Unmarshal(data, &blogs); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return map[string]blog.BlogType{}, err
	}

	blogMetadata := map[string]blog.BlogType{}
	for _, b := range blogs.Blogs {
		blogMetadata[b.Slug] = b
	}
	return blogMetadata, nil
}
