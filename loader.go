package orchid

import (
	"github.com/go-gl/glow/gl"
	"image/png"
	"os"
	"image"
)

type Model struct {
	vaoID       uint32
	vertexCount int32
	indexCount  int32
}

type Loader struct {
	vaoIDs, vboIDs, textureIDs []uint32
}

type Texture struct {
	glTextureID uint32
}

// Generates a VBO and VAO which represent a model
func (l Loader) MakeModel(vertices []float32, indices []uint32) Model {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	l.vboIDs = append(l.vboIDs, ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vboIDs = append(l.vboIDs, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaoIDs = append(l.vaoIDs, vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)

	var m Model
	m.vaoID = vao
	m.vertexCount = int32(len(vertices) / 3)
	m.indexCount = int32(len(indices))

	return m
}

func LoadTexture(filename string, wrapS, wrapT int32) (*Texture, error) {

	infile, err := os.Open(filename)
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

	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrapT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR) // magnification filter

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
}
