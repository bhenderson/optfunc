package opt

// DO NOT EDIT: automatically generated

import (
	"time"
)

func (fs *FlagSet) Bool(name, usage string, f func(bool)) {
	var p bool
	fs.BoolVar(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Bool(name, usage string, f func(bool)) {
	CommandLine.Bool(name, usage, f)
}

func (fs *FlagSet) Duration(name, usage string, f func(time.Duration)) {
	var p time.Duration
	fs.DurationVar(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Duration(name, usage string, f func(time.Duration)) {
	CommandLine.Duration(name, usage, f)
}

func (fs *FlagSet) Float64(name, usage string, f func(float64)) {
	var p float64
	fs.Float64Var(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Float64(name, usage string, f func(float64)) {
	CommandLine.Float64(name, usage, f)
}

func (fs *FlagSet) Int(name, usage string, f func(int)) {
	var p int
	fs.IntVar(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Int(name, usage string, f func(int)) {
	CommandLine.Int(name, usage, f)
}

func (fs *FlagSet) Int64(name, usage string, f func(int64)) {
	var p int64
	fs.Int64Var(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Int64(name, usage string, f func(int64)) {
	CommandLine.Int64(name, usage, f)
}

func (fs *FlagSet) String(name, usage string, f func(string)) {
	var p string
	fs.StringVar(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func String(name, usage string, f func(string)) {
	CommandLine.String(name, usage, f)
}

func (fs *FlagSet) Uint(name, usage string, f func(uint)) {
	var p uint
	fs.UintVar(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Uint(name, usage string, f func(uint)) {
	CommandLine.Uint(name, usage, f)
}

func (fs *FlagSet) Uint64(name, usage string, f func(uint64)) {
	var p uint64
	fs.Uint64Var(&p, name, p, usage)
	fs.Add(name, func() { f(p) })
}

func Uint64(name, usage string, f func(uint64)) {
	CommandLine.Uint64(name, usage, f)
}
