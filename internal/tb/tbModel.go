package tb

type TwinForm struct {
	Name            string        `json:"Name"`
	Caption         string        `json:"Caption"`
	Width           int           `json:"Width"`
	Height          int           `json:"Height"`
	StartUpPosition string        `json:"StartUpPosition"`
	Children        []TwinControl `json:"_children"`
}

//
// (C) Bilbo Backends 2025
//    All what I've seen in Visual Basic user-controls
// is Text-elements, which contains Text Property (e.g. CommandBox, TextBox)
// and Caption-elements, which contains Caption Property (e.g. Form, Label, Command)
//    First idea is divide them by 2 groups: Text-Changed (by user) / Text-Unchangable (by user too)
// I suggest is bad idea, but it better than nothing... Next time I'll rebuild determination logic
//

type TwinControl struct {
	ClassName string `json:"_className"`
	Name      string `json:"Name"`
	Caption   string `json:"Caption"`
	Left      int    `json:"Left"`
	Top       int    `json:"Top"`
	Width     int    `json:"Width"`
	Height    int    `json:"Height"`
}

type TwinTextControl struct {
	ClassName string `json:"_className"`
	Name      string `json:"Name"`
	Text      string `json:"Text"`
	Left      int    `json:"Left"`
	Top       int    `json:"Top"`
	Width     int    `json:"Width"`
	Height    int    `json:"Height"`
}

type TwinAttribute struct {
	Name      string
	Value     string
	Arguments []string
}

type TwinObject struct {
	Name       string
	Body       string
	Type       string
	Attributes []TwinAttribute
}
