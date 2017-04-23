package orchid

import (
	"github.com/go-gl/glow/gl"
)

type Renderer struct{}

func (r Renderer) Prepare() {
	gl.ClearColor(0, .4, .6, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r Renderer) Render(m Model, s ShaderProgram) {
	s.Start()
	gl.BindVertexArray(m.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.DrawElements(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	//gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
	s.Stop()
}
