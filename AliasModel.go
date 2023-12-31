package main

import (
	"encoding/xml"
)

type AliasPackage struct {
	XMLName    xml.Name     `xml:"AliasPackage" json:"-"`
	AliasList  []Alias      `xml:"Alias" json:"-"`
	AliasGroup []AliasGroup `xml:"AliasGroup" json:"-"`
}

type Alias struct {
	IsActive   string       `xml:"isActive,attr" json:"isActive"`
	IsFolder   string       `xml:"isFolder,attr" json:"isFolder"`
	Name       string       `xml:"name" json:"name"`
	Regex      string       `xml:"regex" json:"regex,omitempty"`
	Script     string       `xml:"script" json:"script"`
	AliasList  []Alias      `xml:"Alias" json:"-"`
	AliasGroup []AliasGroup `xml:"AliasGroup" json:"-"`
}

type AliasGroup struct {
	Alias
}
