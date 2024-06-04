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
			listAreas := gpui.ListAreas(col, 3)
			for j, lis := range listAreas {
				button := &gpui.Button{FillColor: rl.Red, Label: "Button " + fmt.Sprint(i*j)}
				button.Render(g, lis, frameEvents)
			}
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
