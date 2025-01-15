package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type JSONSerializable interface {
	TimerEntity | KeyEntity | ScriptEntity | AliasEntity
}

type ScriptHandler interface {
	GetName() string
	GetScript() string
	SetScript(script string)
}

func writeJsonToFilewriteJsonToFile[T JSONSerializable](data *[]T, parentDir, fileName string) {
	if len(*data) == 0 {
		return
	}

	jsonFilePath := filepath.Join(parentDir, fileName)

	// Read previous data, if any
	var prev []T
	prevData, err := os.ReadFile(jsonFilePath)
	if err == nil {
		json.Unmarshal(prevData, &prev)
	}

	finalData := append(prev, *data...)
	jsonData, err := json.MarshalIndent(finalData, "", "       ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func writeScriptFiles[T ScriptHandler](script T, parentDir string, seenNames map[string]bool) {
	scriptFileName := strings.ReplaceAll(script.GetName(), " ", "_")
	scriptFilePath := filepath.Join(parentDir, scriptFileName+".lua")
	if len(script.GetScript()) > 0 && !containsIllegalCharacters(scriptFileName) {
		if !seenNames[scriptFilePath] {
			if err := os.WriteFile(scriptFilePath, []byte(script.GetScript()), 0644); err != nil {
				panic(err)
			}
			// Clear script content after writing to avoid it reappear in the json
			script.SetScript("")
			seenNames[scriptFilePath] = true
		}
	}

}
