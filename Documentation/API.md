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