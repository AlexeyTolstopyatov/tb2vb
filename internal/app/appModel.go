package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"tb2vb/internal/tb"
)

//
// (C) Bilbo Backends 2025
// TwinBasic project files divides by 2 kinds ->
// 		*.tbform -- Twin Basic forms constructor
//		*.twin	 -- Twin Basic code of form
//					Contains all events and other business-logic
//					functions/procedures.
//	  I appologize, Standard Visual Basic .frm file has not store
// the Constructor method (Sub New()) because, all Projects, what I've done
// had not Class identification. They were look like Module files but
// have -frm prefix in naming.
//    So, it will be 1'st task - compile VB-code of events from .twin file
// and collect all constructor data (JSON-format) from .tbform file.
//	  Next task in schedule is -- determine TwinBasic Attributes ->
// replace them with Visual Basic 5 attribute-expressions
//
// [PredeclaredId()] replaces with
// Attribute PredeclaredId = True
//

// Init
// Main constructor method for filling
// TwinBasic-module structure
func Init(twin string, tbform string) {
	// Read and Write Constructor path
	data, err := os.ReadFile(tbform)
	if err != nil {
		panic(err)
	}

	var form tb.TwinForm
	err = json.Unmarshal(data, &form)
	if err != nil {
		panic(err)
	}
	code, err := os.ReadFile(twin)
	if err != nil {
		panic(err)
	}

	frmContent := generateConstructorPart(form)
	frmEvents := generateCodePart(bytes.NewBuffer(code))

	frmContent += frmEvents

	// Saving...
	err = os.WriteFile(".frm", []byte(frmContent), 0644)
	if err != nil {
		fmt.Println("Unable to write result", err)
		return
	}

	fmt.Printf("done")
}

func generateConstructorPart(form tb.TwinForm) string {
	content := fmt.Sprintf(`VERSION 5.00
Begin VB.Form %s\
	Caption         =   "%s"
	ClientHeight    =   %d
	ClientWidth     =   %d
	StartUpPosition =   %s
`,
		form.Name,
		form.Caption,
		form.Height,
		form.Width,
		form.StartUpPosition)
	// Enumerate known controls in Form
	for _, control := range form.Children {
		controlType := tb.GetVbControlByString(control.ClassName)
		if controlType == "" {
			continue
		}

		content += fmt.Sprintf(`
	Begin %s %s
    	Caption         =   "%s"
        Height          =   %d
        Left            =   %d
        Top             =   %d
        Width           =   %d
    End`,
			controlType,
			control.Name,
			control.Caption,
			control.Height,
			control.Left,
			control.Top,
			control.Width)
	}
	// End operator follows after code
	return content
}

func generateCodePart(code *bytes.Buffer) string {
	codestr := string(code.Bytes())
	lines := strings.Split(codestr, "\n")
	clines := make([]string, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := tb.GetTokens(line)
		object := tb.FindObject(tokens)

		if object.Type != "" {
			line = tb.GetVbObject(object)
		} else {
			fmt.Println("Error")
		}
		clines = append(clines, line)
	}
	return strings.Join(clines, "\n")
}
