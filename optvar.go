package optfunc

// DO NOT EDIT: automatically generated

import (
	"strconv"
	"time"
)

func Bool(name, usage string, f func(bool)) (Value, string, string) {
	return boolFunc(f), name, usage
}

type boolFunc func(bool)

func (f boolFunc) Set(s string) error {
	v, err := strconv.ParseBool(s)
	f(bool(v))
	return err
}
func (f boolFunc) String() string   { return "" }
func (f boolFunc) IsBoolFlag() bool { return true }

func Duration(name, usage string, f func(time.Duration)) (Value, string, string) {
	return durationFunc(f), name, usage
}

type durationFunc func(time.Duration)

func (f durationFunc) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	f(time.Duration(v))
	return err
}
func (f durationFunc) String() string { return "" }

func Float64(name, usage string, f func(float64)) (Value, string, string) {
	return float64Func(f), name, usage
}

type float64Func func(float64)

func (f float64Func) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	f(float64(v))
	return err
}
func (f float64Func) String() string { return "" }

func Int(name, usage string, f func(int)) (Value, string, string) {
	return intFunc(f), name, usage
}

type intFunc func(int)

func (f intFunc) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	f(int(v))
	return err
}
func (f intFunc) String() string { return "" }

func Int64(name, usage string, f func(int64)) (Value, string, string) {
	return int64Func(f), name, usage
}

type int64Func func(int64)

func (f int64Func) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	f(int64(v))
	return err
}
func (f int64Func) String() string { return "" }

func String(name, usage string, f func(string)) (Value, string, string) {
	return stringFunc(f), name, usage
}

type stringFunc func(string)

func (f stringFunc) Set(s string) error {
	v, err := s, error(nil)
	f(string(v))
	return err
}
func (f stringFunc) String() string { return "" }

func Uint(name, usage string, f func(uint)) (Value, string, string) {
	return uintFunc(f), name, usage
}

type uintFunc func(uint)

func (f uintFunc) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	f(uint(v))
	return err
}
func (f uintFunc) String() string { return "" }

func Uint64(name, usage string, f func(uint64)) (Value, string, string) {
	return uint64Func(f), name, usage
}

type uint64Func func(uint64)

func (f uint64Func) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	f(uint64(v))
	return err
}
func (f uint64Func) String() string { return "" }

// Copy of flag.Value
type Value interface {
	Set(string) error
	String() string
}
