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
				handleScriptGroups(&scriptPackage.ScriptGroup, localBaseDir)
				handleScripts(&scriptPackage.Scripts, localBaseDir)
			case "AliasPackage":
				err = decoder.DecodeElement(&aliasPackage, &se)
				if err != nil {
					panic(err)
				}
				localBaseDir := filepath.Join(baseDir, "aliases")
				handleAliasGroups(&aliasPackage.AliasGroup, localBaseDir)
				handleAliases(&aliasPackage.AliasList, localBaseDir)
			case "TimerPackage":
				err = decoder.DecodeElement(&timerPackage, &se)
				if err != nil {
					panic(err)
				}
				localBaseDir := filepath.Join(baseDir, "timers")
				handleTimerGroups(&timerPackage.TimerGroup, localBaseDir)
				handleTimers(&timerPackage.Timers, localBaseDir)
			case "KeyPackage":
				err = decoder.DecodeElement(&keyPackage, &se)
				if err != nil {
					panic(err)
				}
				localBaseDir := filepath.Join(baseDir, "keys")
				handleKeyGroups(&keyPackage.KeyGroup, localBaseDir)
				handleKeys(&keyPackage.KeyList, localBaseDir)
			}
		}
	}
}
