package main

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func handleConfig(f *zip.File, baseDir string) {
	rc, err := f.Open()
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(rc)
	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(err)
	}
	config := parseLuaConfig(string(content))
	description := config.Description
	config.Description = ""
	output, _ := json.MarshalIndent(config, "", "       ")

	configFilePath := filepath.Join(baseDir, "mfile")
	err = os.WriteFile(configFilePath, output, 0644)
	if err != nil {
		panic(err)
	}

	readmeFilePath := filepath.Join(baseDir, "README.md")
	err = os.WriteFile(readmeFilePath, []byte(description), 0644)
	if err != nil {
		panic(err)
	}
}

func parseLuaConfig(content string) PackageConfig {
	var config PackageConfig
	var captureDescription bool
	var descriptionBuilder strings.Builder

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if captureDescription {
			if strings.Contains(line, "]]") {
				captureDescription = false
				descriptionBuilder.WriteString(strings.Trim(line, "[\"]"))
				config.Description = descriptionBuilder.String()
				continue
			}
			descriptionBuilder.WriteString(line + "\n")
			continue
		}

		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "mpackage":
				config.Package = strings.Trim(value, "[\"]")
			case "author":
				config.Author = strings.Trim(value, "[\"]")
			case "icon":
				config.Icon = strings.Trim(value, "[\"]")
			case "title":
				config.Title = strings.Trim(value, "[\"]")
			case "description":
				if !captureDescription && strings.Contains(value, "]]") {
					config.Description = strings.Trim(value, "[\"]")
					continue
				}
				if strings.Contains(value, "[[") {
					captureDescription = true
				}
				descriptionBuilder.WriteString(strings.Trim(value, "[\"]"))
			}
		}
	}

	return config
}
