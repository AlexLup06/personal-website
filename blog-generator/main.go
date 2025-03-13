package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

const pathToRoot string = "../"
const pathToBlogs string = pathToRoot + "blogs"

func mdToHTML(md []byte) []byte {
	mdConverter := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	var buf bytes.Buffer
	err := mdConverter.Convert(md, &buf)
	if err != nil {
		return nil // Return nil if conversion fails (handle error appropriately in your app)
	}
	return buf.Bytes()
	// return nil
}

func main() {
	err := filepath.Walk(pathToBlogs, processMarkdown)
	if err != nil {
		fmt.Println(err)
	}
}

func processMarkdown(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var title, date, intro string
	var contentLines []string
	inContent := false

	for scanner.Scan() {
		line := scanner.Text()

		// If we encounter an empty line, assume content starts
		if line == "" {
			inContent = true
			continue
		}

		// If still in metadata, extract values
		if !inContent {
			if strings.HasPrefix(line, "Title: ") {
				title = strings.TrimPrefix(line, "Title: ")
			} else if strings.HasPrefix(line, "Date: ") {
				date = strings.TrimPrefix(line, "Date: ")
			} else if strings.HasPrefix(line, "Intro: ") {
				intro = strings.TrimPrefix(line, "Intro: ")
			}
		} else {
			// Collect the remaining lines as blog content
			contentLines = append(contentLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Join the content lines into a single string
	content := strings.Join(contentLines, "\n")

	// generate slug
	slug := generateSlug(title, date)

	blogMetadata := BlogMetadata{
		Title: title,
		Date:  date,
		Intro: intro,
		Slug:  slug,
	}

	// append Metadata to project-root/blog-generator/blogMetadata.json
	allBlogMetadata, err := loadBlogMetadata()
	if err != nil {
		return err
	}

	allBlogMetadata[blogMetadata.Slug] = blogMetadata
	err = writeBlogMetadata(allBlogMetadata)
	if err != nil {
		return err
	}

	// write HTML to file in project-root/frontend/src/views/blog/blog-html/<<slug>>.html
	html := mdToHTML([]byte(content))

	err = writeToFile(pathToRoot+"blog-html/"+slug+".html", html)

	return nil
}

type BlogMetadata struct {
	Title string `json:"title"`
	Date  string `json:"date"`
	Intro string `json:"intro"`
	Slug  string `json:"slug"`
}

func loadBlogMetadata() (map[string]BlogMetadata, error) {
	metadataPath := filepath.Join(pathToRoot+"blog-html", "blogMetadata.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		fmt.Println("Error reading blog metadata:", err)
		return map[string]BlogMetadata{}, err
	}

	var blogs struct {
		Blogs []BlogMetadata `json:"blogs"`
	}
	if err := json.Unmarshal(data, &blogs); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return map[string]BlogMetadata{}, err
	}

	blogMetadata := map[string]BlogMetadata{}
	for _, b := range blogs.Blogs {
		blogMetadata[b.Slug] = b
	}
	return blogMetadata, nil
}

func generateSlug(title, date string) string {
	titleSeperatedByMinus := strings.ReplaceAll(title, " ", "-")
	return date + "_" + titleSeperatedByMinus
}

func writeBlogMetadata(blogMetadata map[string]BlogMetadata) error {
	metadataPath := filepath.Join(pathToRoot+"blog-html", "blogMetadata.json")

	// Convert the map back to JSON structure
	var blogs struct {
		Blogs []BlogMetadata `json:"blogs"`
	}

	for _, b := range blogMetadata {
		blogs.Blogs = append(blogs.Blogs, b)
	}

	// Marshal to JSON with indentation
	jsonData, err := json.MarshalIndent(blogs, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	// Overwrite the file with the new content
	err = os.WriteFile(metadataPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing blog metadata:", err)
		return err
	}

	return nil
}

func writeToFile(filepath string, data []byte) error {
	// os.WriteFile automatically creates or overwrites the file
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
