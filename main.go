package main

import (
	"flag"
	"log"

	"github.com/hapi666/163CloudMusic/crawl"
)

var (
	topListName string
	song        string
	music       = make(map[string]string)
)

func init() {
	flag.StringVar(&topListName, "l", "云音乐新歌榜", "What list do you want to get on netease cloud?")
	// flag.StringVar(&params["id"], "id", " ", "topList's id.")
	flag.StringVar(&song, "Search Key Word", "寒鸦少年", "the song that you want to search.")
}

func main() {
	flag.Parse()
	crawl.TopList(topListName)
	songID, err := crawl.SearchSongID(song)
	if err != nil {
		log.Printf("Failed to Search the song's id:%v", err)
	}
	crawl.Comment(songID)
}
