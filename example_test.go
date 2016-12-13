package optfunc_test

import (
	"flag"
	"fmt"

	"github.com/bhenderson/optfunc"
)

var (
	args = []string{"-q", "-qa", "rab oof"}
	fs   = flag.NewFlagSet("my-command", flag.PanicOnError)
)

func Example() {
	fs.Var(optfunc.Bool("q", "run a query", query))
	fs.Var(optfunc.String("qa", "query all the things", queryAll))

	fs.Parse(args)
	// Output: hello from query
	// hello rab oof
}

func query(b bool) {
	if b {
		fmt.Println("hello from query")
	}
}

func queryAll(s string) {
	fmt.Println("hello", s)
}
