package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Controller struct {
	inputChannel chan *InputMessage
}

var (
	KeyMsg     = 0
	ClickMsg   = 1
	controller *Controller
	mouse      Mouse
)

type InputMessage struct {
	messageType int
	keyEvent    *KeyEvent
	mouseEvent  *MouseClickEvent
}

type KeyEvent struct {
	key      int
	action   int
	modifier int
}

type Mouse struct {
	x, y int
}

type MouseClickEvent struct {
	x, y     int
	buttonID int
	action   int
	modifier int
}

func CreateController(eventBufferSize int) *Controller {
	inputChannel := make(chan *InputMessage, eventBufferSize)

	controller = &Controller{inputChannel}

	return controller

}

func KeyCallbackHandler(w *glfw.Window, key glfw.Key, scancode int, actions glfw.Action, mods glfw.ModifierKey) {

	ke := KeyEvent{int(key), int(actions), int(mods)}

	im := InputMessage{KeyMsg, &ke, nil}

	controller.inputChannel <- &im
}

func MousePosCallbackHandler(w *glfw.Window, x, y float64) {

	width, height := w.GetSize()

	tx := (x + 1.0) / 2.0 * float64(width)  // Convert x into pixel coordinates based on size of screen
	ty := (y + 1.0) / 2.0 * float64(height) // Convert y into pixel coordinates based on size of screen

	mouse.x = int(tx)
	mouse.y = int(ty)

}

func MouseClickCallbackHandler(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	click := MouseClickEvent{mouse.x, mouse.y, int(button), int(action), int(mods)}

	message := InputMessage{ClickMsg, nil, &click}

	controller.inputChannel <- &message
}

func MouseScrollCallbackHandler(w *glfw.Window, xoffset, yoffset float64) {

}

func GetMousePos() (int, int) {
	return mouse.x, mouse.y
}

// inputchannel <-
// inputchannel <- nil

// game.step() {
// 	timelasttick = atime
// 	for input <- inputChannel != nil {
// 		Process events
// 		if curtime - lastime > ticktime
// 			break
// 	}
// 	update enemies
// 	do attacks
// 	update position

// 	movex * dtime
// }
