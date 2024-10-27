package feed

import (
	"fmt"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"

	"github.com/gorilla/feeds"
)

func cookContent(link string,content string) string {
	prefix := fmt.Sprintf("> 本文该渲染由 reblog 生成，可能存在排版问题，最佳体验请前往：\n[%s](%s)\n\n", link, link)
	return prefix + content
}

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
		link := fmt.Sprintf(app.Config().Rss.ContentLinkFormat, site.Url, article.Slug)

		feed.Items = append(feed.Items, &feeds.Item{
			Title:       article.Title,
			Description: article.Desc,
			Content:     markdownService.Render(cookContent(link, article.Content)),
			Link:        &feeds.Link{Href: link},
			Created:     article.CreatedAt,
		})
	}

	rssString, err := feed.ToAtom()

	if err != nil {
		return "", err
	}

	return rssString, nil
}
