// +build ignore

package main

import (
	"bytes"
	"os"

	"go/format"

	"github.com/bhenderson/go/src/text/template"
)

var flagTypes = []struct{ Name, Type string }{
	{"Bool", "bool"},
	{"Duration", "time.Duration"},
	{"Float64", "float64"},
	{"Int", "int"},
	{"Int64", "int64"},
	{"String", "string"},
	{"Uint", "uint"},
	{"Uint64", "uint64"},
}

var temp = template.Must(template.New("on").Parse(`package opt

// DO NOT EDIT: automatically generated

import (
	"time"
)

{{ range . }}

func (fs *FlagSet) {{.Name}}(name, usage string, f func({{.Type}})) {
	var p {{.Type}}
	fs.{{.Name}}Var(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func {{.Name}}(name, usage string, f func({{.Type}})) {
	CommandLine.{{.Name}}(name, usage, f)
}
{{ end }}
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
