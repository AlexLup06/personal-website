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

type BlogMetadata struct {
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Intro string   `json:"intro"`
	Slug  string   `json:"slug"`
	Tags  []string `json:"tags"`
}

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
		return nil
	}
	return buf.Bytes()
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

	if info.IsDir() || info.Name() == "blog-template.md" {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var title, date, intro string
	var tags []string
	var contentLines []string
	inContent := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "--start--" {
			inContent = true
			continue
		}

		if !inContent {
			if strings.HasPrefix(line, "Title: ") {
				title = strings.TrimPrefix(line, "Title: ")
			} else if strings.HasPrefix(line, "Date: ") {
				date = strings.TrimPrefix(line, "Date: ")
			} else if strings.HasPrefix(line, "Intro: ") {
				intro = strings.TrimPrefix(line, "Intro: ")
			} else if strings.HasPrefix(line, "Tags: ") {
				tagsString := strings.TrimPrefix(line, "Tags: ")
				tags = strings.Split(tagsString, ",")
			}
		} else {
			contentLines = append(contentLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	content := strings.Join(contentLines, "\n")
	slug := generateSlug(title, date)
	blogMetadata := BlogMetadata{
		Title: title,
		Date:  date,
		Intro: intro,
		Slug:  slug,
		Tags:  tags,
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

	var blogs struct {
		Blogs []BlogMetadata `json:"blogs"`
	}

	for _, b := range blogMetadata {
		blogs.Blogs = append(blogs.Blogs, b)
	}

	jsonData, err := json.MarshalIndent(blogs, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	err = os.WriteFile(metadataPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing blog metadata:", err)
		return err
	}

	return nil
}

func writeToFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
