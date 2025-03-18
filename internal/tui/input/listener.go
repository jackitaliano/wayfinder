package input

import (
    "os"
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
