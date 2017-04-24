package orchid

import (
	"github.com/go-gl/glow/gl"
)

type Renderer struct{}

func (r Renderer) Prepare() {
	gl.ClearColor(0, .4, .6, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r Renderer) Render(tm TexturedModel) {
	m := tm.Model
	t := tm.Texture

	gl.BindVertexArray(m.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.textureID)
	gl.DrawElements(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	//gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.BindVertexArray(0)
}
