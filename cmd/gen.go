package main

import (
	"github.com/redish101/reblog/internal/db"
	"github.com/redish101/reblog/internal/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db.DB())

	g.ApplyBasic(
		model.Site{},
		model.Article{},
		model.User{},
		model.Friend{},
	)

	g.Execute()
}
