package vb

// Visual Basic Form unit specification
// All Visual Basic units starts with fs-prefix of unit
// and naming attribute:
//		Modules 	(.bas)		modVisualBasicModule
//		Form		(.frm)		frmMain
//		Class		(.cls)		clsXmlDocument

// Little constant definitions of Visual Basic units
// opening. (Starts from Visual Basic 5)
const (
	Version = "VERSION 5.00"
	Begin   = "Begin VB."
	End     = "End"
	Tab     = "\t"
)

// Form describes model of transformed TwinBasic format to
// Visual Basic Classic forms
type Form struct {
	Name       string
	Properties map[string]string
	Controls   map[string]interface{}
}
