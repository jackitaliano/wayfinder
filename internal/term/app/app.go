package app

import (
	// "fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jackitaliano/wayfinder/internal/term/buffer"
	"github.com/jackitaliano/wayfinder/internal/term/cursor"
	"github.com/jackitaliano/wayfinder/internal/term/io"
)

type TermSpecifier string

func Startup() {
    cursor.SaveCursorPos(os.Stdout)
    buffer.EnableAlternate(os.Stdout)
    io.EnableRawMode()
}

func Cleanup() {
    buffer.DisableAlternate(os.Stdout)
    io.DisableRawMode()
    cursor.RestoreCursorPos(os.Stdout)

    time.Sleep(time.Millisecond)
}

func GetSize() (int, int) {
    cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	height, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}

	return width, height
}


