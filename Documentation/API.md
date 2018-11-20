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
      