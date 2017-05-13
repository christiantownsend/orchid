package window

import (
	"github.com/go-gl/glow/gl"
)

var (
	optionsSet bool
	Options    RunOptions
)

type RunOptions struct {
	Title string

	Width, Height int

	Fullscreen bool

	MSAA int

	TextureInterpolationMode int32
}

func NewRunOptions(title string, width int, height int, fullscreen bool, MSAA int, texIntMode int32) RunOptions {

	o := RunOptions{
		Title:      title,
		Width:      width,
		Height:     height,
		Fullscreen: fullscreen,
		MSAA:       MSAA,
		TextureInterpolationMode: texIntMode,
	}

	return o
}

func SetRunOptions(o RunOptions) {

	if o.Width == 0 {
		o.Width = 800
	}
	if o.Height == 0 {
		o.Height = 600
	}
	if o.MSAA != 0 && o.MSAA != 2 && o.MSAA != 4 && o.MSAA != 8 && o.MSAA != 16 {
		o.MSAA = 0
	}
	if o.TextureInterpolationMode != gl.LINEAR && o.TextureInterpolationMode != gl.NEAREST {
		o.TextureInterpolationMode = gl.LINEAR
	}

	Options = o
}
