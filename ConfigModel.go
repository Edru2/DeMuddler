package main

type PackageConfig struct {
	Package      string `json:"package"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Version      string `json:"version"`
	Author       string `json:"author"`
	Icon         string `json:"icon"`
	Dependencies string `json:"dependencies"`
	OutputFile   bool   `json:"outputFile"`
}
