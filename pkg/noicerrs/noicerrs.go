package noicerrs

import (
	bytes "bytes"
	cherry "git.containerum.net/ch/cherry"
	template "text/template"
)

const ()

func ErrUnableToOpenTOMLfile(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "unable to open file", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x1}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableToParseTOMLfile(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "unable to parse TOML file", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x2}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableToCreatePackageDir(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "unable to create package name", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x3}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableToWriteSourcefile(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "unable to write source file", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x4}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUndefinedSID(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "undefined SID", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x5}, Details: []string{"must be > 0"}}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUndefinedPackageName(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "undefined package name", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x6}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUndefinedKind(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "undefined error kind", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x7}, Details: []string{"in error {{.ErrorName}}"}}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUndefinedStatusHTTP(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "undefined status HTTP", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x8}, Details: []string{"in error {{.ErrorName}}"}}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableToWriteJSONfile(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "unable to write json file", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0x9}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrConflictingKinds(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "conflicting kinds", StatusHTTP: 418, ID: cherry.ErrID{SID: 0x109, Kind: 0xa}, Details: []string(nil)}
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
	err = templ.Execute(buf, map[string]string{})
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
