package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func handleAliases(aliasList *[]Alias, parentDir string) {
	if len(*aliasList) == 0 {
		return
	}
	var jsonFile []Alias
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	for _, alias := range *aliasList {
		aliasFileName := strings.ReplaceAll(alias.Name, " ", "_")
		aliasFilePath := filepath.Join(parentDir, aliasFileName+".lua")
		if len(alias.Script) > 0 && !containsIllegalCharacters(aliasFileName) {
			if err := os.WriteFile(aliasFilePath, []byte(alias.Script), 0644); err != nil {
				panic(err)
			}
			alias.Script = ""
		}
		jsonFile = append(jsonFile, alias)
	}

	jsonFilePath := filepath.Join(parentDir, "aliases.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func handleAliasGroups(groups *[]AliasGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		handleAliases(&((*groups)[i].AliasList), groupPath)
		handleAliasGroups(&((*groups)[i].AliasGroup), groupPath)
	}
}
