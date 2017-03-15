package slide

import (
	"github.com/gregoryv/stejg/act"
	"strings"
)

var (
	light, dark, red         = "999999", "222222", "red"
	h2Color                  = "555555"
	size, x, y, ident  int32 = 78, size, gold(1, size), 0
	fontColor, bgColor       = light, dark
	fontSize                 = size
)

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
	clear()
	parts := strings.Split(txt, "\n")
	for _, line := range parts {
		if strings.Index(line, "# ") == 0 {
			fontColor, fontSize, ident = light, gold(1, size), 0
			li(line[2:])
			continue
		}
		if strings.Index(line, "## ") == 0 {
			fontColor, fontSize, ident = h2Color, gold(1, size), gold(2, size)
			li(line[3:])
			continue
		}
		if line != "" {
			fontColor, fontSize, ident = light, gold(2, size), gold(1, size)
			li(line)
		}
	}
}

func li(txt string) {
	a := &act.Event{
		Code:      act.NONE,
		Delay:     1,
		Text:      txt,
		FontColor: fontColor,
		FontSize:  int(fontSize),
		X:         x + ident,
		Y:         y,
	}
	ident = 0
	y += fontSize + gold(2, fontSize) // New line
	send(a)
}

func clear() {
	send(&act.Event{
		Code:    act.CLEAR,
		BgColor: bgColor,
	})
	send(&act.Event{
		Code:    act.HIDE,
		BgColor: bgColor,
	})
}

func send(a *act.Event) {
	act.SendEvent(a, "localhost:9994")
}
