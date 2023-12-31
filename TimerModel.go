package main

import (
	"encoding/xml"
)

type TimerPackage struct {
	XMLName    xml.Name     `xml:"TimerPackage" json:"-"`
	Timers     []Timer      `xml:"Timer" json:"-"`
	TimerGroup []TimerGroup `xml:"TimerGroup" json:"-"`
}

type Timer struct {
	IsActive     string       `xml:"isActive,attr" json:"isActive"`
	IsFolder     string       `xml:"isFolder,attr" json:"-"`
	Name         string       `xml:"name" json:"name"`
	Time         string       `xml:"time" json:"-"`
	Command      string       `xml:"command" json:"command"`
	Hours        string       `xml:"-" json:"hours"`
	Minutes      string       `xml:"-" json:"minutes"`
	Seconds      string       `xml:"-" json:"seconds"`
	Milliseconds string       `xml:"-" json:"milliseconds"`
	Script       string       `xml:"script" json:"script"`
	Timers       []Timer      `xml:"Timer" json:"-"`
	TimerGroup   []TimerGroup `xml:"TimerGroup" json:"-"`
}

type TimerGroup struct {
	Timer
}
