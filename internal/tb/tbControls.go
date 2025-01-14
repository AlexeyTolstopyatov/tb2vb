package tb

func GetVbControlByString(className string) string {
	switch className {
	case "Frame":
		return "VB.Frame"
	case "TextBox":
		return "VB.TextBox"
	case "CommandButton":
		return "VB.CommandButton"
	case "Label":
		return "VB.Label"
	case "ListBox":
		return "VB.ListBox"
	default:
		return ""
	}
}
