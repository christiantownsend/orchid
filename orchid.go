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
	options Options
	LINEAR  int32 = gl.LINEAR
	NEAREST int32 = gl.NEAREST
)

func staticShaderBindFunc(s ShaderProgram) {
	s.bindAttribute(0, "position")
	s.bindAttribute(1, "textureCoords")
}

func Run(o Options) {

	CreateWindow(o.Title, o.Width, o.Height, o.Fullscreen, o.MSAA)
	defer DestroyWindow()

	var loader Loader
	var renderer Renderer

	vertexBufferData := []float32{
		-0.7, 0.7, 0,
		-0.5, -0.5, 0,
		0.7, -0.7, 0,
		0.5, 0.5, 0}

	textureCoords := []float32{
		0, 0,
		0, 1,
		1, 1,
		1, 0}

	indices := []uint32{
		0, 1, 3,
		3, 1, 2}

	model := loader.MakeModel(vertexBufferData, textureCoords, indices)
	texture, err := loader.LoadTexture("textureTest", gl.REPEAT, gl.REPEAT)
	if err != nil {
		LogError(err)
	}

	texturedModel := TexturedModel{model, texture}

	staticShader, err := CreateShaderProgram("shaders/static.vert", "shaders/static.frag", staticShaderBindFunc)
	if err != nil {
		LogError(err)
	}

	for !window.ShouldClose() {
		renderer.Prepare()
		staticShader.Start()
		renderer.Render(texturedModel)
		staticShader.Stop()
		Maintainance()
	}

	CleanShaderPrograms()
	loader.Clean()
}

func SetRunOptions(title string, width int, height int, fullscreen bool, MSAA int, texIntMode int32) Options {
	var o Options
	o.Title = title
	o.Width = width
	o.Height = height
	o.Fullscreen = fullscreen
	o.MSAA = MSAA
	o.TextureInterpolationMode = texIntMode

	options = o

	return o
}
