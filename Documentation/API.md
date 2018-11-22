## site: 网易、QQ、酷狗、百度、虾米

### 网易音乐 API
- 搜索
   > http://music.163.com/weapi/search/get
    - 请求方式: POST
    - 请求参数需要加密
   ```aidl
      {
          "s": "搜索关键字",
          "offset": "0",
          "limit": "10",
          "type": "1",
          "total": "true"
          "n": "1000"
      }
   ```
- 详情(根据音乐ID获取对应的信息)   
   > http://music.163.com/api/v3/song/detail?id=35847388&c=[{"id":"35847388"}]
     - 请求方式: GET
        ```aidl
        id: 音乐ID
        c: [{"id":"音乐ID"}]    
        ```
- 歌词
   > http://music.163.com/api/song/lyric?id=418603133&os=linux&lv=-1&kv=-1&tv=-1
     - 请求方式: GET
     ```aidl
    "id": "歌曲ID", 
    "os": "linux", 
    "lv": "-1", 
    "kv": "-1", 
    "tv": "-1"
    ```      
- 专辑
   > http://music.163.com/api/v1/album/专辑ID
     - 请求方式: POST
     ```aidl
      {
          "id": "专辑ID", 
          "total": "true", 
          "offset": "0", 
          "limit": "1000", 
          "ext": "true", 
          "private_cloud": "true"
      }
     ```  
- 艺人
   > http://music.163.com/api/v1/artist/艺人ID
     - 请求方式: POST
       ```aidl
        {
           "albumId": "艺人ID", 
           "ext": "true", 
           "private_cloud": "true", 
           "top": 50
        }
       ```
