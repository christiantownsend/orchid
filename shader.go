package orchid

import (
	"io/ioutil"
)

type ShaderProgram struct {
	programID, vertID, fragID int
}

func loadShader(filepath string, shaderType int) {
	b, err := ioutil.ReadFile(filepath)

	if err != nil {
		LogError(err)
	}

	cstring := string(b)

}
