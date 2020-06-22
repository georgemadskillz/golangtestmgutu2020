package cui

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

// Screen is handling terminal window output
type Screen struct {
	size       winsize
	clearFuncs map[string]func()
	buffer     bytes.Buffer
}

// Winsize stores terminal window sizes
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Init Screen params
func (scr *Screen) Init() {
	// Update window size
	scr.size.update()

	// Set clear screen system calls functions
	scr.clearFuncs = make(map[string]func())
	scr.clearFuncs["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	scr.clearFuncs["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	scr.Clear()
}

// Clear screen by system call from clearFunc map
func (scr *Screen) Clear() {
	clearFunc, ok := scr.clearFuncs[runtime.GOOS]
	if ok {
		clearFunc()
	} else {
		panic("Cannot clear screen, unsupported OS!")
	}
}

// Draw screen by writing screen buffer to Stdout
func (scr *Screen) Draw() {
	fmt.Printf("SCREEN BUFFER: %v\n", scr.buffer)
	//scr.buffer.WriteTo(os.Stdout)
}

func (ws *winsize) update() {

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
}

// GetWidth returns number of columns of terminal window
func (scr *Screen) GetWidth() uint16 {
	scr.size.update()
	return scr.size.Col
}

// GetHeight returns number of rows of terminal window
func (scr *Screen) GetHeight() uint16 {
	scr.size.update()
	return scr.size.Row
}

// GetBuff returns pointer to the screen buffer
func (scr *Screen) GetBuff() *bytes.Buffer {
	return &scr.buffer
}
