package models

import (
	"bufio"
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
)

func TestErrorFuncGeneration(test *testing.T) {
	tomlerr := &TOMLerror{
		Name:       "ErrCantGenerateCode",
		Message:    "error while generating code",
		StatusHTTP: 500,
		SID:        "noice",
		Kind:       2049,
		Details: []string{
			"noice {.thing}",
		},
		Comment: `test 
		error func`,
	}
	pkg := jen.NewFile("test")
	pkg.Add(tomlerr.GenerateSource())
	src := fmt.Sprintf("%#v", pkg)
	fileSet := token.NewFileSet()
	_, err := parser.ParseFile(fileSet, "", src, parser.AllErrors)
	test.Logf("\n%s", enumerateLines(src))
	if err != nil {
		test.Fatalf("error while validating ast: %v", err)
	}
}

func TestServiceGenerator(test *testing.T) {
	service := &Service{
		Name: "te",
		SID:  "sid",
		Error: []TOMLerror{
			{
				Name:    "InvalidMessage",
				Kind:    201,
				Message: "invalid message",
				Details: []string{
					"{{.Foo}}",
				},
				StatusHTTP: 409,
			},
			{
				Name:       "NotEnoghCheese",
				Kind:       200,
				Message:    "not enough cheese",
				StatusHTTP: 500,
			},
		},
		Keys: map[string]string{
			"Foo": "bar",
		},
	}
	src, err := service.GenerateSourceString()
	test.Logf("\n%s", enumerateLines(src))
	if err != nil {
		test.Fatalf("error while validating ast: %v", err)
	}
}

func enumerateLines(text string) string {
	var lines []string
	i := 0
	scanner := bufio.NewScanner(bytes.NewReader([]byte(text)))
	for scanner.Scan() {
		i++
		lines = append(lines, fmt.Sprintf("%03d : %s", i, scanner.Text()))
	}
	return strings.Join(lines, "\n")
}
