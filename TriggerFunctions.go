package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func convertToJSONTrigger(trigger Trigger) JSONTrigger {
	jsonTrigger := JSONTrigger{
		Name:           trigger.Name,
		IsActive:       trigger.IsActive,
		IsFolder:       trigger.IsFolder,
		Command:        trigger.MCommand,
		Multiline:      trigger.IsMultiline,
		MultilineDelta: strconv.Itoa(trigger.ConditionLineDelta),
		Matchall:       trigger.IsPerlSlashGOption,
		Filter:         trigger.IsFilterTrigger,
		FireLength:     trigger.MStayOpen,
		SoundFile:      trigger.MSoundFile,
		Highlight:      trigger.IsColorizerTrigger,
		HighlightFG:    trigger.IsColorTriggerFg,
		HighlightBG:    trigger.IsColorTriggerBg,
	}

	for i, content := range trigger.RegexCodeList {
		jsonPattern := Pattern{
			Pattern: content,
			Type:    patternNumberToType(trigger.RegexCodeProperties[i]),
		}
		jsonTrigger.Patterns = append(jsonTrigger.Patterns, jsonPattern)
	}

	for _, child := range trigger.Triggers {
		jsonTrigger.Children = append(jsonTrigger.Children, convertToJSONTrigger(child))
	}

	return jsonTrigger
}

func patternNumberToType(patternNumber int) string {
	typeMap := map[int]string{
		0: "substring",
		1: "regex",
		2: "startOfLine",
		3: "exactMatch",
		4: "lua",
		5: "spacer",
		6: "colour",
		7: "prompt",
	}

	triggerType, exists := typeMap[patternNumber]
	if !exists {
		return "substring" // Default value
	}
	return triggerType
}

func handleTriggers(triggers *[]Trigger, parentDir string) {
	if len(*triggers) == 0 {
		return
	}
	var jsonFile []JSONTrigger
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	for _, trigger := range *triggers {
		scriptFileName := strings.ReplaceAll(trigger.Name, " ", "_")
		scriptFilePath := filepath.Join(parentDir, scriptFileName+".lua")
		jsonTrigger := convertToJSONTrigger(trigger)
		if len(trigger.Script) == 0 || containsIllegalCharacters(scriptFileName) {
			jsonTrigger.Script = trigger.Script
		} else {
			if err := os.WriteFile(scriptFilePath, []byte(trigger.Script), 0644); err != nil {
				panic(err)
			}
		}
		jsonFile = append(jsonFile, jsonTrigger)
	}

	jsonFilePath := filepath.Join(parentDir, "triggers.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func handleTriggerGroups(groups *[]TriggerGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		handleTriggers(&((*groups)[i].Triggers), groupPath)
		handleTriggerGroups(&((*groups)[i].TriggerGroup), groupPath)
	}
}
