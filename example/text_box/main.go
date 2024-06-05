package main

import (
	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	lastText := ""

	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		areas := gpui.ColumnarAreas(windowCtx.Area(), 3)
		for _, area := range areas {
			_, lastText = gpui.TextBox(g, area, frameEvents, rl.White, lastText)
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
