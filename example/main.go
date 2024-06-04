package main

import (
	"fmt"

	gpui "github.com/amirrezaask/gogpui"
)

func handler(windowCtx gpui.WindowContext, frameEvents []gpui.Event) gpui.Renderable {
	if len(frameEvents) > 0 {
		fmt.Println("======")
		for _, evt := range frameEvents {
			fmt.Printf("event: %+v\n", evt)
		}
	}
	return &gpui.Button{}
}

func main() {
	gpui.
		New(handler).
		Start()
}
