package gogpui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Text(g *GPUI, area rl.Rectangle, frameEvents []Event, TextSize int, Color color.RGBA, String string) {
	g.DrawTextAt(String, TextSize, rl.Vector2{X: area.X, Y: area.Y}, Color)
}

func Button(g *GPUI,
	area rl.Rectangle,
	frameEvents []Event,
	TextSize int,
	FillColor color.RGBA,
	TextColor color.RGBA,
	Label string) bool {
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
		g.DrawFilledRectangle(area, FillColor)
	} else {
		g.DrawRectangle(area, 2, FillColor)
	}

	g.DrawTextAt(Label, TextSize, rl.Vector2{
		X: area.X + area.Width*2/3,
		Y: area.Y + area.Height/2,
	}, TextColor)

	return pressed
}
func Border(screenArea rl.Rectangle) rl.Rectangle {
	return rl.Rectangle{
		X:      screenArea.X + 1,
		Y:      screenArea.Y + 1,
		Width:  screenArea.Width - 1,
		Height: screenArea.Height - 1,
	}
}

func TextBox(g *GPUI, screanArea rl.Rectangle, frameEvents []Event, TextColor color.RGBA, lastText string) (hasFocus bool, text string) {
	text = lastText
	for _, evt := range frameEvents {
		if evt.Type == EventType_MouseClick {
			mouseClick := evt.Data.(MouseClickEvent)
			if mouseClick.Position.X >= screanArea.X && mouseClick.Position.Y >= screanArea.Y &&
				mouseClick.Position.X <= screanArea.X+screanArea.Width && mouseClick.Position.Y <= screanArea.Y+screanArea.Height {
				hasFocus = true
			}
		} else if evt.Type == EventType_KeyPress {
			char := evt.Data.(KeyPressEvent).GetAsciiChar()
			if char != 0 {
				text += string(char)
			}
		}
	}

	g.DrawRectangle(screanArea, 1, rl.Red)
	g.DrawTextAt(text, 40, rl.Vector2{X: screanArea.X, Y: screanArea.Y}, rl.White)

	return hasFocus, text
}

func ColumnarAreas(screenArea rl.Rectangle, columnCount int) []rl.Rectangle {
	columnWidth := screenArea.Width / float32(columnCount)
	var columns []rl.Rectangle
	for i := 0; i < columnCount; i++ {
		columns = append(columns, rl.Rectangle{
			X:      float32(i)*columnWidth + screenArea.X,
			Width:  columnWidth,
			Y:      screenArea.Y,
			Height: screenArea.Height,
		})
	}

	return columns
}

func ListAreas(screenArea rl.Rectangle, listItemCount int) []rl.Rectangle {
	itemHeight := screenArea.Height / float32(listItemCount)
	var items []rl.Rectangle
	for i := 0; i < listItemCount; i++ {
		items = append(items, rl.Rectangle{
			X:      screenArea.X,
			Width:  screenArea.Width,
			Y:      float32(i)*itemHeight + screenArea.Y,
			Height: itemHeight,
		})
	}

	return items
}

func GridAreas(screenArea rl.Rectangle, rowsCount int, columnsCount int) []rl.Rectangle {
	var items []rl.Rectangle
	columns := ColumnarAreas(screenArea, columnsCount)
	for _, col := range columns {
		listAreas := ListAreas(col, rowsCount)
		items = append(items, listAreas...)
	}
	return items
}
