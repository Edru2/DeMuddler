package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func handleTimers(timers *[]Timer, parentDir string) {
	if len(*timers) == 0 {
		return
	}
	var jsonFile []Timer
	if parentDir != "" {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			panic(err)
		}
	}

	for _, timer := range *timers {
		timerFileName := strings.ReplaceAll(timer.Name, " ", "_")
		timerFilePath := filepath.Join(parentDir, timerFileName+".lua")
		if len(timer.Script) > 0 && !containsIllegalCharacters(timerFileName) {
			if err := os.WriteFile(timerFilePath, []byte(timer.Script), 0644); err != nil {
				panic(err)
			}
			timer.Script = ""
		}
		ConvertTimeToSingleProperties(&timer)
		jsonFile = append(jsonFile, timer)
	}

	jsonFilePath := filepath.Join(parentDir, "timers.json")
	jsonData, err := json.MarshalIndent(jsonFile, "", "       ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func handleTimerGroups(groups *[]TimerGroup, baseDir string) {
	for i := range *groups {
		groupPath := filepath.Join(baseDir, (*groups)[i].Name)
		handleTimers(&((*groups)[i].Timers), groupPath)
		handleTimerGroups(&((*groups)[i].TimerGroup), groupPath)
	}
}
func ConvertTimeToSingleProperties(timer *Timer) Timer {
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
