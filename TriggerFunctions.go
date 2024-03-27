package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func convertToJSONTrigger(parentDir string, trigger Trigger) JSONTrigger {
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
		HighlightFG:    trigger.MFgColor,
		HighlightBG:    trigger.MBgColor,
	}

	for i, content := range trigger.RegexCodeList {
		jsonPattern := Pattern{
			Pattern: content,
			Type:    patternNumberToType(trigger.RegexCodeProperties[i]),
		}

		if jsonPattern.Type == "colour" {
			jsonPattern.Pattern = extractColorTriggerColors(content)
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
			jsonTrigger.Children = append(jsonTrigger.Children, convertToJSONTrigger(parentDir, child))
		}
	}
	writeTriggerScript(parentDir, &trigger, &jsonTrigger)
	return jsonTrigger
}

func convertToAnsiColor(colorStr string) (ansiColor string) {
	switch colorStr {
	case "-2":
		ansiColor = "IGNORE" 
	case "0":
		ansiColor = "DEFAULT" 
	case "1":
		ansiColor = "8" // Light black (dark gray)
	case "2":
		ansiColor = "0" // Black
	case "3":
		ansiColor = "9" // Light red
	case "4":
		ansiColor = "1" // Red
	case "5":
		ansiColor = "10" // Light green
	case "6":
		ansiColor = "2" // Green
	case "7":
		ansiColor = "11" // Light yellow
	case "8":
		ansiColor = "3" // Yellow
	case "9":
		ansiColor = "12" // Light blue
	case "10":
		ansiColor = "4" // Blue
	case "11":
		ansiColor = "13" // Light magenta
	case "12":
		ansiColor = "5" // Magenta
	case "13":
		ansiColor = "14" // Light cyan
	case "14":
		ansiColor = "6" // Cyan
	case "15":
		ansiColor = "15" // Light white
	case "16":
		ansiColor = "7" // White (light gray)
	default:
		ansiColor = colorStr // Use color directly if it doesn't match any case
	}
	return ansiColor
}

func extractColorTriggerColors(colorStr string) string {
	var fgColor, bgColor string
	fgIndex := strings.Index(colorStr, "FG")
	bgIndex := strings.Index(colorStr, "BG")

	if fgIndex != -1 && bgIndex != -1 {
		fgColor = colorStr[fgIndex+2 : bgIndex]
		bgColor = colorStr[bgIndex+2:]
	}
	return fmt.Sprintf("%s,%s", convertToAnsiColor(fgColor), convertToAnsiColor(bgColor))
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
		jsonTrigger := convertToJSONTrigger(parentDir, trigger)
		jsonFile = append(jsonFile, jsonTrigger)
	}

	return jsonFile

}

func writeTriggerScript(parentDir string, trigger *Trigger, jsonTrigger *JSONTrigger) {
	scriptFileName := strings.ReplaceAll(trigger.Name, " ", "_")
	scriptFilePath := filepath.Join(parentDir, scriptFileName+".lua")
	if len(trigger.Script) == 0 || containsIllegalCharacters(scriptFileName) {
		jsonTrigger.Script = trigger.Script
	} else {
		if err := os.WriteFile(scriptFilePath, []byte(trigger.Script), 0644); err != nil {
			panic(err)
		}
	}
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
