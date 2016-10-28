package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agatan/japawrap"
)

func main() {
	s, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	w := japawrap.New(`<span class="wordwrap">`, `</span>`)
	fmt.Println(w.Do(string(s)))
}
