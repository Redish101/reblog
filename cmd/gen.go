package main

import (
	"github.com/ChuqiCloud/acmeidc/internal/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.ApplyBasic(
		model.User{},
	)

	g.Execute()
}
