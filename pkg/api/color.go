// color.go contains everything required to identify any color by its
// foreground and foreground values.
package api

import "fmt"

// -----------------------------------------------------------------------------
// Public constants: colors
// -----------------------------------------------------------------------------

// Cell colors. You can combine these with multiple attributes using
// a bitwise OR ('|'). Colors can't combine with other colors.
const (
	ColorDefault Attr = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorDarkGray
	ColorLightRed
	ColorLightGreen
	ColorLightYellow
	ColorLightBlue
	ColorLightMagenta
	ColorLightCyan
	ColorLightGray
)

// Cell attributes, it is possible to use multiple attributes by combining them
// using bitwise OR ('|'). Although, colors cannot be combined. But you can
// combine attributes and a single color.
//
// It's worth mentioning that some platforms don't support certain attributes.
// For example windows console doesn't support AttrUnderline. And on some
// terminals applying AttrBold to background may result in blinking text. Use
// them with caution and test your code on various terminals.
const (
	AttrBold Attr = 1 << (iota + 9)
	AttrBlink
	AttrHidden
	AttrDim
	AttrUnderline
	AttrCursive
	AttrReverse
	max_attr
)

// Various default colors.
var (
	ColorBlackAndWhite   = NewColor(ColorBlack, ColorWhite)
	ColorRedAndWhite     = NewColor(ColorRed, ColorWhite)
	ColorGreenAndWhite   = NewColor(ColorGreen, ColorWhite)
	ColorYellowAndWhite  = NewColor(ColorYellow, ColorWhite)
	ColorBlueAndWhite    = NewColor(ColorBlue, ColorWhite)
	ColorMagentaAndWhite = NewColor(ColorMagenta, ColorWhite)
	ColorCyanAndWhite    = NewColor(ColorCyan, ColorWhite)
)

// -----------------------------------------------------------------------------
// Public functions
// -----------------------------------------------------------------------------

// ColorToString function returns a string for the given color.
func ColorToString(color Attr) string {
	switch color {
	case ColorBlack:
		return "black"
	case ColorRed:
		return "red"
	case ColorGreen:
		return "green"
	case ColorYellow:
		return "yellow"
	case ColorBlue:
		return "blue"
	case ColorMagenta:
		return "magenta"
	case ColorCyan:
		return "cyan"
	case ColorWhite:
		return "white"
	case ColorDarkGray:
		return "darkgray"
	case ColorLightRed:
		return "lightred"
	case ColorLightGreen:
		return "lightgreen"
	case ColorLightYellow:
		return "lightyellow"
	case ColorLightBlue:
		return "lightblue"
	case ColorLightMagenta:
		return "lightmagenta"
	case ColorLightCyan:
		return "lightcyan"
	case ColorLightGray:
		return "lightgray"
	case ColorDefault:
		fallthrough
	default:
		return "default"
	}
}

// ColorToAttr function returns the color Attr from a string.
func ColorToAttr(color string) Attr {
	switch color {
	case "black":
		return ColorBlack
	case "red":
		return ColorRed
	case "green":
		return ColorGreen
	case "yellow":
		return ColorYellow
	case "blue":
		return ColorBlue
	case "magenta":
		return ColorMagenta
	case "cyan":
		return ColorCyan
	case "white":
		return ColorWhite
	case "darkgray":
		return ColorDarkGray
	case "lightred":
		return ColorLightRed
	case "lightgreen":
		return ColorLightGreen
	case "lightyellow":
		return ColorLightYellow
	case "lightblue":
		return ColorLightBlue
	case "lightmagenta":
		return ColorLightMagenta
	case "lightcyan":
		return ColorLightCyan
	case "lightgray":
		return ColorLightGray
	case "default":
		fallthrough
	default:
		return ColorDefault
	}
}

// -----------------------------------------------------------------------------
//
// Color
//
// -----------------------------------------------------------------------------

// Color structure defines any color for any element in the engine by its
// foreground and background values.
// Fg Attr is the foreground color value.
// Bg Attr is the background color value.
type Color struct {
	Fg Attr `json:"fg"`
	Bg Attr `json:"bg"`
}

// NewColor function creates a new Color instance with given foreground and
// background values.
func NewColor(fg Attr, bg Attr) *Color {
	return &Color{
		Fg: fg,
		Bg: bg,
	}
}

// ClonePoint functions creates a new Color instances with same attributes as
// the given Color.
func CloneColor(color *Color) *Color {
	return &Color{
		Fg: color.Fg,
		Bg: color.Bg,
	}
}

// NewColorFromString function creates a new Color instance with given
// foreground and background values provided as strings.
func NewColorFromString(fg string, bg string) *Color {
	return &Color{
		Fg: ColorToAttr(fg),
		Bg: ColorToAttr(bg),
	}
}

// -----------------------------------------------------------------------------
// Color public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Color instance.
func (c *Color) Clone(color *Color) {
	c.Fg = color.Fg
	c.Bg = color.Bg
}

// Set method assigns new foreground and background color values.
func (c *Color) Set(fg Attr, bg Attr) {
	c.Fg = fg
	c.Bg = bg
}

// Get method returns foreground and background values for the instace.
func (c *Color) Get() (Attr, Attr) {
	return c.Fg, c.Bg
}

// ToStringAttr method returns color information as a string with Attr values.
func (c *Color) ToStringAttr() string {
	return fmt.Sprintf("[%d:%d]", c.Fg, c.Bg)
}

// IsEqual method returns if the give Color is equal to the instance with the
// same foreground and background values.
func (c *Color) IsEqual(color *Color) bool {
	return (c.Fg == color.Fg) && (c.Bg == color.Bg)
}

// ToStringAttr method returns color information as a string with string
// values.
func (c *Color) ToString() string {
	return fmt.Sprintf("[%s:%s]", ColorToString(c.Fg), ColorToString(c.Bg))
}

// SaveToDict method saves the instance information as a map.
func (c *Color) SaveToDict() map[string]any {
	result := map[string]any{}
	result["fg"] = c.Fg
	result["bg"] = c.Bg
	return result
}
