package main

import (
	"os"
	"path/filepath"
)

func writeScripts(scripts *[]ScriptEntity, parentDir string, seenNames map[string]bool) {
	for i := range *scripts {
		script := &(*scripts)[i]
		if script.IsFolder == "no" {
			handleScripts(&script.Scripts, parentDir, seenNames)
		}
		writeScriptFiles(script, parentDir, seenNames)
	}
}

func writeScriptJson(scripts *[]ScriptEntity, parentDir string) {
	writeJsonToFilewriteJsonToFile(scripts, parentDir, "scripts.json")
}

func handleScripts(scripts *[]ScriptEntity, parentDir string, seenNames map[string]bool) {
	if len(*scripts) == 0 {
		return
	}
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	writeScripts(scripts, parentDir, seenNames)
	writeScriptJson(scripts, parentDir)
}

func handleScriptGroups(groups *[]ScriptEntity, baseDir string, seenNames map[string]bool) {
	for i := range *groups {
		group := &(*groups)[i]

		groupPath := filepath.Join(baseDir, group.Name)
		if err := os.MkdirAll(groupPath, 0755); err != nil {
			panic(err)
		}

		handleScripts(&group.Scripts, groupPath, seenNames)
		handleScriptGroups(&group.ScriptGroup, groupPath, seenNames)
		group.Scripts = nil
		var singleGroup []ScriptEntity
		singleGroup = append(singleGroup, *group)
		writeScripts(&singleGroup, baseDir, seenNames)
		writeScriptJson(&singleGroup, baseDir)
	}
}
