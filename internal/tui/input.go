package tui

import (
	"os"
)

type Mode string

const (
    NORMAL Mode = "NORMAL"
    INSERT Mode = "INSERT"
)

type Input struct {
    Mode Mode
    screen *Screen
}

func NewInput(screen *Screen) Input {


    return Input{
        NORMAL,
        screen,
    }

}

func (i *Input) HandleKey(inputKey byte) {
    key, handled := Keys[inputKey]

    i.screen.Buffer.StatusLine.LastInput = inputKey
    i.screen.Buffer.StatusLine.LastInputKey = string(inputKey)

    if !handled {
        i.screen.Buffer.StatusLine.LastInputMap = ""
        return
    }

    i.screen.Buffer.StatusLine.LastInputMap = key

    if i.Mode == NORMAL {
        i.HandleNormalKeys(key)
        return
    }

    if i.Mode == INSERT {
        i.HandleInsertKeys(key)
        return
    }

}

func (i *Input) HandleNormalKeys(key string) {
    if key == "j" {
        i.screen.MoveCursorDown()
    }
    if key == "k" {
        i.screen.MoveCursorUp()
    }

    if key == "h" {
        i.screen.MoveCursorLeft()
    }

    if key == "l" {
        i.screen.MoveCursorRight()
    }

    if key == "RET" {
        i.screen.Buffer.MoveCursorDown()
    }

    if key == "i" {
        i.Mode = INSERT
        i.screen.Buffer.CursorInsertMode()
    }

    if key == "a" {
        i.Mode = INSERT
        i.screen.Buffer.CursorAppendMode()
    }

    if key == "A" {
        i.Mode = INSERT
        i.screen.Buffer.CursorEnd()
        i.screen.Buffer.CursorAppendMode()
    }

    if key == "I" {
        i.Mode = INSERT
        i.screen.Buffer.CursorHome()
        i.screen.Buffer.CursorInsertMode()
    }

    if key == "o" {
        i.Mode = INSERT
        i.screen.Buffer.AppendLineBelow()
        i.screen.Buffer.CursorInsertMode()
    }

    if key == "O" {
        i.Mode = INSERT
        i.screen.Buffer.AppendLineAbove()
        i.screen.Buffer.CursorInsertMode()
    }

    if key == "D" {
        i.screen.Buffer.DeleteToEnd()
    }

    if key == "x" {
        i.screen.Buffer.DeleteChar()
    }

    if key == "0" {
        i.screen.Buffer.CursorHome()
    }

    if key == "$" {
        i.screen.Buffer.CursorEnd()
    }
}

func (i *Input) HandleInsertKeys(key string) {
    if key == "ESC" {
        i.Mode = NORMAL
        i.screen.Buffer.CursorNormalMode()
        return
    }

    if key == "LF" {
        i.screen.Buffer.CarryLine()
        return
    }

    if key == "DEL" {
        i.screen.Buffer.Backspace()
        return
    }

    if key == "TAB" {
        i.screen.InsertChar(" ")
        return
    }

    i.screen.InsertChar(key)
}

func ListenForKeys(keyChan chan byte) {
    inputBuffer := make([]byte, 1)
    go func() {
        for {
            os.Stdin.Read(inputBuffer)
            keyChan <- inputBuffer[0]
        }
    }()
}
