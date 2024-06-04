package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textColor = rl.White
	font      rl.Font
)

func handler(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
	if len(frameEvents) > 0 {
		fmt.Println("======")
		for _, evt := range frameEvents {
			fmt.Printf("event: %+v\n", evt)
		}
	}

	for _, evt := range frameEvents {
		if evt.Type == gpui.EventType_KeyPress {
			if textColor == rl.White {
				textColor = rl.Red
			} else {
				textColor = rl.White
			}
		}
	}
	text := &gpui.Text{TextSize: 50, String: "Hello", Color: textColor}
	button := &gpui.Button{Label: "Button1", FillColor: rl.White}

	text.Render(g, windowCtx.Area(), frameEvents)
	if button.Render(g, rl.Rectangle{
		X: 200, Y: 200, Height: windowCtx.Height / 2, Width: windowCtx.Width / 2,
	}, frameEvents).(gpui.ButtonState).Pressed {
		fmt.Println("Pressed the button")
	}
}

func main() {
	gpui.
		New(handler).
		WithFont("LiberationMono.ttf").
		Start()
}
