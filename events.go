package gogpui

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MOD_CTRL uint8 = 1 << iota
	MOD_ALT
	MOD_SHIFT
)

const (
	EventType_KeyPress = iota + 1
	EventType_KeyPressRepeat
	EventType_MouseClick
	EventType_ScrollWheel
	EventType_DragDrop
)
const (
	SCROLL_SPEED_Y = 5
	SCROLL_SPEED_X = 10
)

type Event struct {
	Type int
	Data any
}

type ScrollWheelEvent struct {
	V    rl.Vector2
	Mods uint8
}

type MouseClickEvent struct {
	MouseKeyCode int32
	Mods         uint8
	Position     rl.Vector2
}

type KeyPressEvent struct {
	Mods    uint8
	KeyCode int32
}

func (k KeyPressEvent) GetAsciiChar() byte {
	return 'A'
}

type KeyPressRepeatEvent struct {
	Mods    uint8
	KeyCode int32
}

type DragDropEvent struct {
	File string
}
