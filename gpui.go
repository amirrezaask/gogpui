package gogpui

import (
	"image/color"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowContext struct {
	Height float32
	Width  float32
}

func (w *WindowContext) Area() rl.Rectangle {
	return rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(w.Width),
		Height: float32(w.Height),
	}
}

type Handler func(g *GPUI, windowContext WindowContext, frameEvents []Event)

type GPUI struct {
	fontName string
	fontData []byte
	fontMap  map[int]rl.Font
	handler  Handler
}

func (g *GPUI) DrawTextAt(text string, fontsize int, at rl.Vector2, c color.RGBA) {
	rl.DrawTextEx(g.GetFont(fontsize), text, rl.Vector2{X: at.X, Y: at.Y}, float32(fontsize), 0, c)
}

func (g *GPUI) DrawRectangle(area rl.Rectangle, lineThick int, borderColor color.RGBA) {
	rl.DrawRectangleLinesEx(area, 2, borderColor)
}

func (g *GPUI) DrawFilledRectangle(area rl.Rectangle, fillColor color.RGBA) {
	rl.DrawRectangleRec(area, fillColor)
}

func (g *GPUI) GetFont(fontSize int) rl.Font {
	if font, exists := g.fontMap[fontSize]; exists {
		return font
	}

	g.fontMap[fontSize] = rl.LoadFontEx(g.fontName, int32(fontSize), nil)
	return g.fontMap[fontSize]
}

type Option func(*GPUI)

func (g *GPUI) WithFont(font string) *GPUI {
	g.fontName = font
	return g
}

func New(handler Handler) *GPUI {
	return &GPUI{handler: handler, fontMap: make(map[int]rl.Font)}
}

func (g *GPUI) Start() {
	var (
		screenWidth  = int32(1280)
		screenHeight = int32(720)
	)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)

	rl.InitWindow(screenWidth, screenHeight, "TITLE")
	defer rl.CloseWindow()

	fontData, err := os.ReadFile(g.fontName)
	if err != nil {
		panic(err)
	}
	g.fontData = fontData

	rl.SetExitKey(rl.KeyNull)

	rl.SetTargetFPS(60)

	camera := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 0},
		Target:   rl.Vector2{X: 0, Y: 0},
		Rotation: 0.0,
		Zoom:     1.0,
	}
	for !rl.WindowShouldClose() {
		var frameEvents []Event
		windowCtx := WindowContext{
			Height: float32(rl.GetScreenHeight()),
			Width:  float32(rl.GetScreenWidth()),
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
			lastKey := rl.IsKeyPressed(i)
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
		for i := int32(0); i <= int32(350); i++ {
			lastKey := rl.IsKeyPressedRepeat(i)
			if lastKey {
				frameEvents = append(frameEvents, Event{
					Type: EventType_KeyPressRepeat,
					Data: KeyPressRepeatEvent{
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

		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.Black)
		g.handler(g, windowCtx, frameEvents)
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
