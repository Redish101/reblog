package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var dist embed.FS

func GetUIFS() fs.FS {
	return dist
}
