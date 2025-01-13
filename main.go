package main

import (
	"path/filepath"
)

func main() {
	filePath := parseFlags()
	validateFilePath(filePath)
	r := openZipFile(filePath)
	defer r.Close()

	workDir := getWorkingDirectory()
	fileName := extractFileName(filePath)

	baseDir := filepath.Join(workDir, fileName, "src")
	processZipFiles(r, baseDir, workDir, fileName)
}
