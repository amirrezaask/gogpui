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
	switch {
	case k.KeyCode == rl.KeySpace:
		return ' '
	case k.KeyCode == rl.KeyEnter:
		return '\n'
	case k.KeyCode == rl.KeyTab:
		return '\t'
	case k.KeyCode == rl.KeyKpDivide:
		return '/'
	case k.KeyCode == rl.KeyKpMultiply:
		return '*'
	case k.KeyCode == rl.KeyKpSubtract:
		return '-'
	case k.KeyCode == rl.KeyKpAdd:
		return '+'
	case k.KeyCode == rl.KeyKpEnter:
		return '\n'
	case k.KeyCode == rl.KeyKpEqual:
		return '='
	case k.KeyCode == rl.KeyApostrophe:
		return '`'
	case k.KeyCode == rl.KeyComma:
		return ','
	case k.KeyCode == rl.KeyMinus:
		return '-'
	case k.KeyCode == rl.KeyPeriod:
		return '.'
	case k.KeyCode == rl.KeySlash:
		return '/'
	case k.KeyCode == rl.KeyZero:
		return '0'
	case k.KeyCode == rl.KeyOne:
		return '1'
	case k.KeyCode == rl.KeyTwo:
		return '2'
	case k.KeyCode == rl.KeyThree:
		return '3'
	case k.KeyCode == rl.KeyFour:
		return '4'
	case k.KeyCode == rl.KeyFive:
		return '5'
	case k.KeyCode == rl.KeySix:
		return '6'
	case k.KeyCode == rl.KeySeven:
		return '7'
	case k.KeyCode == rl.KeyEight:
		return '8'
	case k.KeyCode == rl.KeyNine:
		return '9'
	case k.KeyCode == rl.KeySemicolon:
		return ';'
	case k.KeyCode == rl.KeyEqual:
		return '='
	case k.KeyCode == rl.KeyA:
		return 'a'
	case k.KeyCode == rl.KeyB:
		return 'b'
	case k.KeyCode == rl.KeyC:
		return 'c'
	case k.KeyCode == rl.KeyD:
		return 'd'
	case k.KeyCode == rl.KeyE:
		return 'e'
	case k.KeyCode == rl.KeyF:
		return 'f'
	case k.KeyCode == rl.KeyG:
		return 'g'
	case k.KeyCode == rl.KeyH:
		return 'h'
	case k.KeyCode == rl.KeyI:
		return 'i'
	case k.KeyCode == rl.KeyJ:
		return 'j'
	case k.KeyCode == rl.KeyK:
		return 'k'
	case k.KeyCode == rl.KeyL:
		return 'l'
	case k.KeyCode == rl.KeyM:
		return 'm'
	case k.KeyCode == rl.KeyN:
		return 'n'
	case k.KeyCode == rl.KeyO:
		return 'o'
	case k.KeyCode == rl.KeyP:
		return 'p'
	case k.KeyCode == rl.KeyQ:
		return 'q'
	case k.KeyCode == rl.KeyR:
		return 'r'
	case k.KeyCode == rl.KeyS:
		return 's'
	case k.KeyCode == rl.KeyT:
		return 't'
	case k.KeyCode == rl.KeyU:
		return 'u'
	case k.KeyCode == rl.KeyV:
		return 'v'
	case k.KeyCode == rl.KeyW:
		return 'w'
	case k.KeyCode == rl.KeyX:
		return 'x'
	case k.KeyCode == rl.KeyY:
		return 'y'
	case k.KeyCode == rl.KeyZ:
		return 'z'
	default:
		return 0
	}
}

type KeyPressRepeatEvent struct {
	Mods    uint8
	KeyCode int32
}

type DragDropEvent struct {
	File string
}