- 歌单
   > http://music.163.com/weapi/v3/playlist/detail
     - 请求方式: POST 
     - 请求参数需要加密 
       ```aidl
       {
           "s": "0",
           "id": 歌单ID,
           "n": '1000',
           "t": '0',
       }    
       ```               

 ### QQ音乐 API
 - 搜索
    > https://c.y.qq.com/soso/fcgi-bin/client_search_cp?w=hello&format=json&p=1&n=10&aggr=1&lossless=1&cr=1&new_json=1
     - 请求方式: GET
        ```aidl
        W: 搜索关键字
        format: 返回值形式    
        p: 页数
        n: 
        aggr:
        lossless:
        cr:
        new_json:
        ```
 - 详情(根据音乐ID获取对应的信息)   
    > https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg?songmid=000bdxcy3ta8Q3&platform=yqq&format=json
      - 请求方式: GET
         ```aidl
         songmid: 音乐ID
         platform: 平台
         format: 请求返回类型       
         ```
 - 歌词
    > https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg?songmid=000bdxcy3ta8Q3&g_tk=%!s(float64=5381)
      - 请求方式: GET
          ```aidl
         "songmid": "歌曲ID", 
         "g_tk": 5381
         ```         
 - 专辑
   > https://c.y.qq.com/v8/fcg-bin/fcg_v8_album_detail_cp.fcg?newsong=1&albummid=专辑ID号&platform=mac&format=json
     - 请求方式: GET
        ```aidl
        "albummid": "专辑ID", 
        "platform": "mac", 
        "format": "json", 
        "newsong": "1"
        ```   
 - 艺人
    > https://c.y.qq.com/v8/fcg-bin/fcg_v8_singer_track_cp.fcg?singermid=003Nz2So3XXYek&begin=0&num=50&order=listen&platform=mac&newsong=1
      - 请求方式: GET
        ```aidl
        "singermid": "艺人ID", 
        "begin": "0", 
        "num": "50", 
        "order": "listen", 
        "platform": "mac", 
        "newsong": "1"
        ``` 
 - 歌单
    > https://c.y.qq.com/v8/fcg-bin/fcg_v8_playlist_cp.fcg?id=1144188779&format=json&newsong=1&platform=jqspaframe.json
      - 请求方式: GET
        ```aidl
        "id": "歌单ID", 
        "format": "json", 
        "newsong": "1", 
        "platform": "jqspaframe.json"
        ```              
 
 ### 虾米音乐 API
 - 搜索
    > http://api.xiami.com/web?v=2.0&key=hello&limit=30&page=1&r=search/songs&app_key=1
      - 请求方式: GET
        ```aidl
        key: 搜索关键字
        v: 版本号2.0  
        p: 页数
        limit: 上限数
        r: search/songs规则
        app_key: 1
        
        - 请求头必须加上 Referer: http://h.xiami.com/ 否则会被认为非法请求
        ```
 - 详情(根据音乐ID获取对应的信息)   
    > http://www.xiami.com/song/playlist/id/1807124714/object_name/default/object_id/0/cat/json
      - 请求方式: GET
         ```aidl
         1807124714: 音乐ID(用此ID手动构造请求URL)     
         ```   
 - 歌词
    暂时只知道可以在歌曲详情中获取对应歌词，其他方法尚未成功。\
    从 `详情` 中获取歌词的链接, 这里有点特别奇怪的是，我代码里面请求歌词链接获取到的歌词
    是不需要做额外的处理，但是在浏览器上看到的歌词返回值是有<数字>这种符号的。
    
 - 专辑
    > https://www.xiami.com/album/songs/id/专辑ID
      - 请求方式: GET \
      :lock: 此接口待更换，不建议使用, 频繁请求便会被ban掉禁止访问。 
   
 - 下载   
    虾米下载直链获取方法：\
    先请求下面的URL，从返回的html页面中拿到经过凯撒加密后的字符串，由字符串的首位分析得出此
    凯撒矩阵的行数，然后拼接出编码的URL，最后解码URL获取真实的下载直链。
    > http://www.xiami.com/widget/xml-single/uid/0/sid/歌曲ID
      - 请求方式: GET

    
 ### 酷狗音乐 API
 - 搜索
     > http://mobilecdn.kugou.com/api/v3/search/song?keyword=hello&pagesize=10&page=1&area_code=1&plat=2&version=8990&api_ver=1&correct=1&tag=1&sver=5&showtype=10
      - 请求方式: GET
         ```aidl
         keyword: 搜索关键字
         pagesize: 搜索页数 
         page: 当前页
         area_code: 区域代码
         plat:平台代号
         version: 版本号
         api_ver: api接口版本号
         correct: 
         tag:
         sver:
         showtype:     
         ```
 - 详情(根据音乐ID获取对应的信息)   
    > http://m.kugou.com/app/i/getSongInfo.php?cmd=playInfo&hash=d353b69a3b7f1a250000c5daabb8a4f1
      - 请求方式: GET
        ```aidl
        "cmd": "playInfo", 
        "hash": "歌曲 hash ID"  
        ```
 - 歌词 \
    这里获取歌词的接口其实也是获取歌曲详情的接口，相比上面的接口这个附带了歌词信息
    > http://www.kugou.com/yy/index.php?r=play/getdata&hash=CB7EE97F4CC11C4EA7A1FA4B516A5D97
    - 请求方式: GET
      ```aidl
      "r": "play/getdata", 
      "hash": "歌曲 Hash ID"
      ```  
 - 专辑
    > http://mobilecdn.kugou.com/api/v3/album/song?albumid=1645030&area_code=1&plat=2&page=1&pagesize=-1&version=8990
      - 请求方式: GET
        ```aidl
        "albumid": "专辑ID", 
        "area_code": "1", 
        "plat": "2", 
        "page": "1", 
        "pagesize": "-1", 
        "version": "8990"
        ```   
 - 艺人
    > http://mobilecdn.kugou.com/api/v3/singer/song?version=8990&singerid=3520&area_code=1&page=1&plat=0&pagesize=50
      - 请求方式: GET       
        ```aidl
         "singerid": "艺人ID", 
         "area_code": "1", 
         "page": "1", 
         "plat": "0", 
         "pagesize": "50", 
         "version": "8990"
        ``` 
 - 歌单
   > http://mobilecdn.kugou.com/api/v3/special/song?specialid=119859&area_code=1&plat=2&page=1&pagesize=-1&version=8990
     - 请求方式: GET
       ```aidl
       "specialid": "歌单ID", 
       "area_code": "1", 
       "plat": "2", 
       "page": "1", 
       "pagesize": "-1", 
       "version": "8990" 
       ```         
           
 - 下载\
    step1:请求如下URL获取hash值
    > http://media.store.kugou.com/v1/get_res_privilege
     - 请求方式: POST
       ```aidl
        {
           "relate":"1",
           "userid":"0",
           "vip":"0",
           "appid":"1000",
           "token":"",
           "behavior":"download",
           "area_code":"1",
           "clientver":"8990",
           "resource":[{"id":0,"type":"audio","hash":"歌曲的id(酷狗的歌曲id是hash加密的)"}]}
       ```  
    step2: 将获取到的hash值加盐(kugouv2)后进行MD5加密, 然后请求如下URL获取下载直链
    > http://trackercdn.kugou.com/i/v2/
     - 请求方式: GET
     ```aidl
    "hash": "D2BB1EE2952B8EC60EAAFBE129190101", 
    "key": "1e08e8c8b68f98598654863732b25a7c", 
    "pid": "3", 
    "behavior": "play",   这里传play和download结果是一样的
    "cmd": "25", 
    "version": "8990"
    ```             
     
 ### 百度音乐 API
 - 搜索
     > http://musicapi.taihe.com/v1/restserver/ting?page_no=1&page_size=10&from=qianqianmini&method=baidu.ting.search.merge&isNew=1&platform=darwin&version=11.2.1&query=hello
      - 请求方式: GET
         ```aidl
         query: 搜索关键字
         page_size: 搜索页数 
         page_no: 当前页
         from: api源 qianqianmini
         method: 方法参数
         isNew:
         platform:平台参数  
         version: 版本号
         ```     
 - 详情(根据音乐ID获取对应的信息)   
    >  http://musicapi.taihe.com/v1/restserver/ting?platform=darwin&version=1.0.0&e=vchaKgAigShr/UkgYA0bw0joanSuEPhWqIaaazfh0aUT/N4CRAIraAHDhnXg6DeQ&songid=578055564&from=qianqianmini&method=baidu.ting.song.getInfos&res=1
      - 请求方式: GET
         ```aidl
         id: 音乐ID 
         from:  api源？
         method: 方法类别
         res: 
         platform: 平台
         version: 版本号
         e: 将音乐ID和时间戳合并与密钥进行加密处理得出来的加密字符串   
         ```   
 - 歌词
    > http://musicapi.taihe.com/v1/restserver/ting?songid=578055564&from=qianqianmini&method=baidu.ting.song.lry&platform=darwin&version=1.0.0
      - 请求方式: GET
         ```aidl
         "songid": "歌曲ID", 
         "from": "qianqianmini", 
         "method": "baidu.ting.song.lry", 
         "platform": "darwin", 
         "version": "1.0.0"   
         ``` 
 - 专辑
    > http://musicapi.taihe.com/v1/restserver/ting?album_id=541223882&from=qianqianmini&method=baidu.ting.album.getAlbumInfo&platform=darwin&version=11.2.1
      - 请求方式: GET
        ```aidl
        "album_id": "专辑ID", 
        "from": "qianqianmini", 
        "method": "baidu.ting.album.getAlbumInfo", 
        "platform": "darwin", 
        "version": "11.2.1"
        ```  
 - 艺人
    > http://musicapi.taihe.com/v1/restserver/ting?version=11.2.1&artistid=362&from=qianqianmini&method=baidu.ting.artist.getSongList&limits=50&platform=darwin&offset=0&tinguid=0
      - 请求方式: GET
        ```aidl
        "artistid": "艺人ID", 
        "from": "qianqianmini", 
        "method": "baidu.ting.artist.getSongList", 
        "limits": "50", 
        "platform": "darwin", 
        "offset": "0", 
        "tinguid": "0", 
        "version": "11.2.1"
        ```  
 - 歌单
    > http://musicapi.taihe.com/v1/restserver/ting?method=baidu.ting.diy.gedanInfo&platform=darwin&version=11.2.1&listid=364201689&from=qianqianmini
      - 请求方式: GET
        ```aidl
         "listid": "歌单ID", 
         "from": "qianqianmini", 
         "method": "baidu.ting.diy.gedanInfo", 
         "platform": "darwin", 
         "version": "11.2.1"   
        ```              
                 
 - 下载
    > http://musicapi.taihe.com/v1/restserver/ting?version=1.0.0&e=vchaKgAigShr/UkgYA0bw1nX9xQpRiZHIzBqV/lBoBUDS2QfPcumcfd92CJXkVA2&songid=578055564&from=qianqianmini&method=baidu.ting.song.getInfos&res=1&platform=darwin
     - 请求方式: GET
     - :lock:注意: 非大陆IP似乎无法请求获得下载直链
     ```aidl
       "songid": "歌曲ID",
       "from": "qianqianmini", 
       "method": "baidu.ting.song.getInfos", 
       "res": "1", 
       "platform": "darwin", 
       "version": "1.0.0", 
       "e": "将音乐ID和时间戳合并与密钥进行加密处理得出来的加密字符串"
    ```             