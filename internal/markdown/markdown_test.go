package markdown

import (
	_ "embed"
	"testing"
)

//go:embed markdown_spec.md
var content string

func BenchmarkMarkdown(b *testing.B) {
	renderer := NewRenderer()

	for i := 0; i < b.N; i++ {
		renderer.Render(content)
	}
}
