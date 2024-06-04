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
			button := &gpui.Button{FillColor: rl.Red, Label: "Button " + fmt.Sprint(i)}
			button.Render(g, col, frameEvents)
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
