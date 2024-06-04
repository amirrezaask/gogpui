package gogpui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	TextSize int
	Color    color.RGBA
	String   string
}

func (t *Text) Render(g *GPUI, area rl.Rectangle, frameEvents []Event) State {
	g.DrawTextAt(t.String, t.TextSize, rl.Vector2{X: area.X, Y: area.Y}, t.Color)
	return nil
}

type Button struct {
	FillColor color.RGBA
	Label     string
}
type ButtonState struct {
	Pressed bool
}

func (b *Button) Render(g *GPUI, area rl.Rectangle, frameEvents []Event) State {
	pressed := false
	for _, evt := range frameEvents {
		if evt.Type == EventType_MouseClick {
			mouseClick := evt.Data.(MouseClickEvent)
			if mouseClick.Position.X >= area.X && mouseClick.Position.Y >= area.Y &&
				mouseClick.Position.X <= area.X+area.Width && mouseClick.Position.Y <= area.Y+area.Height {
				pressed = true
			}
		}
	}
	if pressed {
		g.DrawFilledRectangle(area, b.FillColor)
	} else {
		g.DrawRectangle(area, 2, b.FillColor)
	}

	return ButtonState{Pressed: pressed}
}

func ColumnarAreas(screenArea rl.Rectangle, columnCount int) []rl.Rectangle {
	columnWidth := screenArea.Width / float32(columnCount)
	var columns []rl.Rectangle
	for i := 0; i < columnCount; i++ {
		columns = append(columns, rl.Rectangle{
			X:      float32(i) * columnWidth,
			Width:  float32(i)*columnWidth + columnWidth,
			Y:      screenArea.Y,
			Height: screenArea.Height,
		})
	}

	return columns
}
