package models

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"

	"github.com/containerum/cherry"
	"github.com/dave/jennifer/jen"
)

const (
	CherryPath = "github.com/containerum/cherry"
)

type TOMLerror struct {
	Name       string
	Message    string
	StatusHTTP int
	SID        cherry.ErrSID
	Kind       cherry.ErrKind
	Comment    string
	Details    []string
}

func (terror *TOMLerror) SourceCodeID() string {
	name := strings.TrimFunc(strings.Title(terror.Name),
		func(r rune) bool {
			return unicode.IsPunct(r) || unicode.IsSpace(r)
		})
	if !strings.HasPrefix(name, "Err") {
		name = "Err" + name
	}
	return name
}

func (terror *TOMLerror) Cherry() *cherry.Err {
	return &cherry.Err{
		Message:    terror.Message,
		StatusHTTP: terror.StatusHTTP,
		ID: cherry.ErrID{
			SID:  cherry.ErrSID(terror.SID),
			Kind: cherry.ErrKind(terror.Kind),
		},
		Details: terror.Details,
	}
}
func (terror *TOMLerror) GenerateSource() *jen.Statement {
	fn := jen.Func().Id(terror.SourceCodeID()).
		Params(jen.Id("params").Op("...").Func().Params(jen.Op("*").Qual(CherryPath, "Err"))).
		Op("*").Qual(CherryPath, "Err").
		Block(jen.Id("err").Op(":=").Lit(terror.Cherry()),
			jen.For(jen.Id("_").Op(",").Id("param").Op(":=").Range().Id("params")).Block(
				jen.Id("param").Call(jen.Id("err")),
			),
			jen.For(jen.Id("i").Op(",").Id("detail").Op(":=").Range().Id("err").Dot("Details").Block(
				jen.Id("det").Op(":=").Id("renderTemplate").Call(jen.Id("detail")),
				jen.Id("err").Dot("Details").Index(jen.Id("i")).Op("=").Id("det"),
			)),
			jen.Return(jen.Id("err")),
		)
	if terror.Comment != "" {
		return jen.Comment(terror.SourceCodeID() + " error ").Line().
			Add(buildComments(terror.Comment).Add(fn))
	}
	return fn
}

func sanitizeCommentLine(commentLine string) string {
	commentLine = strings.TrimSpace(commentLine)
	commentLine = strings.TrimPrefix(commentLine, "//")
	commentLine = strings.TrimPrefix(commentLine, "/*")
	commentLine = strings.TrimSuffix(commentLine, "*/")
	commentLine = strings.Replace(commentLine, "\n", "", -1)
	return commentLine
}
func buildComments(text string) *jen.Statement {
	var comments *jen.Statement
	i := 0
	scanner := bufio.NewScanner(bytes.NewReader([]byte(text)))
	for scanner.Scan() {
		commentLine := sanitizeCommentLine(scanner.Text())
		if commentLine == "" {
			continue
		}
		if i == 0 {
			i++
			comments = jen.Comment(commentLine)
			continue
		}
		comments.Add(jen.Line().Comment(commentLine))
	}
	return comments.Line()
}
