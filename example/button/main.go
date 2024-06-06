package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		pressed := gpui.Button(g,
			gpui.Rectangle{X: 200, Y: 200, Height: windowCtx.Height / 2, Width: windowCtx.Width / 2},
			frameEvents,
			15,
			rl.White,
			rl.White,
			"Button1",
		)
		if pressed {
			fmt.Println("Pressed")
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
