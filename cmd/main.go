package main

import (
	"fmt"
	"os"
	"tb2vb/internal/app"
)

func main() {
	args := os.Args
	if len(args) < 5 {
		fmt.Println("(C) Bilbo Backends")
		fmt.Println("./tb2vb -f path/frmMain.tbform -c path/frmMain.twin")
		fmt.Println("\t -f\tSets TwinBasic Form file (JSON-like) file (.tbform)")
		fmt.Println("\t -c\tSets TwinBasic module file (.twin)")
		return
	}

	form, code := app.GetFilesFromArgs(args)
	app.Init(form, code)
}
