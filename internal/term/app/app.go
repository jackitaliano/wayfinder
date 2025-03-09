package app

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

    "github.com/jackitaliano/wayfinder/internal/term/buffer"
    "github.com/jackitaliano/wayfinder/internal/term/io"
)

type TermSpecifier string

func Startup() {
    buffer.EnableAlternate(os.Stdin)
    io.EnableRawMode()
}

func Cleanup() {
    buffer.DisableAlternate(os.Stdin)
    io.DisableRawMode()
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


