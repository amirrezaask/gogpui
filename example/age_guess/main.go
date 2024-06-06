package main

import (
	"fmt"
	"strconv"
	"time"

	gpui "github.com/amirrezaask/gogpui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var birth string
	gpui.New(func(g *gpui.GPUI, windowCtx gpui.WindowContext, frameEvents []gpui.Event) {
		windowArea := windowCtx.Area()

		gpui.TextBox("textbox", g, gpui.Rectangle{
			X:      windowArea.X + (windowArea.Width)*1/4,
			Y:      windowArea.Y + (windowArea.Height)*1/2,
			Width:  windowArea.Width * 1 / 5,
			Height: windowArea.Height * 1 / 10,
		}, frameEvents, rl.Red, &birth)

		pressed := gpui.Button("button",
			g,
			gpui.Rectangle{
				X:      windowArea.X + (windowArea.Width)*1/4,
				Y:      windowArea.Y + (windowArea.Height)*3/4,
				Width:  windowArea.Width * 1 / 5,
				Height: windowArea.Height * 1 / 10,
			},
			frameEvents,
			15,
			rl.White,
			rl.White,
			"Guess",
		)
		if pressed {
			birthYear, err := strconv.ParseInt(birth, 10, 64)
			if err != nil {
				fmt.Println("enter valid year")
			} else {
				fmt.Println("your age is ", time.Now().Year()-int(birthYear)+1)
			}
		}
	}).
		WithFont("LiberationMono.ttf").
		Start()
}

