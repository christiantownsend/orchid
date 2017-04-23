package orchid

import (
	"github.com/go-gl/glow/gl"
)

type Renderer struct{}

func (r Renderer) Prepare() {
	gl.ClearColor(1, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r Renderer) Render(m Model) {
	gl.BindVertexArray(m.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.DrawElements(gl.TRIANGLES, m.vertexCount, gl.UNSIGNED_INT, nil)
	//gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
}
