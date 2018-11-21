## site: 网易、QQ、酷狗、百度、虾米

### 网易音乐 API
- 搜索
   > http://music.163.com/weapi/search/get
    - 请求方式: POST
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
   
 - 下载   
    虾米下载直链获取方法：
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
    > http://m.kugou.com/app/i/getSongInfo.php
      - 请求方式: POST
        ```aidl
        {
          "hash": "音乐ID",
          "cmd": "playInfo",
          "from": "mkugou"
        }    
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