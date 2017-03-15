package act

import (
	"fmt"
	"time"
)

const (
	ERR = int32(iota)
	NONE
	QUIT
	PAUSE
	CLEAR
	HIDE
	SAVE
	REDRAW

	// Effects
	FADE_OUT
	PULSE
)

var NamedCommands = map[string]int32{
	"clear": CLEAR,
	"quit":  QUIT,
	"pause": PAUSE,
	"hide":  HIDE,
	"save":  SAVE,
}

var NamedEffects = map[string]int32{
	"none":     NONE,
	"fade_out": FADE_OUT,
	"pulse":    PULSE,
}

var EventNames = map[int32]string{}

func init() {
	for name, code := range NamedCommands {
		EventNames[code] = name
	}
	for name, code := range NamedEffects {
		EventNames[code] = name
	}
}

type Event struct {
	Code      int32
	ImageURI  string
	Position  string
	X, Y      int32
	Delay     time.Duration
	FontColor string
	Font      string
	FontSize  int
	BgColor   string
	Text      string
	Multiply  int
}

func NewEvent() *Event {
	return &Event{
		Code:      NONE,
		ImageURI:  "",
		Position:  "",
		X:         0,
		Y:         0,
		Delay:     time.Duration(0),
		FontColor: "white",
		FontSize:  34,
		BgColor:   "000000",
		Text:      "",
	}
}

func (s *Event) String() string {
	return fmt.Sprintf("%s %s", EventNames[s.Code], s.ImageURI)
}
