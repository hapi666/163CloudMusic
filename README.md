# 163CloudMusic

 名字并没有起好，还在想，不着急。

 计划写成命令行版本。
 
 **写的仓促,不免有许多bug,欢迎提 `Issue` 指正。**

  * 完成对榜单的爬取 in  `2018.9.15`   
  
      **涵盖榜单范围**：   
      ```
          "云音乐新歌榜",  
          "云音乐热歌榜",   
          "网易原创歌曲榜",   
          "云音乐飙升榜",   
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
 * 完成对歌曲热评的爬取 in `2018.9.19`      
 **使用方法**：  
      
    ```go
      $ mkdir Project
      
      $ cd Project
      
      $ git clone https://github.com/hapi666/163CloudMusic.git
      
      $ cd 163CloudMusic
        
      $ go run main.go -l 上面任意一个榜单名称 -k 任意一个歌曲名称
    ```
 **示例**
 
    ```
      
    ```
