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

	// only append the children to this triggers.json if they are *not* in a folder
	// this seems weird, but the child triggers / trigger groups will be handled by recursion
	// in handleTriggerGroups.
	// Adding children elements associated with folders will prevent muddler from picking up the correct
	// json and scripts inside those folders.
	if trigger.IsFolder == "no" {
		for _, child := range trigger.Triggers {
			jsonTrigger.Children = append(jsonTrigger.Children, convertToJSONTrigger(child))
		}
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

func handleTriggers(triggers *[]Trigger, parentDir string) []JSONTrigger {
	if len(*triggers) == 0 {
		return nil
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

	return jsonFile

}

func writeJson(jsonFile []JSONTrigger, parentDir string) {
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	jsonFilePath := filepath.Join(parentDir, "triggers.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func handleTriggerGroups(groups *[]TriggerGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		// first, handle all the triggers in this group, but don't write triggers.json yet
		jsonTriggers := handleTriggers(&((*groups)[i].Triggers), groupPath)

		// next, for all triggerGroups in this group, if they are a chain, add them to triggers.json
		triggerGroups := (*groups)[i].TriggerGroup // all the triggerGroups in this group
		for x := range triggerGroups {
			if len(triggerGroups[x].Trigger.RegexCodeList) > 0 { // chains require a trigger pattern
				localJson := handleTriggers(&[]Trigger{triggerGroups[x].Trigger}, groupPath) // this group's trigger data
				jsonTriggers = append(jsonTriggers, localJson...)
			}
		}

		// now write the full triggers.json (trigger data + triggerGroup data)
		writeJson(jsonTriggers, groupPath)
		// recur
		handleTriggerGroups(&((*groups)[i].TriggerGroup), groupPath)
	}
}

// handle the root directory first, before using recursion to build the rest of the TriggerGroups.
func handleTriggerPackage(groups *[]TriggerGroup, triggers *[]Trigger, baseDir string) {
	jsonTriggers := handleTriggers(triggers, baseDir)

	for i := range *groups {
		if len((*groups)[i].Trigger.RegexCodeList) > 0 { // chains require a trigger pattern
			localJson := handleTriggers(&[]Trigger{(*groups)[i].Trigger}, baseDir) // the Trigger data on TriggerGroup
			jsonTriggers = append(jsonTriggers, localJson...)
		}
	}

	writeJson(jsonTriggers, baseDir)

	handleTriggerGroups(groups, baseDir)
}
