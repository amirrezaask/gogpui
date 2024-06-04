package main

import (
	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		text := &gpui.Text{TextSize: 50, String: "Hello", Color: rl.White}
		text.Render(g, windowCtx.Area(), frameEvents)
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
