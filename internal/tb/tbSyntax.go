package tb

import "bytes"

//
// Twin Basic syntax analyser
// Needs for special bracet declarations, used by TwinBasic IDE
//
// TwinBasic Form construt looks like JSON formatted
// file with all VBA/VB6 User-Controls properties
// (e.g. Caption from Label control)
//
// [.tbform] files memorized here and will have sent
// to Visual Basic Syntax transformer
//
// [.twin] files contains TwinBasic IDE Attributes
// which not compatible with Visual Basic Classic syntax (as I remember)
//

var (
	formBytes bytes.Buffer
	codeBytes bytes.Buffer
)

func SetForm() {

}
