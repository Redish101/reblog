package markdown

import "github.com/88250/lute"

type Renderer struct {
	luteEngine *lute.Lute
}

func NewRenderer() *Renderer {
	luteEngine := lute.New()

	return &Renderer{luteEngine: luteEngine}
}

func (r *Renderer) Render(content string) string {
	return r.luteEngine.MarkdownStr("", content)
}
