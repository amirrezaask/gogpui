package gogpui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GPUI) Text(area Rectangle, frameEvents []Event, TextSize int, Color color.RGBA, String string) {
	g.DrawTextAt(String, TextSize, Vector2{X: area.X, Y: area.Y}, Color)
}

func (g *GPUI) Button(
	area Rectangle,
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

	g.DrawTextAt(Label, TextSize, Vector2{
		X: area.X + area.Width*2/3,
		Y: area.Y + area.Height/2,
	}, TextColor)

	return pressed
}

func (g *GPUI) Border(screenArea Rectangle) Rectangle {
	return rl.Rectangle{
		X:      screenArea.X + 1,
		Y:      screenArea.Y + 1,
		Width:  screenArea.Width - 1,
		Height: screenArea.Height - 1,
	}
}

func (g *GPUI) TextBox(
	id string,
	screanArea rl.Rectangle,
	frameEvents []Event,
	TextColor color.RGBA,
	textPointer *string) {
	text := *textPointer
	for _, evt := range frameEvents {
		if evt.Type == EventType_MouseClick {
			mouseClick := evt.Data.(MouseClickEvent)
			if mouseClick.Position.X >= screanArea.X && mouseClick.Position.Y >= screanArea.Y &&
				mouseClick.Position.X <= screanArea.X+screanArea.Width && mouseClick.Position.Y <= screanArea.Y+screanArea.Height {
				g.activeTextBoxID = id
			}
		}
	}

	if g.activeTextBoxID == id {
		for _, evt := range frameEvents {
			if evt.Type == EventType_KeyPress {
				char := evt.Data.(KeyPressEvent).GetAsciiChar()
				if char != 0 {
					text += string(char)
				}
			}
		}

		*textPointer = text
	}

	border := rl.Red
	if g.activeTextBoxID == id {
		border = rl.Blue
	}
	g.DrawRectangle(screanArea, 1, border)
	g.DrawTextAt(*textPointer, 40, Vector2{X: screanArea.X, Y: screanArea.Y}, rl.White)
}

func ColumnarAreas(screenArea Rectangle, columnCount int) []Rectangle {
	columnWidth := screenArea.Width / float32(columnCount)
	var columns []rl.Rectangle
	for i := 0; i < columnCount; i++ {
		columns = append(columns, Rectangle{
			X:      float32(i)*columnWidth + screenArea.X,
			Width:  columnWidth,
			Y:      screenArea.Y,
			Height: screenArea.Height,
		})
	}

	return columns
}

func ListAreas(screenArea Rectangle, listItemCount int) []Rectangle {
	itemHeight := screenArea.Height / float32(listItemCount)
	var items []Rectangle
	for i := 0; i < listItemCount; i++ {
		items = append(items, Rectangle{
			X:      screenArea.X,
			Width:  screenArea.Width,
			Y:      float32(i)*itemHeight + screenArea.Y,
			Height: itemHeight,
		})
	}

	return items
}

func GridAreas(screenArea Rectangle, rowsCount int, columnsCount int) []Rectangle {
	var items []rl.Rectangle
	columns := ColumnarAreas(screenArea, columnsCount)
	for _, col := range columns {
		listAreas := ListAreas(col, rowsCount)
		items = append(items, listAreas...)
	}
	return items
}
