package main

import (
	"flag"
	"github.com/gregoryv/stejg/slide"
	"io/ioutil"
	"os"
)

var size = flag.Int("s", 72, "size of title, other text is adapted using golden mean")
var file = flag.String("f", "", "text file with slide content")
var fontColor = flag.String("fc", "999999", "font color")

func main() {
	flag.Parse()

	txt, err := ioutil.ReadFile(*file)
	if err != nil {
		print(err.Error(), "\n")
		flag.Usage()
		os.Exit(1)
	}

	slide.SetSize(int32(*size))
	slide.SetFontColor(*fontColor)
	slide.Basic(string(txt))
}
