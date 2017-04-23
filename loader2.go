package orchid

import (
	"github.com/go-gl/glow/gl"
)

type Model struct {
	vaoID       uint32
	vertexCount int32
}

type Loader struct {
	vaoIDs, vboIDs []uint32
}

// Generates a VBO and VAO which represent a model
func (l Loader) MakeModel(vertices []float32) Model {
	// var ibo uint32
	// gl.GenBuffers(1, &ibo)
	// l.vboIDs = append(l.vboIDs, ibo)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(&indices[0]), gl.STATIC_DRAW)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vboIDs = append(l.vboIDs, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaoIDs = append(l.vaoIDs, vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)

	var m Model
	m.vaoID = vao
	m.vertexCount = int32(len(vertices) / 3)

	return m
}

func (l Loader) Clean() {
	for _, vao := range l.vaoIDs {
		gl.DeleteVertexArrays(1, &vao)
	}

	for _, vbo := range l.vboIDs {
		gl.DeleteBuffers(1, &vbo)
	}
}
