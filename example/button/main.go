package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		button := &gpui.Button{Label: "Button1", FillColor: rl.White}
		buttonState := button.Render(g, rl.Rectangle{X: 200, Y: 200, Height: windowCtx.Height / 2, Width: windowCtx.Width / 2}, frameEvents).(gpui.ButtonState)
		if buttonState.Pressed {
			fmt.Println("Pressed")
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}
