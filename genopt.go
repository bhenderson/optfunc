// +build ignore

package main

import (
	"bytes"
	"os"
	"strings"

	"go/format"

	"github.com/bhenderson/go/src/text/template"
)

var flagTypes = []flagType{
	{"Bool", "bool", "strconv.ParseBool(s)"},
	{"Duration", "time.Duration", "strconv.ParseInt(s, 0, 64)"},
	{"Float64", "float64", "strconv.ParseFloat(s, 64)"},
	{"Int", "int", "strconv.ParseInt(s, 0, 64)"},
	{"Int64", "int64", "strconv.ParseInt(s, 0, 64)"},
	{"String", "string", "s, error(nil)"},
	{"Uint", "uint", "strconv.ParseUint(s, 0, 64)"},
	{"Uint64", "uint64", "strconv.ParseUint(s, 0, 64)"},
}

type flagType struct{ Name, Type, ParseFunc string }

func (ft flagType) Fname() string {
	return strings.ToLower(ft.Name) + "Func"
}

var temp = template.Must(template.New("on").Parse(`package opt

// DO NOT EDIT: automatically generated

import (
	"time"
	"strconv"
)

{{ range . }}

func {{.Name}}(name, usage string, f func({{.Type}})) (Value, string, string) {
	return {{.Fname}}(f), name, usage
}

type {{.Fname}} func({{.Type}})

func (f {{.Fname}}) Set(s string) error {
	v, err := {{.ParseFunc}}
	f({{.Type}}(v))
	return err
}
func (f {{.Fname}}) String() string { return "" }
{{ if eq .Type "bool" -}}
func (f {{.Fname}}) IsBoolFlag() bool { return true }
{{- end }}
{{ end }}

// Copy of flag.Value
type Value interface {
	Set(string) error
	String() string
}

`))

func main() {
	var buf bytes.Buffer
	err := temp.Execute(&buf, flagTypes)
	if err != nil {
		panic(err)
	}
	s, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	f, err := os.Create("optvar.go")
	if err != nil {
		panic(err)
	}
	f.Write(s)
}
