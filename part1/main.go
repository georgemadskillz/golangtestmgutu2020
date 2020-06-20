package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

func main() {
	cyclesCnt := 0
	var scr screen

	scr.init()

	reader := bufio.NewReader(os.Stdin)

	for {
		scr.buffer.Reset()
		scr.size.update()
		cols := int(scr.size.Col)
		rows := int(scr.size.Row)

		for c := 0; c < rows; c++ {

			if c == 0 {
				scr.buffer.WriteRune('┌')
				for i := 0; i < cols-2; i++ {
					scr.buffer.WriteRune('─')
				}
				scr.buffer.WriteRune('┐')
				scr.buffer.WriteRune('\n')
			} else if c == rows-1 {
				scr.buffer.WriteRune('└')
				for i := 0; i < cols-2; i++ {
					scr.buffer.WriteRune('─')
				}
				scr.buffer.WriteRune('┘')
			} else {
				var str string
				scr.buffer.WriteRune('│')
				str = fmt.Sprintf("Цикл выполнения программы #%v     нажмите Enter для перерисовки экрана..", cyclesCnt)
				scr.buffer.WriteString(str)

				for i := 0; i < cols-len([]rune(str))-2; i++ {
					scr.buffer.WriteRune(' ')
				}
				scr.buffer.WriteRune('│')
				scr.buffer.WriteRune('\n')
			}

		}

		scr.buffer.WriteTo(os.Stdout)

		_, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		} // if result == '1'

		cyclesCnt++

		//clearScreen() ???
	}

}

type screen struct {
	size       winsize
	clearFuncs map[string]func()
	buffer     bytes.Buffer
}

func (scr *screen) init() {
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
}

func (scr *screen) clear() {
	clearFunc, ok := scr.clearFuncs[runtime.GOOS]
	if ok {
		clearFunc()
	} else {
		panic("Cannot clear screen, unsupported OS!")
	}
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
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
