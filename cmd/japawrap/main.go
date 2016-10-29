package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/agatan/japawrap"
)

var version string = "0.0.1"

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) process(w *japawrap.Wrapper, r io.Reader) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		s := sc.Text()
		fmt.Fprintln(c.outStream, w.Do(s))
	}
	return sc.Err()
}

func (c *CLI) Run() int {
	var v bool
	flag.BoolVar(&v, "version", false, "Print version information and quit")
	var open string
	var close string
	flag.StringVar(&open, "open", `<span class="wordwrap">`, "Open string")
	flag.StringVar(&close, "close", `</span>`, "Close string")
	flag.Parse()

	if v {
		fmt.Fprintf(c.outStream, "japawrap %s\n", version)
		return 0
	}

	w := japawrap.New(open, close)

	if len(flag.Args()) == 0 {
		if err := c.process(w, os.Stdin); err != nil {
			fmt.Fprintf(c.errStream, err.Error())
			return 1
		}
		return 0
	}

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
		if err := c.process(w, input); err != nil {
			fmt.Fprintf(c.errStream, err.Error())
			return 1
		}
	}

	return 0
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run())
}
