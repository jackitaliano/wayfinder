package io

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// Platform-specific ioctl constants
const (
	TCGETS  = 0x5401 // Linux
	TCSETS  = 0x5402
	TIOCGETA = 0x40487413 // macOS/BSD
	TIOCSETA = 0x80487414
)

// Enable raw mode (disable line buffering)
func EnableRawMode() {
	fd := int(os.Stdin.Fd())
	termios, err := getTermios(fd)
	if err != nil {
		fmt.Println("Error getting terminal attributes:", err)
		os.Exit(1)
	}

	// Modify termios to disable canonical mode & echo
	newTermios := *termios
	newTermios.Lflag &^= syscall.ICANON | syscall.ECHO

	err = setTermios(fd, &newTermios)
	if err != nil {
		fmt.Println("Error setting terminal attributes:", err)
		os.Exit(1)
	}
}

// Disable raw mode (restore normal terminal behavior)
func DisableRawMode() {
	fd := int(os.Stdin.Fd())
	termios, err := getTermios(fd)
	if err != nil {
		fmt.Println("Error getting terminal attributes:", err)
		os.Exit(1)
	}

	// Restore original settings
	err = setTermios(fd, termios)
	if err != nil {
		fmt.Println("Error restoring terminal attributes:", err)
		os.Exit(1)
	}
}

// getTermios fetches the terminal attributes based on the platform
func getTermios(fd int) (*syscall.Termios, error) {
	var termios syscall.Termios
	var req uintptr
	if isMacOS() {
		req = TIOCGETA
	} else {
		req = TCGETS
	}

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), req, uintptr(unsafe.Pointer(&termios)))
	if err != 0 {
		return nil, err
	}
	return &termios, nil
}

// setTermios sets the terminal attributes based on the platform
func setTermios(fd int, termios *syscall.Termios) error {
	var req uintptr
	if isMacOS() {
		req = TIOCSETA
	} else {
		req = TCSETS
	}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), req, uintptr(unsafe.Pointer(termios)))
	if err != 0 {
		return err
	}
	return nil
}

// Detect if the OS is macOS
func isMacOS() bool {
	return runtime.GOOS == "darwin" || runtime.GOOS == "macos"
}

