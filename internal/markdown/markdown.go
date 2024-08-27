package markdown

import "gitlab.com/golang-commonmark/markdown"

type Renderer struct {
	markdown *markdown.Markdown
}

func NewRenderer() *Renderer {
	md := markdown.New(markdown.XHTMLOutput(true))

	return &Renderer{markdown: md}
}

func (r *Renderer) Render(content string) string {
	return r.markdown.RenderToString([]byte(content))
}
