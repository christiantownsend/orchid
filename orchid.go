package orchid

import "github.com/go-gl/glow/gl"

type Options struct {
	Title string

	Width, Height int

	Fullscreen bool

	MSAA int

	TextureInterpolationMode int32
}

var (
	optionsSet bool
	options    Options
	LINEAR     int32 = gl.LINEAR
	NEAREST    int32 = gl.NEAREST
)

func CreateRunOptions(title string, width int, height int, fullscreen bool, MSAA int, texIntMode int32) Options {
	var o Options
	o.Title = title
	o.Width = width
	o.Height = height
	o.Fullscreen = fullscreen
	o.MSAA = MSAA
	o.TextureInterpolationMode = texIntMode

	return o
}

func SetRunOptions(o Options) {

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

	options = o
	optionsSet = true
}
