package main

import (
	"encoding/xml"
)

type KeyPackage struct {
	XMLName  xml.Name    `xml:"KeyPackage" json:"-"`
	KeyList  []KeyEntity `xml:"Key" json:"-"`
	KeyGroup []KeyEntity `xml:"KeyGroup" json:"-"`
}

type KeyEntity struct {
	Name        string      `xml:"name" json:"name"`
	IsActive    string      `xml:"isActive,attr" json:"isActive"`
	Command     string      `xml:"command" json:"command"`
	Script      string      `xml:"script" json:"script"`
	Keys        string      `xml:"-" json:"keys"`
	KeyCode     string      `xml:"keyCode" json:"-"`
	KeyModifier string      `xml:"keyModifier" json:"-"`
	KeyList     []KeyEntity `xml:"Key" json:"children,omitempty"`
	IsFolder    string      `xml:"isFolder,attr" json:"isFolder"`
	KeyGroup    []KeyEntity `xml:"KeyGroup" json:"-"`
}

type Modifier struct {
	Bitmask int
	Name    string
}

var keyModifiers = []Modifier{
	{Bitmask: 0x02000000, Name: "shift"},
	{Bitmask: 0x04000000, Name: "ctrl"},
	{Bitmask: 0x08000000, Name: "alt"},
	{Bitmask: 0x20000000, Name: "keypad"},
}

// Global map of Qt key codes to key names
var keyMap = map[int]string{
	0x2d:       "minus",
	0x2a:       "asterisk",
	0x2f:       "slash",
	0x2b:       "plus",
	0x01000003: "backspace",
	0x01000004: "return",
	0x01000005: "enter",
	0x01000006: "ins",
	0x01000007: "del",
	0x01000009: "print",
	0x0100000b: "clear",
	0x01000010: "home",
	0x01000011: "end",
	0x01000012: "left",
	0x01000013: "up",
	0x01000014: "right",
	0x01000015: "down",
	0x01000016: "pgup",
	0x01000017: "pgdn",
	0x01000024: "caps",
	0x01000025: "num",
	0x01000026: "scroll",
	0x01000030: "F1",
	0x01000031: "F2",
	0x01000032: "F3",
	0x01000033: "F4",
	0x01000034: "F5",
	0x01000035: "F6",
	0x01000036: "F7",
	0x01000037: "F8",
	0x01000038: "F9",
	0x01000039: "F10",
	0x0100003a: "F11",
	0x0100003b: "F12",
	0x0100003c: "F13",
	0x0100003d: "F14",
	0x0100003e: "F15",
	0x0100003f: "F16",
	0x01000040: "F17",
	0x01000041: "F18",
	0x41:       "a",
	0x42:       "b",
	0x43:       "c",
	0x44:       "d",
	0x45:       "e",
	0x46:       "f",
	0x47:       "g",
	0x48:       "h",
	0x49:       "i",
	0x4a:       "j",
	0x4b:       "k",
	0x4c:       "l",
	0x4d:       "m",
	0x4e:       "n",
	0x4f:       "o",
	0x50:       "p",
	0x51:       "q",
	0x52:       "r",
	0x53:       "s",
	0x54:       "t",
	0x55:       "u",
	0x56:       "v",
	0x57:       "w",
	0x58:       "x",
	0x59:       "y",
	0x5a:       "z",
	0x30:       "0",
	0x31:       "1",
	0x32:       "2",
	0x33:       "3",
	0x34:       "4",
	0x35:       "5",
	0x36:       "6",
	0x37:       "7",
	0x38:       "8",
	0x39:       "9",
	0x20:       "space",
	0x21:       "!",
	0x22:       `"`,
	0x23:       "#",
	0x24:       "$",
	0x25:       "%",
	0x26:       "&",
	0x27:       "'",
	0x28:       "(",
	0x29:       ")",
	0x2c:       ",",
	0x2e:       ".",
	0x7c:       "|",
}

func (k *KeyEntity) GetName() string {
	return k.Name
}

func (k *KeyEntity) GetScript() string {
	return k.Script
}

func (k *KeyEntity) SetScript(script string) {
	k.Script = script
}
