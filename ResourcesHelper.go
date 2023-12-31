package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func handleResources(f *zip.File, destDir string) {
	destPath := filepath.Join(destDir, f.Name)

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
			panic(err)
		}

		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		outFile, err := os.Create(destPath)
		if err != nil {
			panic(err)
		}

		defer outFile.Close()
		if _, err = io.Copy(outFile, rc); err != nil {
			panic(err)
		}
		if strings.HasPrefix(f.Name, ".mudlet/Icon/") {
			content, err := io.ReadAll(rc)
			if err != nil {
				panic(err)
			}
			fileName := strings.TrimPrefix(filepath.FromSlash(f.Name), ".mudlet/Icon/")
			os.WriteFile(filepath.Join(destDir, fileName), content, 0644)

		}
	}
}
