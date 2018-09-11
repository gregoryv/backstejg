// texter, text actor for stejg, stejg.7de.se
package main

import (
	"bytes"
	"flag"
	"github.com/gregoryv/backstejg/text"
	"io/ioutil"
	"os"
	"path/filepath"
)

var size = flag.Int("fs", 18, "font size of title, other text is adapted using golden mean")
var fontColor = flag.String("fc", "999999", "font color")
var x = flag.Int("x", 0, "x position")
var y = flag.Int("y", 0, "y position")
var showPath = flag.Bool("p", false, "show path to file")

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		print("Usage: texter [OPTIONS] text_file\n\nOptions\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Load the entire file, useless to try to render very large
	// files anyway so this is ok.
	file := flag.Args()[0]
	txt, err := ioutil.ReadFile(file)
	if err != nil {
		print(err.Error(), "\n")
		flag.Usage()
		os.Exit(1)
	}

	if *showPath {
		format := text.NewPlain()
		format.Render(string(file))
		*y += *size // Render text below filename
	}

	ext := filepath.Ext(file)
	format := text.NewPlain()
	format.SetSize(int32(*size))
	format.SetFontColor(*fontColor)
	format.SetPosition(int32(*x), int32(*y))

	switch ext {
	case ".md":
		format.RenderMarkdown(string(txt))
	default:
		tab := []byte("\t")
		spaces := []byte("    ")
		txt = bytes.Replace(txt, tab, spaces, -1)
		format.Render(string(txt))
	}
}
