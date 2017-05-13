package window

import (
	"log"
	"runtime"

	glfw "github.com/go-gl/glfw/v3.2/glfw"
	gl "github.com/go-gl/glow/gl"

	input "orchid/input"
)

func init() {
	runtime.LockOSThread()
}

// NewWindow will create a new window with a glfw context
func NewWindow() *glfw.Window {

	if !optionsSet {
		var o RunOptions
		SetRunOptions(o)
	}

	title := Options.Title
	width := Options.Width
	height := Options.Height
	fullscreen := Options.Fullscreen
	MSAA := Options.MSAA

	// Initialize glfw
	err := glfw.Init()
	log.Fatal(err)

	// Get values for video mode from the primary monitor
	monitor := glfw.GetPrimaryMonitor()

	var mode *glfw.VidMode
	if monitor != nil {
		mode = monitor.GetVideoMode()
	} else {
		mode = &glfw.VidMode{
			Width:       1,
			Height:      1,
			RedBits:     8,
			GreenBits:   8,
			BlueBits:    8,
			RefreshRate: 60,
		}
	}

	// Configure fullscreen
	if fullscreen {
		width = mode.Width
		height = mode.Height
		glfw.WindowHint(glfw.Decorated, 0)
	} else {
		monitor = nil
	}

	// Window hints
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.Samples, MSAA) // Set MSAA (antialiasing) levels

	// Create window
	window, err := glfw.CreateWindow(width, height, title, monitor, nil)
	log.Fatal(err)

	window.MakeContextCurrent()

	// Initialize OpenGL
	err = gl.Init()
	log.Fatal(err)

	gl.Enable(gl.MULTISAMPLE) // Enable MSAA

	// Set viewport size
	width, height = window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	window.SetKeyCallback(input.KeyCallbackHandler)
	window.SetMouseButtonCallback(input.MouseClickCallbackHandler)
	window.SetCursorPosCallback(input.MousePosCallbackHandler)
	window.SetScrollCallback(input.MouseScrollCallbackHandler)

	return window
}

func Maintainance(window *glfw.Window) {
	window.SwapBuffers()
	glfw.PollEvents()
}

func DestroyWindow() {
	glfw.Terminate()
}
