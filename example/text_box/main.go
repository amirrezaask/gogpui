package main

import (
	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var text1 string
	var text2 string
	var text3 string

	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		areas := gpui.ColumnarAreas(windowCtx.Area(), 3)
		g.TextBox("#text1", areas[0], frameEvents, rl.White, &text1)
		g.TextBox("#text2", areas[1], frameEvents, rl.White, &text2)
		g.TextBox("#text3", areas[2], frameEvents, rl.White, &text3)
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
