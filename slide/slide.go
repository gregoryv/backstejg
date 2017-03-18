package slide

import (
	"github.com/gregoryv/stejg/act"
	"strings"
)

var (
	size, x, y, ident int32 = 78, size, gold(1, size), 0
	fontColor               = "999999"
	fontSize                = size
)

func SetSize(s int32) {
	size = s
}

func SetFontColor(color string) {
	fontColor = color
}

func gold(min int, s int32) int32 {
	res := float32(s)
	if min <= 0 {
		return size
	}
	for {
		res = res / 1.61 // golden mean
		min--
		if min == 0 {
			break
		}
	}
	return int32(res)
}

func Basic(txt string) {
	parts := strings.Split(txt, "\n")
	for _, line := range parts {
		if strings.Index(line, "# ") == 0 {
			fontSize, ident = gold(1, size), 0
			li(line[2:], "FreeSerif")
			continue
		}
		if strings.Index(line, "## ") == 0 {
			fontSize, ident = gold(2, size), 0
			li(line[3:], "FreeSerif")
			continue
		}
		if line == "" {
			line = " "
		}
		fontSize, ident = gold(3, size), 0
		li(line, "FreeSans")
	}
}

func li(txt, font string) {
	a := &act.Event{
		Code:      act.NONE,
		Delay:     1,
		Text:      txt,
		FontColor: fontColor,
		FontSize:  int(fontSize),
		Font:      font,
		X:         x + ident,
		Y:         y,
	}
	ident = 0
	y += fontSize + gold(2, fontSize) // New line
	send(a)
}

func send(a *act.Event) {
	act.SendEvent(a, "localhost:9994")
}
