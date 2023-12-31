package main

import (
	"encoding/xml"
)

type ScriptPackage struct {
	XMLName     xml.Name      `xml:"ScriptPackage" json:"-"`
	Scripts     []Script      `xml:"Script" json:"-"`
	ScriptGroup []ScriptGroup `xml:"ScriptGroup" json:"-"`
}

type Script struct {
	IsActive         string        `xml:"isActive,attr" json:"isActive"`
	IsFolder         string        `xml:"isFolder,attr" json:"isFolder"`
	Name             string        `xml:"name" json:"name"`
	EventHandlerList []string      `xml:"eventHandlerList>string" json:"eventHandlerList,omitempty"`
	Script           string        `xml:"script" json:"script"`
	Scripts          []Script      `xml:"Script" json:"-"`
	ScriptGroup      []ScriptGroup `xml:"ScriptGroup" json:"-"`
}

type ScriptGroup struct {
	Script
}
