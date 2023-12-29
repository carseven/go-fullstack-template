package handler

import (
	"github.com/carseven/go-fullstack-template/view/cases/blog"
	layout "github.com/carseven/go-fullstack-template/view/layouts"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v4"
)

type MdHandler struct{}

var mds = `# Title
## Title 2
Example of text

- List 1
- List 2

[link](www.google.com)
`

func (h MdHandler) HandleMd(c echo.Context) error {
	blogHtml := mdToHTML([]byte(mds))
	blogView := blog.View(string(blogHtml))
	return render(c, layout.Base(blogView))
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
