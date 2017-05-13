package loader

// Model represents the ID of the VAO associated with it, as well as the vertex and index counts used in rendering
type Model struct {
	vaoID       uint32
	vertexCount int32
	indexCount  int32
}

func (m *Model) VaoID() uint32 {
	return m.vaoID
}

func (m *Model) VertexCount() int32 {
	return m.vertexCount
}

func (m *Model) IndexCount() int32 {
	return m.indexCount
}

func NewModel(vaoID uint32, vertexCount, indexCount int32) *Model {
	m := Model{
		vaoID:       vaoID,
		vertexCount: vertexCount,
		indexCount:  indexCount,
	}

	return &m
}

// TexturedModel represents a raw Model and a Texture object together
type TexturedModel struct {
	model   *Model
	texture *Texture
}

func NewTexturedModel(model *Model, texture *Texture) *TexturedModel {
	tm := TexturedModel{
		model:   model,
		texture: texture,
	}

	return &tm
}

func (tm *TexturedModel) Model() *Model {
	return tm.model
}

func (tm *TexturedModel) Texture() *Texture {
	return tm.texture
}
