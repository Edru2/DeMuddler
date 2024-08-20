package main

import (
	"encoding/xml"
)

type ScriptPackage struct {
	XMLName     xml.Name       `xml:"ScriptPackage" json:"-"`
	Scripts     []ScriptEntity `xml:"Script" json:"-"`
	ScriptGroup []ScriptEntity `xml:"ScriptGroup" json:"-"`
}

type ScriptEntity struct {
	Name             string         `xml:"name" json:"name"`
	IsActive         string         `xml:"isActive,attr" json:"isActive"`
	EventHandlerList []string       `xml:"eventHandlerList>string" json:"eventHandlerList,omitempty"`
	Script           string         `xml:"script" json:"script"`
	Scripts          []ScriptEntity `xml:"Script" json:"children,omitempty"`
	ScriptGroup      []ScriptEntity `xml:"ScriptGroup" json:"-"`
	IsFolder         string         `xml:"isFolder,attr" json:"isFolder"`
}

func (s *ScriptEntity) GetName() string {
	return s.Name
}

func (s *ScriptEntity) GetScript() string {
	return s.Script
}

func (s *ScriptEntity) SetScript(script string) {
	s.Script = script
}
