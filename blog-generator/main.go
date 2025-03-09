package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"path/filepath"
)

const pathToRoot string = "../"
const pathToBlogs string = pathToRoot + "blogs"

var mds = `# header

Sample text.

[link](http://example.com)
`

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	ast := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(ast, renderer)
}

func main() {
	err := filepath.Walk(pathToBlogs, processMarkdown)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Date(2022, 6, 4, 0, 0, 0, 0, time.UTC).Format("2006-01-02"))

}

func processMarkdown(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return err
	}

	// generate Metadata and html
	md := []byte(contents)
	html := mdToHTML(md)
	fmt.Println(string(html))

	// append Metadata to project-root/blog-generator/blogMetadata.json

	// write HTML to file in project-root/frontend/src/views/blog/blog-html/<<slug>>.html

	return nil
}
