# CloudMusicBox
---
 **菜🐔一只,不免有许多bug,欢迎提 `Issue` 指正。**
---

* 使用小视频

  ![very good|412x297](https://github.com/hapi666/163CloudMusic/blob/master/ezgif.com-resize.gif)

---

  * 完成对榜单的爬取 in  `2018.9.15`   
**涵盖榜单范围**：

```go
      "云音乐新歌榜",  //每天更新一次
      "云音乐热歌榜",  //每周四更新一次
      "网易原创歌曲榜",//每周四更新一次
      "云音乐飙升榜",  //每天更新一次
      "云音乐电音榜",   
      "UK排行榜周榜",   
      "美国Billboard周",   
      "KTV嗨榜",   
      "iTunes榜",   
      "Hit FM Top榜",   
      "日本Oricon周",   
      "韩国Melon排行榜周榜",   
      "韩国Mnet排行榜周",   
      "韩国Melon原声周榜",   
      "中国TOP排行榜(港台榜)",   
      "中国TOP排行榜(内地榜)",   
      "香港电台中文歌曲龙虎榜",   
      "华语金曲榜",   
      "中国嘻哈榜",   
      "法国 NRJ EuroHot 30周榜",   
      "台湾Hito排行榜",   
      "Beatport全球电子舞曲榜",   
      "云音乐ACG音乐榜",   
      "云音乐嘻哈榜",
```
---
 * 完成对歌曲热评的爬取 in `2018.9.19`  
---
 * 简单完成命令行版本 in `2018.10.17`
 **使用方法**：  
```go

  $ mkdir Project
  
  $ cd Project
  
  $ git clone https://github.com/hapi666/163CloudMusic.git
  
  $ cd 163CloudMusic
    
  $ go run main.go
```
* 上下左右方向键控制光标移动，ESC键推出CloudMusicBox
---
* TODOList
	* 根据网易云音乐排行榜单的更新时间，定时爬取榜单歌曲列表（先拿四大榜单来做），以实现缓存效果。
	* 争取制作成一个小音乐盒，播放音乐，为歌曲点赞，评论。
