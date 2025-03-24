package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"alexlupatsiy.com/personal-website/backend/helpers"
)

type BlogType struct {
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Intro string   `json:"intro"`
	Slug  string   `json:"slug"`
	Tags  []string `json:"tags"`
}

type BlogService struct {
	blogsMetadata []BlogType
	pathToRoot    string
}

func NewBlogService() *BlogService {
	const pathToRoot string = "./"
	var blogsMetadata []BlogType
	blogsMetadata, err := initBlogService(pathToRoot)
	if err != nil {
		fmt.Printf("Error reading blog metadata: %v\n", err)
	}

	fmt.Println("blogsMetadata: ", blogsMetadata)

	return &BlogService{blogsMetadata: blogsMetadata}
}

func initBlogService(pathToRoot string) ([]BlogType, error) {
	metadataPath := filepath.Join(pathToRoot+"blog-html", "blogMetadata.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		fmt.Println("Error reading blog metadata:", err)
		return []BlogType{}, err
	}

	var blogs struct {
		Blogs []BlogType `json:"blogs"`
	}
	if err := json.Unmarshal(data, &blogs); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return []BlogType{}, err
	}

	blogsMetadata := []BlogType{}
	for _, b := range blogs.Blogs {
		blogsMetadata = append(blogsMetadata, b)
	}

	return blogsMetadata, nil
}

func (b *BlogService) GetFilteredMetadata(filter []string) []BlogType {
	if len(filter) == 0 {
		return b.blogsMetadata
	}
	blogs := []BlogType{}
	for _, blogMetadata := range b.blogsMetadata {
		intersectingTags := helpers.Intersect(blogMetadata.Tags, filter)
		if len(intersectingTags) != 0 {
			blogs = append(blogs, blogMetadata)
		}
	}

	return blogs
}

func (b *BlogService) GetBlogMetadata(slug string) (BlogType, bool) {
	for _, value := range b.blogsMetadata {
		if value.Slug == slug {
			return value, true
		}
	}
	return BlogType{}, false
}

func (b *BlogService) GetBlogContent(slug string) (string, error) {
	contentPath := filepath.Join(b.pathToRoot+"blog-html", slug+".html")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (b *BlogService) GetMostRecentBlogs() []BlogType {
	return b.blogsMetadata[len(b.blogsMetadata)-3:]
}
