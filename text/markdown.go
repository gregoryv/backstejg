package text

import (
	"strings"
)

func (p *Plain) RenderMarkdown(txt string) {
	for _, line := range strings.Split(txt, "\n") {
		// Poor mans parsing of markdown, far from complete
		// expanded on a need to basis
		if line == "" {
			line = " "
		}
		switch true {
		case strings.Index(line, "# ") == 0:
			p.FontSize = p.gold(3, p.FontSize)
			p.write(line[2:], "FreeSerif")
		case strings.Index(line, "## ") == 0:
			p.FontSize = p.gold(2, p.FontSize)
			p.write(line[3:], "FreeSerif")
		case strings.Index(line, "### ") == 0:
			p.FontSize = p.gold(1, p.FontSize)
			p.write(line[4:], "FreeSerif")
		case strings.Index(line, "    ") == 0:
			p.FontSize = p.gold(1, p.FontSize)
			p.write(line, "FreeMono")
		default:
			p.FontSize = p.gold(0, p.Size)
			p.write(line, "FreeSans")
		}
	}
}
