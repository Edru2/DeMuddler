package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func handleKeys(keys *[]Key, parentDir string) {
	if len(*keys) == 0 {
		return
	}
	var jsonFile []Key
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	for _, key := range *keys {
		keyFileName := strings.ReplaceAll(key.Name, " ", "_")
		keyFilePath := filepath.Join(parentDir, keyFileName+".lua")
		err := os.WriteFile(keyFilePath, []byte(key.Script), 0644)
		key.Script = ""
		key = convertKeyCodes(key)
		if err != nil {
			panic(err)
		}
		jsonFile = append(jsonFile, key)
	}

	jsonFilePath := filepath.Join(parentDir, "keys.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func handleKeyGroups(groups *[]KeyGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		handleKeys(&((*groups)[i].KeyList), groupPath)
		handleKeyGroups(&((*groups)[i].KeyGroup), groupPath)
	}
}

func convertKeyCodes(key Key) Key {
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

	return key
}
