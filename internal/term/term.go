package term

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TermSpecifier string

func Startup() {
    OpenAltBuf(os.Stdin)
    HideCursor(os.Stdin)
	enableRawMode()
}

func Cleanup() {
    OpenMainBuf(os.Stdin)
    RevealCursor(os.Stdin)
	disableRawMode()
}

func GetTermSize() (int, int) {
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


