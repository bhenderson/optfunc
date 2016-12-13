package opt_test

import (
	"flag"
	"fmt"

	"github.com/bhenderson/opt"
)

var (
	args = []string{"-q", "-qa", "rab oof"}
	fs   = flag.NewFlagSet("my-command", flag.PanicOnError)
)

func Example() {
	fs.Var(opt.Bool("q", "run a query", query))
	fs.Var(opt.String("qa", "query all the things", queryAll))

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
