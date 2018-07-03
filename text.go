package dfhackrpc

import (
	"io"
	"os"

	"github.com/BenLubar/dfhackrpc/dfproto"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

func wrapWriter(w io.Writer) io.Writer {
	if f, ok := w.(*os.File); ok && os.Getenv("TERM") != "dumb" && (isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())) {
		return colorable.NewColorable(f)
	}
	return colorable.NewNonColorable(w)
}

var colors = [...]color.Attribute{
	dfproto.CoreTextFragment_COLOR_BLACK:        color.FgBlack,
	dfproto.CoreTextFragment_COLOR_BLUE:         color.FgBlue,
	dfproto.CoreTextFragment_COLOR_GREEN:        color.FgGreen,
	dfproto.CoreTextFragment_COLOR_CYAN:         color.FgCyan,
	dfproto.CoreTextFragment_COLOR_RED:          color.FgRed,
	dfproto.CoreTextFragment_COLOR_MAGENTA:      color.FgMagenta,
	dfproto.CoreTextFragment_COLOR_BROWN:        color.FgYellow,
	dfproto.CoreTextFragment_COLOR_GREY:         color.FgWhite,
	dfproto.CoreTextFragment_COLOR_DARKGREY:     color.FgHiBlack,
	dfproto.CoreTextFragment_COLOR_LIGHTBLUE:    color.FgHiBlue,
	dfproto.CoreTextFragment_COLOR_LIGHTGREEN:   color.FgHiGreen,
	dfproto.CoreTextFragment_COLOR_LIGHTCYAN:    color.FgHiCyan,
	dfproto.CoreTextFragment_COLOR_LIGHTRED:     color.FgHiRed,
	dfproto.CoreTextFragment_COLOR_LIGHTMAGENTA: color.FgHiMagenta,
	dfproto.CoreTextFragment_COLOR_YELLOW:       color.FgHiYellow,
	dfproto.CoreTextFragment_COLOR_WHITE:        color.FgHiWhite,
}

func colorAttr(c dfproto.CoreTextFragment_Color) color.Attribute {
	if c >= 0 && int(c) < len(colors) {
		return colors[c]
	}
	return color.Reset
}

func (c *Client) outputTextFragment(text *dfproto.CoreTextFragment) {
	_, _ = color.New(colorAttr(text.GetColor())).Fprint(c.out, text.GetText())
}
