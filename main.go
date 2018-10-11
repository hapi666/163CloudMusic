package main

import (
	"flag"
	"log"

	"github.com/hapi666/163CloudMusic/crawl"
)

var (
	topListName string
	song        string
)

func init() {
	flag.StringVar(&topListName, "l", "云音乐新歌榜", "The list that you want to get on netease cloud.")
	flag.StringVar(&song, "k", "年少有为", "The song that you want to search.")
}

func main() {
	flag.Parse()
	crawl.TopList(topListName)
	songID, err := crawl.SongID(song)
	if err != nil {
		log.Printf("Failed to get the song's id:%v", err)
	}
	crawl.HotComment(songID)
}
