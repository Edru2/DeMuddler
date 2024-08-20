package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFlags() string {
	filePath := flag.String("f", "", "Path to the file")
	flag.Parse()
	return *filePath
}

func validateFilePath(filePath string) {
	if filePath == "" {
		fmt.Println("You must specify a file path using the -f flag.")
		os.Exit(1)
	}
}

func openZipFile(filePath string) *zip.ReadCloser {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		panic(err)
	}
	return r
}

func getWorkingDirectory() string {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return workDir
}

func extractFileName(filePath string) string {
	fileNameWithExtension := filepath.Base(filePath)
	return strings.TrimSuffix(fileNameWithExtension, filepath.Ext(fileNameWithExtension))
}

func processZipFiles(r *zip.ReadCloser, baseDir, workDir, fileName string) {
	for _, f := range r.File {
		switch {
		case strings.HasSuffix(strings.ToLower(f.Name), ".xml"):
			handleXML(f, baseDir)
		case strings.EqualFold(f.Name, "config.lua"):
			handleConfig(f, filepath.Join(workDir, fileName))
		default:
			handleResources(f, filepath.Join(baseDir, "resources"))
		}
	}
}

func containsIllegalCharacters(filename string) bool {
	illegalCharacters := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}
	result := strings.ContainsAny(filename, string(illegalCharacters))
	if result {
		fmt.Printf("warn: file %s contains illegal characters, consider changing its name.\n", filename)
	}
	return result
}
