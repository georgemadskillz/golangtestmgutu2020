package main

import (
	"flydb/cui"
	"flydb/ioctrl"
)

/*
Roadmap
1) интерфейс консоли
2) модуль хранения данных в оперативке
3) файловый вод-вывод

Fixes:
! брать размеры терминала с помощью пакета terminal, пока что делаем вручную системными вызовами

*/

func main() {

	go cui.UIcontroller()
	go ioctrl.IOcontroller()

	for {

	}
}
