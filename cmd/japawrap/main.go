package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/agatan/japawrap"
)

var version string = "0.0.1"

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run() int {
	var v bool
	flag.BoolVar(&v, "version", false, "Print version information and quit")
	flag.Parse()

	if v {
		fmt.Fprintf(c.outStream, "japawrap %s", version)
		return 0
	}

	w := japawrap.New(`<span class="wordwrap">`, `</span>`)

	for _, fn := range flag.Args() {
		var input io.Reader
		if fn == "-" {
			input = os.Stdin
		} else {
			fp, err := os.Open(fn)
			if err != nil {
				fmt.Fprintf(c.errStream, err.Error())
				return 1
			}
			defer fp.Close()
			input = fp
		}
		s, err := ioutil.ReadAll(input)
		if err != nil {
			fmt.Fprintf(c.errStream, err.Error())
			return 1
		}
		fmt.Fprintln(c.outStream, w.Do(string(s)))
	}

	return 0
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run())
}
