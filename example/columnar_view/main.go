package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		columns := gpui.ColumnarAreas(windowCtx.Area(), 3)
		for i, col := range columns {
			gpui.Button(g, col, frameEvents, 15, rl.Red, rl.White, "Button "+fmt.Sprint(i))
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
