package orchid

type Options struct {
	Title string

	Width, Height int

	Fullscreen bool

	MSAA int
}

func Run(o Options) {

	CreateWindow(o.Title, o.Width, o.Height, o.Fullscreen, o.MSAA)
	defer DestroyWindow()

	for !window.ShouldClose() {

		Maintainance()
	}

}

func SetRunOptions(title string, width int, height int, fullscreen bool, MSAA int) Options {
	var o Options
	o.Title = title
	o.Width = width
	o.Height = height
	o.Fullscreen = fullscreen
	o.MSAA = MSAA

	return o
}
