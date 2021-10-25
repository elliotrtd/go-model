package kiosk

import "github.com/raceresult/go-model/vbdate"

// Kiosk represents all settings of a kiosk
type Kiosk struct {
	Name            string
	Key             string
	Enabled         bool
	EnabledFrom     vbdate.VBDate
	EnabledTo       vbdate.VBDate
	TransponderMode int
	CSS             string
	Title           string
	Steps           []Step
	AfterSave       []AfterSave
}

type Step struct {
	Type  string
	Label string

	Title         string
	Text          string
	SearchFields  []SearchField
	DisplayFields []DisplayField
	EditFields    []EditField
	Settings      map[string]interface{}
}

type AfterSave struct {
	Type        string
	Value       string
	Destination string
	Filter      string
}

type DisplayField struct {
	Type  string
	Value string
	Label string
}

type EditField struct {
	Label     string
	Field     string
	Special   string
	Mandatory bool
}

type SearchField struct {
	Field string
	Hide  bool
}
