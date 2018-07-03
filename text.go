package dfhackrpc

import (
	"io"
	"os"

	"github.com/BenLubar/dfhackrpc/dfproto/core"
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
	dfproto_core.CoreTextFragment_COLOR_BLACK:        color.FgBlack,
	dfproto_core.CoreTextFragment_COLOR_BLUE:         color.FgBlue,
	dfproto_core.CoreTextFragment_COLOR_GREEN:        color.FgGreen,
	dfproto_core.CoreTextFragment_COLOR_CYAN:         color.FgCyan,
	dfproto_core.CoreTextFragment_COLOR_RED:          color.FgRed,
	dfproto_core.CoreTextFragment_COLOR_MAGENTA:      color.FgMagenta,
	dfproto_core.CoreTextFragment_COLOR_BROWN:        color.FgYellow,
	dfproto_core.CoreTextFragment_COLOR_GREY:         color.FgWhite,
	dfproto_core.CoreTextFragment_COLOR_DARKGREY:     color.FgHiBlack,
	dfproto_core.CoreTextFragment_COLOR_LIGHTBLUE:    color.FgHiBlue,
	dfproto_core.CoreTextFragment_COLOR_LIGHTGREEN:   color.FgHiGreen,
	dfproto_core.CoreTextFragment_COLOR_LIGHTCYAN:    color.FgHiCyan,
	dfproto_core.CoreTextFragment_COLOR_LIGHTRED:     color.FgHiRed,
	dfproto_core.CoreTextFragment_COLOR_LIGHTMAGENTA: color.FgHiMagenta,
	dfproto_core.CoreTextFragment_COLOR_YELLOW:       color.FgHiYellow,
	dfproto_core.CoreTextFragment_COLOR_WHITE:        color.FgHiWhite,
}

func colorAttr(c dfproto_core.CoreTextFragment_Color) color.Attribute {
	if c >= 0 && int(c) < len(colors) {
		return colors[c]
	}
	return color.Reset
}

func (c *Client) outputTextFragment(text *dfproto_core.CoreTextFragment) {
	_, _ = color.New(colorAttr(text.GetColor())).Fprint(c.out, text.GetText())
}
