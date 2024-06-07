package main

import (
	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		g.Text(windowCtx.Area(), frameEvents, 15, rl.White, "Hello")
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
