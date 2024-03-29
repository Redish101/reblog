package main

import (
	"reblog/internal/db"
	"reblog/internal/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db.DB())

	g.ApplyBasic(model.Site{}, model.Article{}, model.User{})

	g.Execute()
}
