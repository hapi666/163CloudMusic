package main

import (
	"code/practices/WYmusic/crawl"
	"flag"
)

var bd string

func main() {
	music := make(map[string]string)
	music["云音乐新歌榜"] = "3779629"
	music["云音乐热歌榜"] = "3778678"
	music["网易原创歌曲榜"] = "2884035"
	music["云音乐飙升榜"] = "19723756"
	music["云音乐电音榜"] = "10520166"
	music["UK排行榜周榜"] = "180106"
	music["美国Billboard周榜"] = "60198"
	music["KTV嗨榜"] = "21845217"
	music["iTunes榜"] = "11641012"
	music["Hit FM Top榜"] = "120001"
	music["日本Oricon周榜"] = "60131"
	music["韩国Melon排行榜周榜"] = "3733003"
	music["韩国Mnet排行榜周榜"] = "60255"
	music["韩国Melon原声周榜"] = "46772709"
	music["中国TOP排行榜(港台榜)"] = "112504"
	music["中国TOP排行榜(内地榜)"] = "64016"
	music["香港电台中文歌曲龙虎榜"] = "10169002"
	music["华语金曲榜"] = "4395559"
	music["中国嘻哈榜"] = "1899724"
	music["法国 NRJ EuroHot 30周榜"] = "27135204"
	music["台湾Hito排行榜"] = "112463"
	music["Beatport全球电子舞曲榜"] = "3812895"
	music["云音乐ACG音乐榜"] = "71385702"
	music["云音乐嘻哈榜"] = "991319590"

	flag.StringVar(&bd, "l", "云音乐新歌榜", "What list do you want to get on netease cloud?")
	flag.Parse()
	crawl.Crawl(music[bd])
}
