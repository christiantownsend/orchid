package loader

type Texture struct {
	textureID uint32
}

func NewTexture(textureID uint32) *Texture {
	t := Texture{
		textureID: textureID,
	}

	return &t
}

func (t Texture) TextureID() uint32 {
	return t.textureID
}
