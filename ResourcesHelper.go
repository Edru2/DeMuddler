package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func handleResources(f *zip.File, destDir string) {
	destPath := filepath.Join(destDir, f.Name)

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
			panic(err)
		}
		return
	}

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
}
