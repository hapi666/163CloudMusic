package main

import (
	"fmt"
	"reflect"

	"github.com/hapi666/163CloudMusic/crawl"

	"github.com/tidwall/gjson"

	"github.com/nsf/termbox-go"
)

var (
	eventQueue = make(chan termbox.Event)
	//moveX int
	moveY                   int
	oneLineWord             = []rune{'网', '易', '云', '音', '乐', '欢', '迎', '您'}
	twoLineWord             = []rune{'1', '.', '获', '取', '榜', '单', '.'}
	threeLineWord           = []rune{'2', '.', '获', '取', '歌', '曲', '热', '评', '.'}
	newSongList             = []rune{'1', '.', '云', '音', '乐', '新', '歌', '榜', '.'}
	hotSongList             = []rune{'2', '.', '云', '音', '乐', '热', '歌', '榜', '.'}
	neteaseOriginalSongList = []rune{'3', '.', '网', '易', '原', '创', '歌', '曲', '榜', '.'}
	cloudMusicSoared        = []rune{'4', '.', '云', '音', '乐', '飙', '升', '榜', '.'}
)

func initDraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 30+2*len(oneLineWord); i += 2 {
		termbox.SetCell(i, 5, oneLineWord[j], termbox.ColorDefault, termbox.ColorBlue)
		j++
	}

	termbox.SetCell(28, moveY, '>', termbox.ColorDefault, termbox.ColorRed)

	k := 0
	for i := 30; i < 30+2*len(twoLineWord); i += 2 {
		termbox.SetCell(i, 8, twoLineWord[k], termbox.ColorDefault, termbox.ColorBlack)
		k++
	}
	l := 0
	for i := 30; i < 30+2*len(threeLineWord); i += 2 {
		termbox.SetCell(i, 11, threeLineWord[l], termbox.ColorDefault, termbox.ColorBlack)
		l++
	}

	termbox.Flush()
}

func enterSongList() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 46; i += 2 {
		termbox.SetCell(i, 5, oneLineWord[j], termbox.ColorDefault, termbox.ColorBlue)
		j++
	}
	termbox.SetCell(28, moveY, '>', termbox.ColorDefault, termbox.ColorRed)
	k := 0
	for i := 30; i < 30+2*len(newSongList); i += 2 {
		termbox.SetCell(i, 8, newSongList[k], termbox.ColorDefault, termbox.ColorBlack)
		k++
	}
	l := 0
	for i := 30; i < 30+2*len(hotSongList); i += 2 {
		termbox.SetCell(i, 11, hotSongList[l], termbox.ColorDefault, termbox.ColorBlack)
		l++
	}
	n := 0
	for i := 30; i < 30+2*len(neteaseOriginalSongList); i += 2 {
		termbox.SetCell(i, 14, neteaseOriginalSongList[n], termbox.ColorDefault, termbox.ColorBlack)
		n++
	}
	p := 0
	for i := 30; i < 30+2*len(cloudMusicSoared); i += 2 {
		termbox.SetCell(i, 17, cloudMusicSoared[p], termbox.ColorDefault, termbox.ColorBlack)
		p++
	}
	termbox.Flush()
}

func enterSongHotComment() {
	//termbox.KeyInsert
}

func setListCell(results []gjson.Result) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	i := 30
	j := 8
	for index, result := range results {
		termbox.SetCell(30, j, rune(index+1), termbox.ColorDefault, termbox.ColorBlack)
		for _, v := range result.Str {
			termbox.SetCell(i, j, v, termbox.ColorDefault, termbox.ColorBlack)
			i++
		}
		j++
	}
	termbox.Flush()
}

func flush(minY, maxY int, f func()) {
	//moveX = 28
	moveY = 8
	f()
loop:
	for ev := range eventQueue {
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				moveY -= 3
			case termbox.KeyArrowDown:
				moveY += 3
			case termbox.KeyEnter:
				//....
				switch moveY {
				case 8:
					// b := true
					switch true {
					case reflect.DeepEqual(f, initDraw):
						flush(8, 17, enterSongList)
					case reflect.DeepEqual(f, enterSongList):
						//call a SetCell function.
						setListCell(crawl.TopList("云音乐新歌榜"))
					}

				case 11:
					//flush(8,,)
					switch true {
					case reflect.DeepEqual(f, initDraw):
						flush(8, 11, enterSongHotComment)
					case reflect.DeepEqual(f, enterSongHotComment):
						//call a SetCell function.
					}
				}
				break loop
			case termbox.KeyEsc:
				break loop
			}
		}
		if moveY > maxY {
			moveY = maxY
		} else if moveY < minY {
			moveY = minY
		}
		f()
	}
}

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
	flush(8, 11, initDraw)
}
