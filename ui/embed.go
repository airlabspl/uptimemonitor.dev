package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var distFS embed.FS

func FS() fs.FS {
	subDir, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(err)
	}

	return subDir
}
