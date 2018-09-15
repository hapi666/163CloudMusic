package crawl

import (
	"code/practices/WYmusic/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

// Crawl crawl 163.music.com to get songs,etc.
func Crawl(id string) {
	params := "{\"id\":\"" + id + "\", \"n\":\"" + "1000" + "\", \"offset\":\"" + "0" + "\", \"limit\":\"" + "1000" + "\", \"total\": \"true\"}"
	p, encSecKey, err := util.EncParams(params)
	form := url.Values{}
	form.Set("params", p)
	form.Set("encSecKey", encSecKey)
	b := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", "http://music.163.com/weapi/v3/playlist/detail", b)
	if err != nil {
		log.Println("Failed to new a request")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "music.163.com")
	req.Header.Set("Referer", "http://music.163.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("Failed to do the request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Failed to read the response body")
	}
	str, err := util.GzipDecode(body)
	if err != nil {
		log.Println("Failed to decode the body.")
	}
	result := gjson.Get(string(str), "playlist.tracks.#.name")
	fmt.Println(result.String())
}
