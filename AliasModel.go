package main

import (
	"encoding/xml"
)

type AliasPackage struct {
	XMLName    xml.Name      `xml:"AliasPackage" json:"-"`
	AliasList  []AliasEntity `xml:"Alias" json:"-"`
	AliasGroup []AliasEntity `xml:"AliasGroup" json:"-"`
}

type AliasEntity struct {
	Name       string        `xml:"name" json:"name"`
	IsActive   string        `xml:"isActive,attr" json:"isActive"`
	Command    string        `xml:"command" json:"command,omitempty"`
	Regex      string        `xml:"regex" json:"regex,omitempty"`
	Script     string        `xml:"script" json:"script"`
	AliasList  []AliasEntity `xml:"Alias" json:"children,omitempty"`
	IsFolder   string        `xml:"isFolder,attr" json:"isFolder"`
	AliasGroup []AliasEntity `xml:"AliasGroup" json:"-"`
}

func (a *AliasEntity) GetName() string {
	return a.Name
}

func (a *AliasEntity) GetScript() string {
	return a.Script
}

func (a *AliasEntity) SetScript(script string) {
	a.Script = script
}
