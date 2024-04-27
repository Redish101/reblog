package rss

import (
	"reblog/internal/model"
	"reblog/internal/query"

	"github.com/gorilla/feeds"
)

func GenerateRSS(articles []*model.Article) (string, error) {
	s := query.Site

	site, err := s.First()

	if err != nil {
		return "", err
	}
	
	feed := feeds.Feed {
		Title: site.Name,
		Description: site.Desc,
		Link: &feeds.Link{Href: site.Url},
	}

	for _, article := range articles {
		feed.Items = append(feed.Items, &feeds.Item{
			Title: article.Title,
			Description: article.Desc,
			Content: article.Content,
			Created: article.CreatedAt,
		})
	}

	rssString, err := feed.ToRss()

	if err != nil {
		return "", err
	}

	return rssString, nil
}
