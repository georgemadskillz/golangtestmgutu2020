package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func main() {
	//initTerm()

	for {
		col := getWidth()
		fmt.Printf("col: %v\n", col)

		fmt.Printf("┌")
		col -= 2
		for i := 0; i < int(col-2); i++ {
			fmt.Printf("─")
		}
		fmt.Printf("┐")
		fmt.Printf("\n")

		fmt.Printf("│ Тестовая табличка на всю длину экрана │")
		fmt.Printf("\n")
		fmt.Printf("└───────────────────┘")
		fmt.Printf("\n")

		//clearScreen()
	}

}

var clearFuncMap map[string]func() //create a map for storing clear funcs

func initTerm() {
	clearFuncMap = make(map[string]func())
	clearFuncMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearFuncMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	clearFunc, ok := clearFuncMap[runtime.GOOS]
	if ok {
		clearFunc()
	} else {
		panic("Cannot clear screen, unsupported OS!")
	}
}
