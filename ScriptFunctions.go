package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func handleScripts(scripts *[]Script, parentDir string) {
	if len(*scripts) == 0 {
		return
	}
	var jsonFile []Script
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	for _, script := range *scripts {
		scriptFileName := strings.ReplaceAll(script.Name, " ", "_")
		scriptFilePath := filepath.Join(parentDir, scriptFileName+".lua")
		if len(script.Script) > 0 && !containsIllegalCharacters(scriptFileName) {
			if err := os.WriteFile(scriptFilePath, []byte(script.Script), 0644); err != nil {
				panic(err)
			}
			script.Script = ""
		}
		jsonFile = append(jsonFile, script)
	}

	jsonFilePath := filepath.Join(parentDir, "scripts.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func handleScriptGroups(groups *[]ScriptGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		handleScripts(&((*groups)[i].Scripts), groupPath)
		handleScriptGroups(&((*groups)[i].ScriptGroup), groupPath)
	}
}
