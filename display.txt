/*
To make the console not render behind the main window:
	go build -ldflags -H=windowsgui
*/

package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/glow/gl"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func main() {

	// Initialize glfw
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Window hints
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create window
	width := 800
	height := 600

	window, err := glfw.CreateWindow(width, height, "OpenGL Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize OpenGL
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	// Set viewport size
	width, height = window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	// Set key callback
	window.SetKeyCallback(keyCallback)

	// Set OpenGL defaults
	gl.ClearColor(1.0, 0, 1.0, 1.0)

	// Loop until window closes
	for !window.ShouldClose() {
		// Rendering
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Maintainance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var keyCallback glfw.KeyCallback = func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Printf("Input: %v\n", key)
	switch key {
	case glfw.KeyEscape:
		w.Destroy()
	}
}
