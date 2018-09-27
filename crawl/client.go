package crawl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hapi666/163CloudMusic/util"
)

func post(req *http.Request) ([]byte, error) {
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
		return nil, fmt.Errorf("Failed to read the response body:%v", err)
	}
	str, err := util.GzipDecode(body)
	if err != nil {
		return nil, fmt.Errorf("Failed to ungzip the body:%v", err)
	}
	return str, nil
}

func encodePayload(params string) (*strings.Reader, error) {
	p, encSecKey, err := util.EncParams(params)
	if err != nil {
		return nil, err
	}
	form := url.Values{}
	form.Set("params", p)
	form.Set("encSecKey", encSecKey)
	return strings.NewReader(form.Encode()), nil
}

func newParams(keyName, keyValue, limitValue string, args ...interface{}) string {
	return "{\"" + keyName + "\":\"" + keyValue + "\", \"n\":\"" + "1000" + "\"" + fmt.Sprint(args...) + ", \"offset\":\"" + "0" + "\", \"limit\":\"" + limitValue + "\", \"total\": \"true\"}"
}
