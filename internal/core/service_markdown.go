package core

import "github.com/redish101/reblog/internal/markdown"

type MarkdownService struct {
	app *App

	renderer *markdown.Renderer
	cache    map[string]string
}

func NewMarkdownService(app *App) *MarkdownService {
	return &MarkdownService{app: app}
}

func (s *MarkdownService) Start() error {
	s.renderer = markdown.NewRenderer()
	s.cache = make(map[string]string)

	return nil
}

func (s *MarkdownService) Stop() error {
	s.renderer = nil

	return nil
}

func (s *MarkdownService) Render(content string) string {
	if cachedResult, isFound := s.cache[content]; isFound {
		return cachedResult
	}

	result := s.renderer.Render(content)
	s.cache[content] = result

	return result
}
