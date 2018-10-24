package main

import (
	"github.com/hapi666/163CloudMusic/crawl"
	"github.com/tidwall/gjson"

	termbox "github.com/nsf/termbox-go"
)

var (
	w int
	h int
	//moveX int
	moveY                   int
	eventQueue              = make(chan termbox.Event)
	results                 []gjson.Result
	boxTitle                = []rune{'网', '易', '云', '音', '乐', '欢', '迎', '您'}
	getSongLists            = []rune{'1', '.', '获', '取', '榜', '单', '.'}
	getSongHotComment       = []rune{'2', '.', '获', '取', '歌', '曲', '热', '评', '.'}
	newSongList             = []rune{'1', '.', '云', '音', '乐', '新', '歌', '榜', '.'}
	hotSongList             = []rune{'2', '.', '云', '音', '乐', '热', '歌', '榜', '.'}
	neteaseOriginalSongList = []rune{'3', '.', '网', '易', '原', '创', '歌', '曲', '榜', '.'}
	cloudMusicSoared        = []rune{'4', '.', '云', '音', '乐', '飙', '升', '榜', '.'}
	newListWords            = []rune{'云', '音', '乐', '新', '歌', '榜'}
)

func initDraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 30+2*len(boxTitle); i += 2 {
		termbox.SetCell(i, 5, boxTitle[j], termbox.ColorDefault, termbox.ColorBlue)
		j++
	}

	termbox.SetCell(28, moveY, '>', termbox.ColorDefault, termbox.ColorRed)

	k := 0
	for i := 30; i < 30+2*len(getSongLists); i += 2 {
		termbox.SetCell(i, 8, getSongLists[k], termbox.ColorDefault, termbox.ColorBlack)
		k++
	}
	l := 0
	for i := 30; i < 30+2*len(getSongHotComment); i += 2 {
		termbox.SetCell(i, 11, getSongHotComment[l], termbox.ColorDefault, termbox.ColorBlack)
		l++
	}

	termbox.Flush()
}

func enterSongList() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 0
	for i := 30; i < 46; i += 2 {
		termbox.SetCell(i, 5, boxTitle[j], termbox.ColorDefault, termbox.ColorBlue)
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

func setListCell() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	j := 1
	// k := 0
	// for i := 30; i < 42; i += 2 {
	// 	termbox.SetCell(i, 0, newListWords[k], termbox.ColorDefault, termbox.ColorBlue)
	// 	k++
	// }
	for _, result := range results {
		i := 32
		termbox.SetCell(30, j, '*', termbox.ColorDefault, termbox.ColorBlack)
		s := []rune(result.Str)
		//termbox-go 只能显示终端屏尺寸的大小，超过此时显示的坐标时就不会再设置cell坐标
		for _, v := range s {
			termbox.SetCell(i, j, v, termbox.ColorDefault, termbox.ColorBlack)
			i += 2
		}
		j++
	}
	termbox.Flush()
}

func flush(minY, maxY int, f func(), pageFlag int) {
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
				switch moveY {
				case 8:
					if pageFlag == 1 {
						flush(8, 17, enterSongList, 2)
					}
					if pageFlag == 2 { //在第二个页面点击第一个的enter键
						results = crawl.TopList("3779629")
						flush(8, 8+len(results), setListCell, 4)
					}
					if pageFlag == 3 { //在第三个页面点击第一个的enter键
						//Call HotComment function.
					}
				case 11:
					if pageFlag == 1 { //在第一个页面点击第二个的enter键
						flush(8, 11, enterSongHotComment, 3)
					}
					if pageFlag == 2 { //在第二个页面点击第二个的enter键
						results = crawl.TopList("3778678")
						flush(8, 8+len(results), setListCell, 4)
					}
				case 14:
					if pageFlag == 2 {
						results = crawl.TopList("2884035")
						flush(8, 8+len(results), setListCell, 4)
					}
				case 17:
					if pageFlag == 2 {
						results = crawl.TopList("19723756")
						flush(8, 8+len(results), setListCell, 4)
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
