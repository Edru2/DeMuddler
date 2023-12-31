package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func containsIllegalCharacters(filename string) bool {
	illegalCharacters := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}
	return strings.ContainsAny(filename, string(illegalCharacters))
}

func main() {
	filePath := flag.String("f", "", "Path to the file")
	flag.Parse()
	if *filePath == "" {
		fmt.Println("You must specify a file path using the -f flag.")
		os.Exit(1)
	}
	r, err := zip.OpenReader(*filePath)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileNameWithExtension := filepath.Base(*filePath)
	fileName := strings.TrimSuffix(fileNameWithExtension, filepath.Ext(fileNameWithExtension))

	baseDir := filepath.Join(workDir, fileName, "src")
	for _, f := range r.File {
		if strings.HasSuffix(strings.ToLower(f.Name), ".xml") {
			handleXML(f, baseDir)
			continue
		}
		if strings.EqualFold(f.Name, "config.lua") {
			handleConfig(f, filepath.Join(workDir, fileName))
			continue
		}
		handleResources(f, filepath.Join(baseDir, "resources"))
	}

}
