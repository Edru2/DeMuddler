package main

import (
	"os"
	"path/filepath"
)

func writeAliasList(aliasList *[]AliasEntity, parentDir string, seenNames map[string]bool) {
	for i := range *aliasList {
		alias := &(*aliasList)[i]
		if alias.IsFolder == "no" {
			handleAliases(&alias.AliasList, parentDir, seenNames)
		}
		writeScriptFiles(alias, parentDir, seenNames)
	}
}

func writeAliasJson(aliasList *[]AliasEntity, parentDir string) {
	writeJsonToFile(aliasList, parentDir, "aliases.json")
}

func handleAliases(aliasList *[]AliasEntity, parentDir string, seenNames map[string]bool) {
	if len(*aliasList) == 0 {
		return
	}
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	writeAliasList(aliasList, parentDir, seenNames)
	writeAliasJson(aliasList, parentDir)
}
func handleAliasGroups(groups *[]AliasEntity, baseDir string, seenNames map[string]bool) {
	for i := range *groups {
		group := &(*groups)[i]

		groupPath := filepath.Join(baseDir, group.Name)
		if err := os.MkdirAll(groupPath, 0755); err != nil {
			panic(err)
		}

		handleAliases(&group.AliasList, groupPath, seenNames)
		handleAliasGroups(&group.AliasGroup, groupPath, seenNames)
		group.AliasList = nil
		var singleGroup []AliasEntity
		singleGroup = append(singleGroup, *group)
		writeAliasList(&singleGroup, baseDir, seenNames)
		writeAliasJson(&singleGroup, baseDir)
	}

}
