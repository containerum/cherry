package exampleErrors

import (
	bytes "bytes"
	cherry "git.containerum.net/ch/cherry"
	template "text/template"
)

const (
	ValidationError = "field {{.Field}}}, want {{.ValidVal}}, have {{.InvalidVal}}"
)

// ErrInvalidCheese error
// returned in case of mouse complaints
func ErrInvalidCheese(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "invalid cheese in the trap", StatusHTTP: 400, ID: cherry.ErrID{SID: 0x49, Kind: 0x78}, Details: []string{"My name is {{.Mouse}}, from {{.Package}}!"}}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}
func renderTemplate(templText string) string {
	buf := &bytes.Buffer{}
	templ, err := template.New("").Parse(templText)
	if err != nil {
		return err.Error()
	}
	err = templ.Execute(buf, map[string]string{"Package": "example-errors", "Mouse": "Jerry"})
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
