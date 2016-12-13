package opt_test

import (
	"flag"
	"fmt"

	"github.com/bhenderson/opt"
)

var (
	args = []string{"-q", "-qa", "arb oof"}
	topt = opt.New(flag.NewFlagSet("my-command", flag.PanicOnError))
)

func Example() {
	topt.Bool("q", "run a query", query)
	topt.String("qa", "query all the things", queryAll)

	topt.Parse(args)
	// Output: hello from query
	// hello arb oof
}

func query(b bool) {
	if b {
		fmt.Println("hello from query")
	}
}

func queryAll(s string) {
	fmt.Println("hello", s)
}
