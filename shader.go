package orchid

import (
	"io/ioutil"

	"fmt"

	"github.com/go-gl/glow/gl"
)

type ShaderProgram struct {
	programID, vertID, fragID int
}

func loadShader(filepath string, shaderType uint32) (uint32, error) {
	b, err := ioutil.ReadFile(filepath)

	if err != nil {
		LogError(err)
	}

	cstring := string(b) + "\x00"

	shaderCode, free := gl.Strs(cstring)
	defer free()

	shaderID := gl.CreateShader(shaderType)
	gl.ShaderSource(shaderID, 1, shaderCode, nil)
	gl.CompileShader(shaderID)
	var status int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, fmt.Errorf("Could not compile shader %s", filepath)
	}

	return shaderID, nil

}
