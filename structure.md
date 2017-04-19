


Game Assumptions:
	Movement is defined as distance per tick

	Ticks are per second

	Movement is done as: velocity * ticks * delta_time

Framework Things:
	- Game loop
	- Window manager
	- Keyboard input
	- Mouse input
	- Loading assets
	- Loading shaders
	- Renderer
		- Only render things that show up in the viewport/camera view (Quad tree?)
	- Camera system