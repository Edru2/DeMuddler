package main

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func writeTimerList(timers *[]TimerEntity, parentDir string, seenNames map[string]bool) {
	for i := range *timers {
		timer := &(*timers)[i]
		if timer.IsFolder == "no" {
			handleTimers(&timer.Timers, parentDir, seenNames)
		}
		writeScriptFiles(timer, parentDir, seenNames)
	}
}

func writeTimerJson(timers *[]TimerEntity, parentDir string) {

	for i := range *timers {
		timer := &(*timers)[i]
		ConvertTimeToSingleProperties(timer)
	}
	writeJsonToFilewriteJsonToFile(timers, parentDir, "timers.json")
}

func handleTimers(timers *[]TimerEntity, parentDir string, seenNames map[string]bool) {
	if len(*timers) == 0 {
		return
	}
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	writeTimerList(timers, parentDir, seenNames)
	writeTimerJson(timers, parentDir)
}

func handleTimerGroups(groups *[]TimerEntity, baseDir string, seenNames map[string]bool) {
	for i := range *groups {
		group := &(*groups)[i]

		groupPath := filepath.Join(baseDir, group.Name)
		if err := os.MkdirAll(groupPath, 0755); err != nil {
			panic(err)
		}

		handleTimers(&group.Timers, groupPath, seenNames)
		handleTimerGroups(&group.TimerGroup, groupPath, seenNames)
		group.Timers = nil
		var singleGroup []TimerEntity
		singleGroup = append(singleGroup, *group)
		writeTimerList(&singleGroup, baseDir, seenNames)
		writeTimerJson(&singleGroup, baseDir)
	}
}
func ConvertTimeToSingleProperties(timer *TimerEntity) TimerEntity {
	t, err := time.Parse("15:04:05.000", timer.Time)
	if err != nil {
		panic(err)
	}
	timer.Hours = strconv.Itoa(t.Hour())
	timer.Minutes = strconv.Itoa(t.Minute())
	timer.Seconds = strconv.Itoa(t.Second())
	timer.Milliseconds = strconv.Itoa(t.Nanosecond() / 1000000)

	return *timer
}
