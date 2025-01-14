package test

import (
	"fmt"
	"strings"
	"tb2vb/internal/tb"
)

func seekAttributesTest() {
	twinBasicCode := `
	[PredeclaredId(True)]
	[Description("This is my awesome form.")]
	[VB_Name("frmMainForm")]
	Begin VB.Form Form1
		Caption         =   "Form1"
		ClientHeight    =   3000
		ClientLeft      =   60
		ClientTop       =   345
		ClientWidth     =   4500
	End
    `
	lines := strings.Split(twinBasicCode, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := tb.GetTokens(line)

		object := tb.FindObject(tokens)
		if object.Type != "" {
			// Generate VB code
			vbCode := tb.GetVbObject(object)
			fmt.Println(vbCode)
		} else {
			fmt.Println("Error")
		}
	}
}
