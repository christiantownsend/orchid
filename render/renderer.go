package render

import (
	"github.com/go-gl/glow/gl"

	loader "orchid/loader"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	var r Renderer

	return &r
}

func (r Renderer) Prepare() {
	gl.ClearColor(0, .4, .6, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r Renderer) Render(tm *loader.TexturedModel) {
	m := tm.Model()
	t := tm.Texture()

	gl.BindVertexArray(m.VaoID())
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.TextureID())
	gl.DrawElements(gl.TRIANGLES, m.IndexCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.BindVertexArray(0)
}
