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
		gpui.TextBox("textbox1", g, areas[0], frameEvents, rl.White, &text1)
		gpui.TextBox("textbox2", g, areas[1], frameEvents, rl.White, &text2)
		gpui.TextBox("textbox3", g, areas[2], frameEvents, rl.White, &text3)
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
