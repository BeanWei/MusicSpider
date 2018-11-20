## site: 网易、QQ、酷狗、百度、虾米

### 网易 API
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