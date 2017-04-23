package orchid

type Options struct {
	Title string

	Width, Height int

	Fullscreen bool

	MSAA int
}

func staticShaderBindFunc(s ShaderProgram) {
	s.bindAttribute(0, "position")
}

func Run(o Options) {

	CreateWindow(o.Title, o.Width, o.Height, o.Fullscreen, o.MSAA)
	defer DestroyWindow()

	var loader Loader
	var renderer Renderer

	// vertexBufferData := []float32{
	// 	-0.5, 0.5, 0,
	// 	-0.5, -0.5, 0,
	// 	0.5, -0.5, 0,
	// 	0.5, 0.5, 0}

	// var indices = []int{
	// 	0, 1, 3,
	// 	3, 1, 2}

	vertices := []float32{-0.5, 0.5, 0, -0.5, -0.5, 0, 0.5, -0.5, 0, 0.5, -0.5, 0, 0.5, 0.5, 0, -0.5, 0.5, 0}

	//model := loader.LoadToVAO(vertexBufferData, indices)
	model := loader.MakeModel(vertices)

	staticShader, err := CreateShaderProgram("shaders/static.vert", "shaders/static.frag", staticShaderBindFunc)
	if err != nil {
		LogError(err)
	}

	for !window.ShouldClose() {
		renderer.Prepare()
		staticShader.Start()
		renderer.Render(model)
		staticShader.Stop()
		Maintainance()
	}

	staticShader.Clean()
	loader.Clean()
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
