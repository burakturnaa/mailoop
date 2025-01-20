package utils

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func SanitizeHTML(input string) string {
	input = removeScriptTags(input)

	tmpl, err := template.New("html").Parse(input)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}

	// Create a buffer to hold the sanitized output
	var result bytes.Buffer

	// Execute the template and write the result to the buffer
	err = tmpl.Execute(&result, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return ""
	}
	return result.String()
}

func removeScriptTags(input string) string {
	input = strings.ReplaceAll(input, "<script>", "")
	input = strings.ReplaceAll(input, "</script>", "")
	return input
}
