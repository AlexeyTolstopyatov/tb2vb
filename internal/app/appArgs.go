package app

import "fmt"

// GetFilesFromArgs returns (tbform , twin) files tuple
func GetFilesFromArgs(args []string) (string, string) {
	// app -x file -y file
	// 0	1	2	3	4
	if len(args) < 5 {
		panic(fmt.Sprintf("%d smaller than 5", len(args)))
	}

	if args[1] == "-c" {
		return args[4], args[2]
	} else {
		return args[2], args[4]
	}
}
