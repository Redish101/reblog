package feed

import (
	"fmt"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"

	"github.com/gorilla/feeds"
)

func GenerateFeed(app *core.App, articles []*model.Article) (string, error) {
	s := app.Query().Site

	site, err := s.First()

	if err != nil {
		return "", err
	}

	user, err := app.Query().User.First()

	if err != nil {
		return "", err
	}

	feed := feeds.Feed{
		Title:       site.Name,
		Description: site.Desc,
		Author:      &feeds.Author{Name: user.Nickname},
		Link:        &feeds.Link{Href: site.Url},
	}

	markdownService, err := core.AppService[*core.MarkdownService](app)

	if err != nil {
		return "", err
	}

	for _, article := range articles {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       article.Title,
			Description: article.Desc,
			Content:     markdownService.Render(article.Content),
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/article/%s", site.Url, article.Slug)},
			Created:     article.CreatedAt,
		})
	}

	rssString, err := feed.ToAtom()

	if err != nil {
		return "", err
	}

	return rssString, nil
}
