package tb

import (
	"fmt"
	"regexp"
	"strings"
)

func GetTokens(input string) []string {
	var tokens []string
	// Regular expression to capture attributes
	// in square brackets:
	// 		[Attribute("value", param2)]
	regex := regexp.MustCompile(`(\[[A-Za-z_]+\s*(\(\s*("[^"]*"|[^)]*?)\s*(,\s*("[^"]*"|[^)]*?)\s*)*\))?\])|([A-Za-z_][A-Za-z_0-9]*)`) // pain Model
	matches := regex.FindAllString(input, -1)

	for _, match := range matches {
		tokens = append(tokens, strings.TrimSpace(match))
	}
	return tokens
}

func FindAttributes(tokens []string) (attributes []TwinAttribute, restTokens []string) {

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Skip attribute
		if token[0] != '[' {
			restTokens = tokens[i:]
			return
		}


		attr := TwinAttribute{}
		attrValue := strings.TrimSuffix(strings.TrimPrefix(token, "["), "]")
		attrArgs := strings.Split(attrValue, "(")
		attr.Name = strings.TrimSpace(attrArgs[0])


		if len(attrArgs) > 1 {

			value := strings.TrimSpace(attrArgs[1])
			value = strings.TrimSuffix(value, ")")

			params := strings.Split(value, ",")

			for _, param := range params {
				attr.Arguments = append(attr.Arguments, strings.TrimSpace(param))
			}
		}

		attributes = append(attributes, attr)
	}
	return
}

func FindObject(tokens []string) (objectDefinition TwinObject) {
	attributes, rest := FindAttributes(tokens)

	if len(rest) < 2 {
		return
	}

	objectDefinition.Type = rest[0]
	objectDefinition.Name = rest[1]
	objectDefinition.Attributes = attributes
	objectDefinition.Body = strings.Join(rest[2:], " ")

	return objectDefinition
}

func convertAttribute(attr TwinAttribute) (vbAttr string) {
	// Searching for known compatible Visual Basic attributes
	switch strings.ToLower(attr.Name) {
	case "predeclaredid":
		if len(attr.Arguments) == 0 || strings.ToLower(attr.Arguments[0]) == "true" {
			vbAttr = "Attribute VB_PredeclaredId = True"
		} else {
			vbAttr = "Attribute VB_PredeclaredId = False"
		}

	case "description":
		if len(attr.Arguments) > 0 {
			vbAttr = fmt.Sprintf("' Description: %s", strings.Trim(attr.Arguments[0], "\""))
		} else {
			vbAttr = fmt.Sprintf("' <No description>")
		}
	case "vb_name":
		if len(attr.Arguments) > 0 {
			vbAttr = fmt.Sprintf("Attribute VB_Name = %q", strings.Trim(attr.Arguments[0], "\""))
		}
	}
	return
}

// GetVbObject constructs Visual Basic
// object from object definition
// returns ready VB expression
func GetVbObject(objectDefinition TwinObject) (vbCode string) {
	for _, attr := range objectDefinition.Attributes {
		convertedAttr := convertAttribute(attr)
		if convertedAttr != "" {
			vbCode += convertedAttr + "\n"
		}
	}

	vbCode += fmt.Sprintf(
		"%s %s %s",
		objectDefinition.Type,
		objectDefinition.Name,
		objectDefinition.Body)

	return
}