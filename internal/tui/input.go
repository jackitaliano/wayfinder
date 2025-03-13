package tui

import (
	"os"
)

type Mode string

const (
    NORMAL Mode = "NORMAL"
    INSERT Mode = "INSERT"
)

func ListenForKeys(keyChan chan byte) {
    inputBuffer := make([]byte, 1)
    go func() {
        for {
            os.Stdin.Read(inputBuffer)
            keyChan <- inputBuffer[0]
        }
    }()
}
