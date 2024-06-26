package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		areas := gpui.ListAreas(windowCtx.Area(), 3)
		for i, area := range areas {
			g.Button(area, frameEvents, 15, rl.Red, rl.White, "Button "+fmt.Sprint(i))
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
