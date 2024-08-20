package main

import (
	"encoding/xml"
)

type TimerPackage struct {
	XMLName    xml.Name      `xml:"TimerPackage" json:"-"`
	Timers     []TimerEntity `xml:"Timer" json:"-"`
	TimerGroup []TimerEntity `xml:"TimerGroup" json:"-"`
}

type TimerEntity struct {
	Name         string        `xml:"name" json:"name"`
	IsActive     string        `xml:"isActive,attr" json:"isActive"`
	Command      string        `xml:"command" json:"command"`
	Time         string        `xml:"time" json:"-"`
	Hours        string        `xml:"-" json:"hours"`
	Minutes      string        `xml:"-" json:"minutes"`
	Seconds      string        `xml:"-" json:"seconds"`
	Milliseconds string        `xml:"-" json:"milliseconds"`
	Script       string        `xml:"script" json:"script"`
	Timers       []TimerEntity `xml:"Timer" json:"children,omitempty"`
	IsFolder     string        `xml:"isFolder,attr" json:"-"`
	TimerGroup   []TimerEntity `xml:"TimerGroup" json:"-"`
}

func (t *TimerEntity) GetName() string {
	return t.Name
}

func (t *TimerEntity) GetScript() string {
	return t.Script
}

func (t *TimerEntity) SetScript(script string) {
	t.Script = script
}
