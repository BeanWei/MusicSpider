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