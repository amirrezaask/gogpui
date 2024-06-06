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
			gpui.Text("text", g, hist, frameEvents, 20, rl.White, "Hello "+fmt.Sprint(i))
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

		gpui.Button("button", g, URLArea, frameEvents, 20, rl.Red, rl.White, "URL ")
		gpui.Button("button", g, responseArea, frameEvents, 20, rl.White, rl.Red, "Response ")
		_ = codeArea

	}).
		WithFont("LiberationMono.ttf").
		Start()
}
