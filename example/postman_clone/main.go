package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		areas := gpui.ColumnarAreas(windowCtx.Area(), 3)

		historyArea := areas[0]
		requestArea := areas[1]
		codeArea := areas[2]

		historyItemAreas := gpui.ListAreas(historyArea, 20)
		for i, hist := range historyItemAreas {
			text := &gpui.Text{TextSize: 20, Color: rl.White, String: "Hello " + fmt.Sprint(i)}
			text.Render(g, hist, frameEvents)
		}

		URLArea := rl.Rectangle{
			X:      requestArea.X,
			Y:      requestArea.Y,
			Width:  requestArea.Width,
			Height: requestArea.Height / 3,
		}
		responseArea := rl.Rectangle{
			X:      requestArea.X,
			Y:      requestArea.Y + requestArea.Height/3,
			Width:  requestArea.Width,
			Height: requestArea.Height * 2 / 3,
		}

		url := &gpui.Button{FillColor: rl.Red, Label: "URL "} //TODO: this should be a text box
		url.Render(g, URLArea, frameEvents)
		response := &gpui.Button{FillColor: rl.Red, Label: "Response "} //TODO: this should be a text box
		response.Render(g, responseArea, frameEvents)
		_ = codeArea

	}).
		WithFont("LiberationMono.ttf").
		Start()
}
