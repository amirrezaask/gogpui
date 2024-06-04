package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textColor = rl.White
)

func handler(windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
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
	text := &gpui.Text{String: "Hello", Color: textColor}
	text.Render()
}

func main() {
	gpui.
		New(handler).
		Start()
}
