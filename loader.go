package orchid

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/go-gl/glow/gl"
)

// TexturedModel represents a raw Model and a Texture object together
type TexturedModel struct {
	model   Model
	texture *Texture
}

// Model represents the ID of the VAO associated with it, as well as the vertex and index counts used in rendering
type Model struct {
	vaoID       uint32
	vertexCount int32
	indexCount  int32
}

// Loader stores information about which VAOs, VBOs, and Textures have been loaded
// Models are created through the Loader
type Loader struct {
	vaoIDs, vboIDs, textureIDs []uint32
}

// Texture represents just the ID of the OpenGL Texture object
type Texture struct {
	textureID uint32
}

// MakeModel generates a VBO and VAO which represent a model
func (l Loader) MakeModel(vertices, textureCoords []float32, indices []uint32) Model {
	// Create new VAO and bind it
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaoIDs = append(l.vaoIDs, vao)
	gl.BindVertexArray(vao)

	// Create a Model object to return
	var m Model
	m.vaoID = vao
	m.vertexCount = int32(len(vertices) / 3)
	m.indexCount = int32(len(indices))

	// Create the index buffer and bind to the current VAO
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	l.vboIDs = append(l.vboIDs, ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	// Create VBOs for the vertices and texture coordinates and store them in the VAO
	l.StoreVBOData(m, 0, 3, vertices)
	l.StoreVBOData(m, 1, 2, textureCoords)

	// Unbind the current VAO
	gl.BindVertexArray(0)

	return m
}

// StoreVBOData will get the VAO of a Model, bind it, then store data in a VBO at index [attribute]
func (l Loader) StoreVBOData(m Model, attribute uint32, stride int32, data []float32) {
	// Get the vaoID for the Model and bind it
	vao := m.vaoID
	gl.BindVertexArray(vao)

	// Create a new VBO and store it in the current VAO at index [attribute]
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vboIDs = append(l.vboIDs, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attribute, stride, gl.FLOAT, false, 0, nil) // Store at index [attribute] with [stride] floats per coordinate
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// Unbind the vao
	gl.BindVertexArray(0)
}

// LoadTexture will take a filename of a png image and return a pointer to the Texture representation of that image
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

// Clean will delete all of the VAOs, VBOs, and Textures that have been loaded
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
