package main

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"path/filepath"
)

func handleXML(f *zip.File, baseDir string) {
	rc, err := f.Open()
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(rc)
	var triggerPackage TriggerPackage
	var scriptPackage ScriptPackage
	var aliasPackage AliasPackage
	var timerPackage TimerPackage
	var keyPackage KeyPackage

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		switch se := token.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "TriggerPackage":
				err = decoder.DecodeElement(&triggerPackage, &se)
				if err != nil {
					panic(err)
				}
				localBaseDir := filepath.Join(baseDir, "triggers")
				handleTriggerPackage(&triggerPackage.TriggerGroup, &triggerPackage.Triggers, localBaseDir)
			case "ScriptPackage":
				err = decoder.DecodeElement(&scriptPackage, &se)
				if err != nil {
					panic(err)
				}
				localBaseDir := filepath.Join(baseDir, "scripts")
				seenNames := make(map[string]bool)
				handleScriptGroups(&scriptPackage.ScriptGroup, localBaseDir, seenNames)
				handleScripts(&scriptPackage.Scripts, localBaseDir, seenNames)
			case "AliasPackage":
				err = decoder.DecodeElement(&aliasPackage, &se)
				if err != nil {
					panic(err)
				}
				seenNames := make(map[string]bool)
				localBaseDir := filepath.Join(baseDir, "aliases")
				handleAliasGroups(&aliasPackage.AliasGroup, localBaseDir, seenNames)
				handleAliases(&aliasPackage.AliasList, localBaseDir, seenNames)
			case "TimerPackage":
				err = decoder.DecodeElement(&timerPackage, &se)
				if err != nil {
					panic(err)
				}
				seenNames := make(map[string]bool)
				localBaseDir := filepath.Join(baseDir, "timers")
				handleTimerGroups(&timerPackage.TimerGroup, localBaseDir, seenNames)
				handleTimers(&timerPackage.Timers, localBaseDir, seenNames)
			case "KeyPackage":
				err = decoder.DecodeElement(&keyPackage, &se)
				if err != nil {
					panic(err)
				}
				seenNames := make(map[string]bool)
				localBaseDir := filepath.Join(baseDir, "keys")
				handleKeyGroups(&keyPackage.KeyGroup, localBaseDir, seenNames)
				handleKeys(&keyPackage.KeyList, localBaseDir, seenNames)
			}
		}
	}
}
