// Package opt adds a DSL on top of the standard flag package

package opt

import (
	"flag"
	"os"
)

var CommandLine = New(flag.CommandLine)

type FlagSet struct {
	flag.FlagSet
	flags map[string]func()
}

func (fs *FlagSet) Add(name string, f func()) {
	fs.flags[name] = f
}

func New(fs *flag.FlagSet) *FlagSet {
	return &FlagSet{
		*fs,
		make(map[string]func()),
	}
}

func Parse() {
	CommandLine.Parse(os.Args[1:])
}

func (fs *FlagSet) Parse(arguments []string) error {
	defer fs.Visit(func(f *flag.Flag) {
		if fu, ok := fs.flags[f.Name]; ok {
			fu()
		}
	})
	return fs.FlagSet.Parse(arguments)
}

//go:generate go run genopt.go
