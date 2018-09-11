package text

import (
	"fmt"
	"github.com/gregoryv/backstejg/act"
	"os"
	"strings"
)

func (p *Plain) gold(min int, s int32) int32 {
	res := float32(s)
	if min <= 0 {
		return p.Size
	}
	for {
		res = res * 1.61 // golden mean
		min--
		if min == 0 {
			break
		}
	}
	return int32(res)
}

func (p *Plain) write(txt, font string) {
	a := &act.Event{
		Code:      act.NONE,
		Delay:     1,
		Text:      txt,
		FontColor: p.FontColor,
		FontSize:  int(p.FontSize),
		Font:      font,
		X:         p.X + p.Ident,
		Y:         p.Y,
	}
	p.Y += p.FontSize + p.gold(0, p.FontSize)/2 // New line
	p.send(a)
}

func (p *Plain) send(a *act.Event) {
	if p.Bind == "" {
		fmt.Fprint(os.Stderr, "Bind is not set")
		os.Exit(1)
	}
	act.SendEvent(a, p.Bind)
}

type Plain struct {
	FontSize   int32
	Ident      int32
	Size, X, Y int32
	FontColor  string
	Bind       string
}

func NewPlain() *Plain {
	size := int32(18)
	p := &Plain{
		FontSize:  size,
		Ident:     0,
		Size:      size,
		Y:         0,
		FontColor: "999999",
		Bind:      os.Getenv("STEJG_BIND"),
	}
	p.X = p.gold(1, size)
	return p
}

func (p *Plain) Render(txt string) {
	for _, line := range strings.Split(txt, "\n") {
		p.write(line, "FreeMono")
	}
}

func (p *Plain) SetPosition(xpos, ypos int32) {
	p.X = xpos
	p.Y = ypos
}

func (p *Plain) SetSize(s int32)           { p.Size = s }
func (p *Plain) SetFontColor(color string) { p.FontColor = color }
