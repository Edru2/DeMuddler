package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func writeKeyList(keys *[]KeyEntity, parentDir string, seenNames map[string]bool) {
	for i := range *keys {
		key := &(*keys)[i]
		if key.IsFolder == "no" {
			handleKeys(&key.KeyList, parentDir, seenNames)
		}
		writeScriptFiles(key, parentDir, seenNames)
	}
}

func writeKeyJson(keys *[]KeyEntity, parentDir string) {

	for i := range *keys {
		key := &(*keys)[i]
		convertKeyCodes(key)
	}
	writeJsonToFile(keys, parentDir, "keys.json")
}

func handleKeys(keys *[]KeyEntity, parentDir string, seenNames map[string]bool) {
	if len(*keys) == 0 {
		return
	}
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	writeKeyList(keys, parentDir, seenNames)
	writeKeyJson(keys, parentDir)
}

func handleKeyGroups(groups *[]KeyEntity, baseDir string, seenNames map[string]bool) {
	for i := range *groups {
		group := &(*groups)[i]

		groupPath := filepath.Join(baseDir, group.Name)
		if err := os.MkdirAll(groupPath, 0755); err != nil {
			panic(err)
		}

		handleKeys(&group.KeyList, groupPath, seenNames)
		handleKeyGroups(&group.KeyGroup, groupPath, seenNames)
		group.KeyList = nil
		var singleGroup []KeyEntity
		singleGroup = append(singleGroup, *group)
		writeKeyList(&singleGroup, baseDir, seenNames)
		writeKeyJson(&singleGroup, baseDir)
	}
}

func convertKeyCodes(key *KeyEntity) {
	var pressedKeys []string
	for _, mod := range keyModifiers {
		value, err := strconv.ParseInt(key.KeyModifier, 10, 32)
		if err != nil {
			fmt.Println("Error parsing KeyModifier:", err)
			continue
		}
		if int(value)&mod.Bitmask != 0 {
			pressedKeys = append(pressedKeys, mod.Name)
		}
	}
	value, err := strconv.ParseInt(key.KeyCode, 10, 32)
	if err != nil {
		fmt.Println("Error parsing KeyCode:", err)
	}
	val, ok := keyMap[int(value)]
	if ok {
		pressedKeys = append(pressedKeys, val)
	}
	key.Keys = strings.Join(pressedKeys, "+")
}
