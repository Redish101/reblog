package rss

import (
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/markdown"
	"github.com/redish101/reblog/internal/model"

	"github.com/gorilla/feeds"
)

func GenerateRSS(app *core.App, articles []*model.Article) (string, error) {
	s := app.Query().Site

	site, err := s.First()

	if err != nil {
		return "", err
	}

	feed := feeds.Feed{
		Title:       site.Name,
		Description: site.Desc,
		Link:        &feeds.Link{Href: site.Url},
	}

	for _, article := range articles {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       article.Title,
			Description: article.Desc,
			Content:     markdown.MarkdownToHtml(article.Content),
			Created:     article.CreatedAt,
		})
	}

	rssString, err := feed.ToAtom()

	if err != nil {
		return "", err
	}

	return rssString, nil
}
