package gogpui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderable interface {
	Render()
}

type Text struct {
	text string
}

func (t *Text) Render() {
}

type Button struct {
	Label string
}

func (b *Button) Render() {

}

type List struct {
	Items []Renderable
}

func (l *List) Render() {

}

const (
	MOD_CTRL uint8 = 1 << iota
	MOD_ALT
	MOD_SHIFT
)

const (
	EventType_KeyPress = iota + 1
	EventType_MouseClick
	EventType_ScrollWheel
	EventType_DragDrop
)
const (
	SCROLL_SPEED_Y = 5
	SCROLL_SPEED_X = 10
)

type Event struct {
	Type int
	Data any
}

type ScrollWheelEvent struct {
	V    rl.Vector2
	Mods uint8
}

type MouseClickEvent struct {
	MouseKeyCode int32
	Mods         uint8
	Position     rl.Vector2
}

type KeyPressEvent struct {
	Mods    uint8
	KeyCode int32
}

type DragDropEvent struct {
	File string
}

type WindowContext struct {
	Height int
	Width  int
}

type Handler func(windowContext WindowContext, frameEvents []Event) Renderable

type GPUI struct {
	handler Handler
}

func New(handler Handler) *GPUI {
	return &GPUI{handler: handler}
}

func pressed(key int32) bool {
	return rl.IsKeyPressed(key) || rl.IsKeyPressedRepeat(key)
}

func (g *GPUI) Start() {
	var (
		screenWidth  = int32(1280)
		screenHeight = int32(720)
	)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)

	rl.InitWindow(screenWidth, screenHeight, "TITLE")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)

	rl.SetTargetFPS(60)

	camera := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 0},
		Target:   rl.Vector2{X: 0, Y: 0},
		Rotation: 0.0,
		Zoom:     1.0,
	}
	var frameEvents []Event
	for !rl.WindowShouldClose() {
		windowCtx := WindowContext{
			Height: rl.GetScreenHeight(),
			Width:  rl.GetScreenWidth(),
		}
		ctrl := rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl)
		alt := rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt)
		shift := rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
		var mods uint8
		if ctrl {
			mods |= MOD_CTRL
		}
		if alt {
			mods |= MOD_ALT
		}
		if shift {
			mods |= MOD_SHIFT
		}
		// Keyboard
		for i := int32(0); i <= int32(350); i++ {
			lastKey := pressed(i)
			if lastKey {
				frameEvents = append(frameEvents, Event{
					Type: EventType_KeyPress,
					Data: KeyPressEvent{
						Mods:    uint8(mods),
						KeyCode: i,
					},
				})

			}
		}
		// Mouse
		for i := int32(rl.MouseButtonLeft); i <= int32(rl.MouseButtonBack); i++ {
			isPressed := rl.IsMouseButtonDown(i)
			if isPressed {
				screenPos := rl.GetMousePosition()
				worldPos := rl.GetScreenToWorld2D(screenPos, camera)
				frameEvents = append(frameEvents, Event{
					Type: EventType_MouseClick,
					Data: MouseClickEvent{
						MouseKeyCode: i,
						Position:     worldPos,
						Mods:         uint8(mods),
					},
				})
			}
		}

		// Mouse Wheel
		wheel := rl.GetMouseWheelMoveV()
		if wheel.X != 0 || wheel.Y != 0 {
			scrollY := SCROLL_SPEED_Y
			scrollX := SCROLL_SPEED_X
			if wheel.Y < 0 {
				scrollY *= -1
			}

			if wheel.X < 0 {
				scrollX *= -1
			}

			frameEvents = append(frameEvents, Event{
				Type: EventType_ScrollWheel,
				Data: ScrollWheelEvent{
					V: rl.Vector2{
						X: float32(scrollX),
						Y: float32(scrollY),
					},
					Mods: uint8(mods),
				},
			})
		}

		if rl.IsFileDropped() {
			files := rl.LoadDroppedFiles()
			for _, file := range files {
				frameEvents = append(frameEvents, Event{
					Type: EventType_DragDrop,
					Data: DragDropEvent{
						File: file,
					},
				})
			}
		}
		renderable := g.handler(windowCtx, frameEvents)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Red)
		rl.BeginMode2D(camera)
		renderable.Render()
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
