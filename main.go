package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer termbox.Close()
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
	w, h = termbox.Size()
	pageFlag := 1
	flush(8, 11, initDraw, pageFlag)
}
