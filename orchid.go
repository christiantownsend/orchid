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

	vertexBufferData := []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0.5, 0.5, 0}

	var indices = []uint32{
		0, 1, 3,
		3, 1, 2}

	// vertices := []float32{-0.5, 0.5, 0, -0.5, -0.5, 0, 0.5, -0.5, 0, 0.5, -0.5, 0, 0.5, 0.5, 0, -0.5, 0.5, 0}

	//model := loader.LoadToVAO(vertexBufferData, indices)
	model := loader.MakeModel(vertexBufferData, indices)

	model2 := loader.MakeModel([]float32{-1, -1, 0, -1, -.2, 0, -.5, -.5, 0, -.5, -1, 0}, []uint32{0, 3, 2, 2, 1, 0})

	staticShader, err := CreateShaderProgram("shaders/static.vert", "shaders/static.frag", staticShaderBindFunc)
	if err != nil {
		LogError(err)
	}

	for !window.ShouldClose() {
		renderer.Prepare()
		renderer.Render(model, staticShader)
		renderer.Render(model2, staticShader)
		Maintainance()
	}

	CleanShaderPrograms()
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
