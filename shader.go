package orchid

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-gl/glow/gl"
)

var shaderPrograms []ShaderProgram

type ShaderProgram struct {
	programID, vertID, fragID uint32
}

type bindAttributesFunc func(ShaderProgram)

func CreateShaderProgram(vertexFilepath, fragmentFilepath string, bindAttributes bindAttributesFunc) ShaderProgram {
	var s ShaderProgram
	var err error
	s.vertID, err = loadShader(vertexFilepath, gl.VERTEX_SHADER)
	if err != nil {
		LogError(err)
	}
	s.fragID, err = loadShader(fragmentFilepath, gl.FRAGMENT_SHADER)
	if err != nil {
		LogError(err)
	}

	s.programID = gl.CreateProgram()
	gl.AttachShader(s.programID, s.vertID)
	gl.AttachShader(s.programID, s.fragID)
	gl.LinkProgram(s.programID)
	gl.ValidateProgram(s.programID)

	bindAttributes(s)

	shaderPrograms = append(shaderPrograms, s)

	return s
}

func (s ShaderProgram) Start() {
	gl.UseProgram(s.programID)
	if s.programID == 1 {
		fmt.Println("IT was!@")
	}
}

func (s ShaderProgram) Stop() {
	gl.UseProgram(0)
}

func (s ShaderProgram) Clean() {
	s.Stop()
	gl.DetachShader(s.programID, s.vertID)
	gl.DetachShader(s.programID, s.fragID)
	gl.DeleteShader(s.vertID)
	gl.DeleteShader(s.fragID)
	gl.DeleteProgram(s.programID)
}

func (s ShaderProgram) bindAttribute(attribute uint32, variableName string) {
	variableNameC, free := gl.Strs(variableName)
	defer free()
	gl.BindAttribLocation(s.programID, attribute, *variableNameC)
}

func CleanShaderPrograms() {
	for _, shader := range shaderPrograms {
		shader.Clean()
	}
}

func loadShader(filepath string, shaderType uint32) (uint32, error) {
	// b, err := ioutil.ReadFile(filepath)

	// if err != nil {
	// 	LogError(err)
	// }

	// cstring := string(b) + "\x00"

	code := ""

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		code += "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	code += "\x00"

	shaderCode, free := gl.Strs(code)
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
