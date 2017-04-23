package orchid

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/go-gl/glow/gl"
)

type TexturedModel struct {
	model   Model
	texture *Texture
}

type Model struct {
	vaoID       uint32
	vertexCount int32
	indexCount  int32
}

type Loader struct {
	vaoIDs, vboIDs, textureIDs []uint32
}

type Texture struct {
	textureID uint32
}

// Generates a VBO and VAO which represent a model
func (l Loader) MakeModel(vertices, textureCoords []float32, indices []uint32) Model {
	// Create new VAO and bind it
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaoIDs = append(l.vaoIDs, vao)
	gl.BindVertexArray(vao)

	// Create the index buffer and bind to the current VAO
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	l.vboIDs = append(l.vboIDs, ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	// Create a VBO for the vertex positions and bind to the current VAO in the 0 index
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vboIDs = append(l.vboIDs, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil) // Store at index 0 with 3 floats per coordinate
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// Create a VBO for the texture coordinates and bind to the current VAO in the 1 index
	var texCoordVBO uint32
	gl.GenBuffers(1, &texCoordVBO)
	l.vboIDs = append(l.vboIDs, texCoordVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, texCoordVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(textureCoords)*4, gl.Ptr(textureCoords), gl.STATIC_DRAW)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil) // Store at index 1 with 2 floats per coordinate
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// Unbind the current VAO
	gl.BindVertexArray(0)

	// Create a Model object and set some data, then return
	var m Model
	m.vaoID = vao
	m.vertexCount = int32(len(vertices) / 3)
	m.indexCount = int32(len(indices))

	return m
}

func (l Loader) LoadTexture(filename string, wrapS, wrapT int32) (*Texture, error) {

	infile, err := os.Open("res/" + filename + ".PNG")
	if err != nil {
		fmt.Println(err)
	}
	defer infile.Close()

	src, err := png.Decode(infile)
	if err != nil {
		fmt.Println(err)
	}

	rgba := image.NewNRGBA(src.Bounds())
	draw.Draw(rgba, rgba.Bounds(), src, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("Unsupported Stride")
	}

	var textureID uint32
	gl.GenTextures(1, &textureID)

	l.textureIDs = append(l.textureIDs, textureID)

	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrapT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, options.TextureInterpolationMode) // minification filter
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, options.TextureInterpolationMode) // magnification filter

	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	var t Texture
	t.textureID = textureID
	return &t, nil

}

func (l Loader) Clean() {
	for _, vao := range l.vaoIDs {
		gl.DeleteVertexArrays(1, &vao)
	}

	for _, vbo := range l.vboIDs {
		gl.DeleteBuffers(1, &vbo)
	}
	for _, texture := range l.textureIDs {
		gl.DeleteTextures(1, &texture)
	}
}
