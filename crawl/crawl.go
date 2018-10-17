package crawl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// TopList crawl 163.music.com to get top-list.
func TopList(topListID string) {
	params := newParams("id", topListID, "1000")
	log.Println(params)
	b, err := encodePayload(params)
	if err != nil {
		log.Printf("Failed to ungzip the payload:%v", err)
	}
	req, _ := http.NewRequest("POST", "http://music.163.com/weapi/v3/playlist/detail", b)
	str, err := post(req)
	if err != nil {
		log.Println(err)
	}
	result := gjson.Get(string(str), "playlist.tracks.#.name")
	log.Println(result.String())
}

// SongID get the corresponding song id by songName.
func SongID(songName string) (string, error) {
	params := newParams("s", songName, "20", ",\"type\":\"1\"")
	log.Println(params)
	b, err := encodePayload(params)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", "http://music.163.com/weapi/search/get", b)
	str, err := post(req)
	if err != nil {
		return "", err
	}
	id := gjson.Get(string(str), "result.songs.#.id")
	if !id.Exists() {
		return "", fmt.Errorf("Don't have the song:%v", songName)
	}
	log.Println(id.Array()[0].String())
	return id.Array()[0].String(), nil
}

// HotComment get a hot comment of the song by the song id.
func HotComment(rid string) {
	log.Println(rid)
	params := newParams("rid", rid, "30")
	fmt.Println(params)
	b, err := encodePayload(params)
	if err != nil {
		log.Printf("Failed to encode the payload:%v", err)
	}
	req, _ := http.NewRequest("POST", "http://music.163.com/weapi/v1/resource/comments/R_SO_4_"+rid, b)
	str, err := post(req)
	if err != nil {
		log.Printf("Faild to send the POST method:%v", err)
	}
	hotComment := gjson.Get(string(str), "hotComments.#.content")
	log.Println(len(hotComment.Array()))
	log.Println(hotComment.String())
}
