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

func (l Loader) LoadToVAO(positions []float32, indices []int) Model {
	vaoID := l.CreateVAO()
	l.BindIndexBuffer(indices)
	l.StoreDataInVBO(0, positions)
	unbindVAO()

	var m Model
	m.vaoID = vaoID
	m.vertexCount = int32(len(indices))

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

func (l Loader) CreateVAO() uint32 {
	var vaoID uint32
	gl.GenVertexArrays(1, &vaoID)
	l.vaoIDs = append(l.vaoIDs, vaoID)
	gl.BindVertexArray(vaoID)
	return vaoID
}

func (l Loader) StoreDataInVBO(attributeNumber uint32, data []float32) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	l.vboIDs = append(l.vboIDs, vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(data), gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attributeNumber, 3, gl.FLOAT, false, 0, nil) // If nil doesn't work, use gl.Ptr(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (l Loader) BindIndexBuffer(indices []int) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	l.vboIDs = append(l.vboIDs, vboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices), gl.Ptr(&indices[0]), gl.STATIC_DRAW)
}

func unbindVAO() {
	gl.BindVertexArray(0)
}
