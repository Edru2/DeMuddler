package main
import(
	"encoding/xml"
)

type TriggerPackage struct {
	XMLName      xml.Name       `xml:"TriggerPackage" json:"-"`
	Triggers     []Trigger      `xml:"Trigger" json:"-"`
	TriggerGroup []TriggerGroup `xml:"TriggerGroup" json:"-"`
}

type Trigger struct {
	IsActive            string         `xml:"isActive,attr"`
	IsFolder            string         `xml:"isFolder,attr"`
	IsMultiline         string         `xml:"isMultiline,attr"`
	IsPerlSlashGOption  string         `xml:"isPerlSlashGOption"`
	IsColorizerTrigger  string         `xml:"isColorizerTrigger"`
	IsFilterTrigger     string         `xml:"isFilterTrigger"`
	IsSoundTrigger      string         `xml:"isSoundTrigger"`
	IsColorTrigger      string         `xml:"isColorTrigger"`
	IsColorTriggerFg    string         `xml:"isColorTriggerFg"`
	IsColorTriggerBg    string         `xml:"isColorTriggerBg"`
	Name                string         `xml:"name"`
	Script              string         `xml:"script"`
	TriggerType         int            `xml:"triggerType"`
	ConditionLineDelta  int            `xml:"conditonLineDelta"`
	MFgColor            string         `xml:"mFgColor"`
	MBgColor            string         `xml:"mBgColor"`
	MSoundFile          string         `xml:"mSoundFile"`
	MStayOpen           string         `xml:"mStayOpen"`
	MCommand            string         `xml:"mCommand"`
	PackageName         string         `xml:"packageName"`
	Path                string         `xml:"path"`
	RegexCodeList       []string       `xml:"regexCodeList>string"`
	RegexCodeProperties []int          `xml:"regexCodePropertyList>integer"`
	Triggers            []Trigger      `xml:"Trigger" json:"children,omitempty"`
	TriggerGroup        []TriggerGroup `xml:"TriggerGroup" json:"-"`
}

type TriggerGroup struct {
	Trigger
}

type JSONTrigger struct {
	Name           string        `json:"name"`
	IsActive       string        `json:"isActive"`
	IsFolder       string        `json:"isFolder"`
	Command        string        `json:"command,omitempty"`
	Multiline      string        `json:"multiline"`
	MultilineDelta string        `json:"multielineDelta"`
	Matchall       string        `json:"matchall"`
	Filter         string        `json:"filter"`
	FireLength     string        `json:"fireLength"`
	SoundFile      string        `json:"soundFile,omitempty"`
	Highlight      string        `json:"highlight"`
	HighlightFG    string        `json:"highlightFG,omitempty"`
	HighlightBG    string        `json:"highlightBG,omitempty"`
	Patterns       []Pattern     `json:"patterns"`
	Script         string        `json:"script"`
	Children       []JSONTrigger `json:"children,omitempty"`
}

type Pattern struct {
	Pattern string `json:"pattern"`
	Type    string `json:"type"`
}

