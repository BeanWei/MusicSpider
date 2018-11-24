<h1 align="center">🔍 MusicSpider 🔎</p>

### 📑 概述
👾 Music Spider 是使用Golang写的音乐聚合爬虫，目前支持的站点包括 网易、QQ、虾米、酷狗、百度。

写这个的目的是为了基于此做其他的进一步开发，为了方便，同时我也写了一个 format 辅助方法来减轻格式化json的工作量。

基于format，我做了一个简单的API接口封装。

### 🌟 来源
感谢 [@metowolf](https://github.com/metowolf) 的这个 [Meting a powerful music API framework](https://github.com/metowolf/Meting) 项目。

这个项目使用的是PHP，然后我使用Golang写了一遍。

项目的接口和加解密的方法有些是来自于它处, 在此一并感谢。

### 💢 说明
此项目目前任然处于开发阶段，代码细节还尚未处理，虽然代码很简单但是因为接口太多，so, 慢慢优化吧。

目前虾米音乐的接口分析还未完善，其他站点的接口可能会不定时改变，如果你对此项目感兴趣，欢迎提交PR，有问题可以提交Issue我们一起解决。

后面会基于此项目做一些其他的开发，比如说做个音乐聚合网站？💖 ，所以说这个项目也会不定时更新，嘻嘻😄

### 🔗 文档说明
所有的原始接口和请求细节、参数，你可以在 \Documentation\API.md 中查看。

所有的原始Json返回值，你可以在 \JsonResultFiles\... 中查看。

### 💡 接口封装
基于 [gin](https://github.com/gin-gonic/gin) golang web 框架写了一个简单的接口封装。

🔎 食用方法

<details>
<summary>http://localhost:8080/api/v1/search/netease/hello</summary>
```json
{
    "result": {
        "count": 30,
        "songs": [
            {
                "album_id": "3377030",
                "album_name": "Hello",
                "singer_id": "46487",
                "singer_name": "Adele",
                "song_id": "35847388",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=35847388"
            },
            {
                "album_id": "3190201",
                "album_name": "Hello",
                "singer_id": "381949",
                "singer_name": "OMFG",
                "song_id": "33211676",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=33211676"
            },
            {
                "album_id": "34930111",
                "album_name": "Hello",
                "singer_id": "12145027",
                "singer_name": "Ayindé",
                "song_id": "436355540",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=436355540"
            },
            {
                "album_id": "3184406",
                "album_name": "Hello",
                "singer_id": "1057278",
                "singer_name": "Barbara Opsomer",
                "song_id": "33166086",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=33166086"
            },
            {
                "album_id": "34914813",
                "album_name": "Hello（DCGO Flip）",
                "singer_id": "12037211",
                "singer_name": "DCGO",
                "song_id": "435948709",
                "song_name": "Hello（DCGO Flip）",
                "source_url": "https://music.163.com/#/song?id=435948709"
            },
            {
                "album_id": "3377045",
                "album_name": "25",
                "singer_id": "46487",
                "singer_name": "Adele",
                "song_id": "36841430",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=36841430"
            },
            {
                "album_id": "3390017",
                "album_name": "Hello (Taps \u0026 JDam Cover)",
                "singer_id": "1077068",
                "singer_name": "LYAR",
                "song_id": "36103120",
                "song_name": "Hello (Taps \u0026 JDam Cover) ",
                "source_url": "https://music.163.com/#/song?id=36103120"
            },
            {
                "album_id": "167355",
                "album_name": "Can't Slow Down (Deluxe Edition)",
                "singer_id": "38125",
                "singer_name": "Lionel Richie",
                "song_id": "1648574",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=1648574"
            },
            {
                "album_id": "34865381",
                "album_name": "I Love You",
                "singer_id": "12002253",
                "singer_name": "杭立旺",
                "song_id": "428391758",
                "song_name": "Hello（3D）",
                "source_url": "https://music.163.com/#/song?id=428391758"
            },
            {
                "album_id": "1872173",
                "album_name": "The N.W.A Legacy, Vol. 2",
                "singer_id": "69344",
                "singer_name": "N.W.A",
                "song_id": "20189570",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=20189570"
            },
            {
                "album_id": "2623237",
                "album_name": "登“封”造极",
                "singer_id": "5349",
                "singer_name": "魏晨",
                "song_id": "27591863",
                "song_name": "哈喽",
                "source_url": "https://music.163.com/#/song?id=27591863"
            },
            {
                "album_id": "529100",
                "album_name": "Replay -君は仆のeverything-",
                "singer_id": "127582",
                "singer_name": "SHINee",
                "song_id": "25727706",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=25727706"
            },
            {
                "album_id": "3438260",
                "album_name": "Hello",
                "singer_id": "858180",
                "singer_name": "Tyler Carter",
                "song_id": "40147463",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=40147463"
            },
            {
                "album_id": "3377421",
                "album_name": "最新热歌慢摇113",
                "singer_id": "29020",
                "singer_name": "Conor Maynard",
                "song_id": "38019041",
                "song_name": "Hello (Cover)",
                "source_url": "https://music.163.com/#/song?id=38019041"
            },
            {
                "album_id": "73467931",
                "album_name": "EVERYD4Y",
                "singer_id": "976163",
                "singer_name": "WINNER",
                "song_id": "1313104171",
                "song_name": "HELLO",
                "source_url": "https://music.163.com/#/song?id=1313104171"
            },
            {
                "album_id": "177188",
                "album_name": "Hello",
                "singer_id": "39337",
                "singer_name": "Martin Solveig",
                "song_id": "1750531",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=1750531"
            },
            {
                "album_id": "34767287",
                "album_name": "Hello (Will \u0026 Tim Remix)",
                "singer_id": "1192002",
                "singer_name": "Will \u0026 Tim",
                "song_id": "419594876",
                "song_name": "Hello (Will \u0026 Tim Remix)",
                "source_url": "https://music.163.com/#/song?id=419594876"
            },
            {
                "album_id": "3066005",
                "album_name": "太陽と月の塔",
                "singer_id": "1028003",
                "singer_name": "ウルトラタワー",
                "song_id": "29719037",
                "song_name": "ハロー",
                "source_url": "https://music.163.com/#/song?id=29719037"
            },
            {
                "album_id": "39082597",
                "album_name": "U",
                "singer_id": "13520323",
                "singer_name": "YahikoNy",
                "song_id": "567183978",
                "song_name": "Hello (Angelika Vee Cover R.M Remix)",
                "source_url": "https://music.163.com/#/song?id=567183978"
            },
            {
                "album_id": "73385818",
                "album_name": "Hello",
                "singer_id": "12113223",
                "singer_name": "RAPPER32",
                "song_id": "1310883113",
                "song_name": "HELLO",
                "source_url": "https://music.163.com/#/song?id=1310883113"
            },
            {
                "album_id": "36705013",
                "album_name": "TEEN, AGE",
                "singer_id": "1080132",
                "singer_name": "SEVENTEEN",
                "song_id": "516358597",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=516358597"
            },
            {
                "album_id": "3093169",
                "album_name": "Новая волна 2014",
                "singer_id": "97506",
                "singer_name": "Mandinga",
                "song_id": "30070232",
                "song_name": "Hello (feat. Fly Project)",
                "source_url": "https://music.163.com/#/song?id=30070232"
            },
            {
                "album_id": "35450132",
                "album_name": "Hello 1",
                "singer_id": "166014",
                "singer_name": "金志文",
                "song_id": "475277342",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=475277342"
            },
            {
                "album_id": "2901266",
                "album_name": "Hello",
                "singer_id": "975255",
                "singer_name": "Mischa",
                "song_id": "28838746",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=28838746"
            },
            {
                "album_id": "2799002",
                "album_name": "AGA",
                "singer_id": "768208",
                "singer_name": "AGA",
                "song_id": "28453002",
                "song_name": "哈啰",
                "source_url": "https://music.163.com/#/song?id=28453002"
            },
            {
                "album_id": "34720325",
                "album_name": "STREET",
                "singer_id": "1081893",
                "singer_name": "Hani",
                "song_id": "414979798",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=414979798"
            },
            {
                "album_id": "35373770",
                "album_name": "HELLO",
                "singer_id": "12373558",
                "singer_name": "海绵老伯",
                "song_id": "472691766",
                "song_name": "HELLO",
                "source_url": "https://music.163.com/#/song?id=472691766"
            },
            {
                "album_id": "34753158",
                "album_name": "Hello",
                "singer_id": "11987087",
                "singer_name": "October Child",
                "song_id": "419250511",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=419250511"
            },
            {
                "album_id": "34701609",
                "album_name": "HELLO",
                "singer_id": "12047464",
                "singer_name": "DR. FRESCH",
                "song_id": "413812575",
                "song_name": "Hello",
                "source_url": "https://music.163.com/#/song?id=413812575"
            },
            {
                "album_id": "2944016",
                "album_name": "괜찮아 사랑이야 Pop OST",
                "singer_id": "41910",
                "singer_name": "Quentin Mosimann",
                "song_id": "28912117",
                "song_name": "Hello (Ira Ver.) (Radio Edit)",
                "source_url": "https://music.163.com/#/song?id=28912117"
            }
        ],
        "source": "netease"
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http://localhost:8080/api/v1/artist/netease/3681</summary>
```json
{
    "result": {
        "brief_desc": "李志，南京知名音乐人。1978年生于江苏金坛。1997年进入东南大学自动控制系，之后一直在南京工作生活至今。以独立音乐人的身份出版发行了六张录音室专辑《被禁忌的游戏》（2004）、《梵高先生》（2005）、《这个世界会好吗》（2006）、《我爱南京》（2009）、《你好郑州》（2010）、《F》（2011），和三张演唱会录音专辑《工体东路没有人》（2009）、《勾三搭四》（2014）、《i/O》（2015）。",
        "cover_url": "http://p3.music.126.net/71bgc394k7FzQ_MPFOoLMg==/6049512976723323.jpg",
        "singer_name": "李志",
        "songs": [
            {
                "song_id": "26508186",
                "song_name": "天空之城",
                "source_url": "https://music.163.com/#/song?id=26508186"
            },
            {
                "song_id": "25867002",
                "song_name": "关于郑州的记忆",
                "source_url": "https://music.163.com/#/song?id=25867002"
            },
            {
                "song_id": "26508240",
                "song_name": "梵高先生",
                "source_url": "https://music.163.com/#/song?id=26508240"
            },
            {
                "song_id": "26523120",
                "song_name": "和你在一起",
                "source_url": "https://music.163.com/#/song?id=26523120"
            },
            {
                "song_id": "26508242",
                "song_name": "你离开了南京，从此没有人和我说话",
                "source_url": "https://music.163.com/#/song?id=26508242"
            },
            {
                "song_id": "26353044",
                "song_name": "忽然",
                "source_url": "https://music.163.com/#/song?id=26353044"
            },
            {
                "song_id": "29724295",
                "song_name": "热河",
                "source_url": "https://music.163.com/#/song?id=29724295"
            },
            {
                "song_id": "26508232",
                "song_name": "山阴路的夏天",
                "source_url": "https://music.163.com/#/song?id=26508232"
            },
            {
                "song_id": "26522011",
                "song_name": "米店",
                "source_url": "https://music.163.com/#/song?id=26522011"
            },
            {
                "song_id": "424474915",
                "song_name": "这个世界会好吗 (2016 unplugged)",
                "source_url": "https://music.163.com/#/song?id=424474915"
            },
            {
                "song_id": "29724292",
                "song_name": "定西",
                "source_url": "https://music.163.com/#/song?id=29724292"
            },
            {
                "song_id": "28406900",
                "song_name": "和你在一起2013版[Live]",
                "source_url": "https://music.163.com/#/song?id=28406900"
            },
            {
                "song_id": "30212877",
                "song_name": "墙上的向日葵(2014i/O版)",
                "source_url": "https://music.163.com/#/song?id=30212877"
            },
            {
                "song_id": "30967318",
                "song_name": "鸵鸟\u0026天空之城\u0026我们不能失去信仰(2014i/O版)",
                "source_url": "https://music.163.com/#/song?id=30967318"
            },
            {
                "song_id": "34200930",
                "song_name": "你离开了南京，从此没有人和我说话 2015现场版",
                "source_url": "https://music.163.com/#/song?id=34200930"
            },
            {
                "song_id": "34200934",
                "song_name": "热河 2015现场版",
                "source_url": "https://music.163.com/#/song?id=34200934"
            },
            {
                "song_id": "26508228",
                "song_name": "寻找",
                "source_url": "https://music.163.com/#/song?id=26508228"
            },
            {
                "song_id": "435948318",
                "song_name": "送别",
                "source_url": "https://music.163.com/#/song?id=435948318"
            },
            {
                "song_id": "26508184",
                "song_name": "结婚",
                "source_url": "https://music.163.com/#/song?id=26508184"
            },
            {
                "song_id": "478490652",
                "song_name": "回答",
                "source_url": "https://music.163.com/#/song?id=478490652"
            },
            {
                "song_id": "26522009",
                "song_name": "听妈妈讲那过去的事情",
                "source_url": "https://music.163.com/#/song?id=26522009"
            },
            {
                "song_id": "26508205",
                "song_name": "被禁忌的游戏",
                "source_url": "https://music.163.com/#/song?id=26508205"
            },
            {
                "song_id": "406072281",
                "song_name": "普希金 (2015动静版)",
                "source_url": "https://music.163.com/#/song?id=406072281"
            },
            {
                "song_id": "26522014",
                "song_name": "再见",
                "source_url": "https://music.163.com/#/song?id=26522014"
            },
            {
                "song_id": "31381877",
                "song_name": "这个世界会好吗2015",
                "source_url": "https://music.163.com/#/song?id=31381877"
            },
            {
                "song_id": "26508185",
                "song_name": "鸵鸟",
                "source_url": "https://music.163.com/#/song?id=26508185"
            },
            {
                "song_id": "26508235",
                "song_name": "杭州",
                "source_url": "https://music.163.com/#/song?id=26508235"
            },
            {
                "song_id": "26508244",
                "song_name": "想起了她",
                "source_url": "https://music.163.com/#/song?id=26508244"
            },
            {
                "song_id": "551338613",
                "song_name": "热河 (相信未来版)",
                "source_url": "https://music.163.com/#/song?id=551338613"
            },
            {
                "song_id": "424474911",
                "song_name": "黑色信封 (2016 unplugged)",
                "source_url": "https://music.163.com/#/song?id=424474911"
            },
            {
                "song_id": "26508221",
                "song_name": "九月",
                "source_url": "https://music.163.com/#/song?id=26508221"
            },
            {
                "song_id": "1297743290",
                "song_name": "金城兰州",
                "source_url": "https://music.163.com/#/song?id=1297743290"
            },
            {
                "song_id": "29724294",
                "song_name": "不多",
                "source_url": "https://music.163.com/#/song?id=29724294"
            },
            {
                "song_id": "435948312",
                "song_name": "兰花草",
                "source_url": "https://music.163.com/#/song?id=435948312"
            },
            {
                "song_id": "424474922",
                "song_name": "春末的南方城市 (2016 unplugged)",
                "source_url": "https://music.163.com/#/song?id=424474922"
            },
            {
                "song_id": "26508182",
                "song_name": "意味",
                "source_url": "https://music.163.com/#/song?id=26508182"
            },
            {
                "song_id": "26353057",
                "song_name": "秋天的老狼",
                "source_url": "https://music.163.com/#/song?id=26353057"
            },
            {
                "song_id": "29724293",
                "song_name": "看见",
                "source_url": "https://music.163.com/#/song?id=29724293"
            },
            {
                "song_id": "435948315",
                "song_name": "Hey Jude",
                "source_url": "https://music.163.com/#/song?id=435948315"
            },
            {
                "song_id": "478490647",
                "song_name": "尽头",
                "source_url": "https://music.163.com/#/song?id=478490647"
            },
            {
                "song_id": "26508230",
                "song_name": "门",
                "source_url": "https://music.163.com/#/song?id=26508230"
            },
            {
                "song_id": "478490649",
                "song_name": "铅笔",
                "source_url": "https://music.163.com/#/song?id=478490649"
            },
            {
                "song_id": "26523123",
                "song_name": "妈妈",
                "source_url": "https://music.163.com/#/song?id=26523123"
            },
            {
                "song_id": "440464353",
                "song_name": "克兰河",
                "source_url": "https://music.163.com/#/song?id=440464353"
            },
            {
                "song_id": "551335618",
                "song_name": "关于郑州的记忆 (相信未来版)",
                "source_url": "https://music.163.com/#/song?id=551335618"
            },
            {
                "song_id": "478491639",
                "song_name": "序曲",
                "source_url": "https://music.163.com/#/song?id=478491639"
            },
            {
                "song_id": "26508239",
                "song_name": "董卓瑶",
                "source_url": "https://music.163.com/#/song?id=26508239"
            },
            {
                "song_id": "440411394",
                "song_name": "一头偶像",
                "source_url": "https://music.163.com/#/song?id=440411394"
            },
            {
                "song_id": "26508236",
                "song_name": "日",
                "source_url": "https://music.163.com/#/song?id=26508236"
            },
            {
                "song_id": "435948316",
                "song_name": "采蘑菇的小姑娘",
                "source_url": "https://music.163.com/#/song?id=435948316"
            }
        ],
        "source": "netease",
        "total_songs": 50
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http://localhost:8080/api/v1/album/kugou/1645030</summary>
```json
{
    "result": {
        "album_name": "",
        "brief_desc": "",
        "cover_url": "",
        "publish_time": "",
        "singer_id": "",
        "singer_name": "周杰伦",
        "songs": [
            {
                "song_id": "09E8DE70A24C97B92A29F6A19F3528A2",
                "song_name": "周杰伦 - 床边故事",
                "source_url": "http://www.kugou.com/song/#hash=09E8DE70A24C97B92A29F6A19F3528A2"
            },
            {
                "song_id": "7122C664CAFF3F7B5BDC76F05C42802A",
                "song_name": "周杰伦 - 说走就走",
                "source_url": "http://www.kugou.com/song/#hash=7122C664CAFF3F7B5BDC76F05C42802A"
            },
            {
                "song_id": "49FCC392E6C9EC7090CA9C9ABB640B05",
                "song_name": "周杰伦 - 一点点",
                "source_url": "http://www.kugou.com/song/#hash=49FCC392E6C9EC7090CA9C9ABB640B05"
            },
            {
                "song_id": "DD90373961E86D70DDEBA5FBA94225BC",
                "song_name": "周杰伦 - 前世情人",
                "source_url": "http://www.kugou.com/song/#hash=DD90373961E86D70DDEBA5FBA94225BC"
            },
            {
                "song_id": "E6D9FF1079D392E2A005ED61A7370835",
                "song_name": "周杰伦 - 英雄",
                "source_url": "http://www.kugou.com/song/#hash=E6D9FF1079D392E2A005ED61A7370835"
            },
            {
                "song_id": "F1D0D0EC62A70C666A42220B4AEB49DE",
                "song_name": "周杰伦、aMEI - 不该",
                "source_url": "http://www.kugou.com/song/#hash=F1D0D0EC62A70C666A42220B4AEB49DE"
            },
            {
                "song_id": "40583041038926F890B71274E213CB65",
                "song_name": "周杰伦 - 土耳其冰淇淋",
                "source_url": "http://www.kugou.com/song/#hash=40583041038926F890B71274E213CB65"
            },
            {
                "song_id": "5FCE4CBCB96D6025033BCE2025FC3943",
                "song_name": "周杰伦 - 告白气球",
                "source_url": "http://www.kugou.com/song/#hash=5FCE4CBCB96D6025033BCE2025FC3943"
            },
            {
                "song_id": "DFE35A8F5618240EBBDF5D9753AE4C7D",
                "song_name": "周杰伦 - Now You See Me",
                "source_url": "http://www.kugou.com/song/#hash=DFE35A8F5618240EBBDF5D9753AE4C7D"
            },
            {
                "song_id": "7620938F040E0B23CAFE788B389FC846",
                "song_name": "周杰伦 - 爱情废柴",
                "source_url": "http://www.kugou.com/song/#hash=7620938F040E0B23CAFE788B389FC846"
            }
        ],
        "source": "kugou",
        "total_songs": 10
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http://localhost:8080/api/v1/playlist/baidu/364201689</summary>
```json
{
    "result": {
        "brief_desc": "由多位歌手合作完成的对唱情歌具有一种超脱的魔力，将情歌里的暗涌百转千回，视听上更有此起彼伏的层叠感。本单精选环球唱片下宝丽金时代的那些经典的珍藏版对唱，用音乐将曾经的美好重组，让我们再次追忆当年的感动与美好。",
        "cover_url": "http://musicugc.cdn.qianqian.com/ugcdiy/pic/2a384fbc55ceb0de1b255853bfb8b3db.jpg",
        "creat_time": "",
        "creator_nickname": "",
        "play_count": "105607",
        "playlist_name": "【环球之音】不能错过的经典对唱情歌",
        "songs": [
            {
                "singer_id": "",
                "singer_name": "张学友,陈慧娴",
                "song_id": "7325804",
                "song_name": "接近",
                "source_url": "http://music.taihe.com/song/7325804"
            },
            {
                "singer_id": "",
                "singer_name": "张国荣,许冠杰",
                "song_id": "7316755",
                "song_name": "沉默是金",
                "source_url": "http://music.taihe.com/song/7316755"
            },
            {
                "singer_id": "",
                "singer_name": "王菲,张学友",
                "song_id": "7343680",
                "song_name": "非常夏日",
                "source_url": "http://music.taihe.com/song/7343680"
            },
            {
                "singer_id": "",
                "singer_name": "关淑怡,谭咏麟",
                "song_id": "7312629",
                "song_name": "明天你是否依然爱我",
                "source_url": "http://music.taihe.com/song/7312629"
            },
            {
                "singer_id": "",
                "singer_name": "陈慧琳,冯德伦",
                "song_id": "7326854",
                "song_name": "北极雪",
                "source_url": "http://music.taihe.com/song/7326854"
            },
            {
                "singer_id": "",
                "singer_name": "高慧君",
                "song_id": "7350102",
                "song_name": "你最珍贵",
                "source_url": "http://music.taihe.com/song/7350102"
            },
            {
                "singer_id": "",
                "singer_name": "熊天平,许茹芸",
                "song_id": "1576268",
                "song_name": "爱情电影",
                "source_url": "http://music.taihe.com/song/1576268"
            },
            {
                "singer_id": "",
                "singer_name": "关淑怡,张学友",
                "song_id": "7340288",
                "song_name": "问",
                "source_url": "http://music.taihe.com/song/7340288"
            },
            {
                "singer_id": "",
                "singer_name": "杨千嬅,林海峰",
                "song_id": "7354099",
                "song_name": "饱满爱",
                "source_url": "http://music.taihe.com/song/7354099"
            },
            {
                "singer_id": "",
                "singer_name": "郑中基,陈慧琳",
                "song_id": "7327494",
                "song_name": "都是你的错",
                "source_url": "http://music.taihe.com/song/7327494"
            },
            {
                "singer_id": "",
                "singer_name": "许志安,陈慧珊",
                "song_id": "7319588",
                "song_name": "苦口良药",
                "source_url": "http://music.taihe.com/song/7319588"
            },
            {
                "singer_id": "",
                "singer_name": "谭咏麟,邝美云",
                "song_id": "7312066",
                "song_name": "分手之后",
                "source_url": "http://music.taihe.com/song/7312066"
            },
            {
                "singer_id": "",
                "singer_name": "陈慧娴,张学友",
                "song_id": "284637833",
                "song_name": "一对寂寞的心",
                "source_url": "http://music.taihe.com/song/284637833"
            },
            {
                "singer_id": "",
                "singer_name": "周慧敏",
                "song_id": "7322100",
                "song_name": "爱到最后",
                "source_url": "http://music.taihe.com/song/7322100"
            },
            {
                "singer_id": "",
                "singer_name": "Michael Au,张学友",
                "song_id": "7333211",
                "song_name": "烟花句",
                "source_url": "http://music.taihe.com/song/7333211"
            },
            {
                "singer_id": "",
                "singer_name": "关淑怡,谭咏麟",
                "song_id": "7341932",
                "song_name": "唱一首好歌",
                "source_url": "http://music.taihe.com/song/7341932"
            },
            {
                "singer_id": "",
                "singer_name": "李克勤,陈苑淇",
                "song_id": "120994492",
                "song_name": "合久必婚",
                "source_url": "http://music.taihe.com/song/120994492"
            },
            {
                "singer_id": "",
                "singer_name": "张学友,许志安,郑中基",
                "song_id": "7341020",
                "song_name": "甲乙丙丁",
                "source_url": "http://music.taihe.com/song/7341020"
            },
            {
                "singer_id": "",
                "singer_name": "王馨平,高明骏",
                "song_id": "7317558",
                "song_name": "今生注定",
                "source_url": "http://music.taihe.com/song/7317558"
            },
            {
                "singer_id": "",
                "singer_name": "郑中基,张学友",
                "song_id": "7341599",
                "song_name": "左右为难",
                "source_url": "http://music.taihe.com/song/7341599"
            },
            {
                "singer_id": "",
                "singer_name": "张淑玲,蔡国权",
                "song_id": "7317045",
                "song_name": "爱到浓时",
                "source_url": "http://music.taihe.com/song/7317045"
            },
            {
                "singer_id": "",
                "singer_name": "迪克牛仔,李翊君",
                "song_id": "7314542",
                "song_name": "回头是岸",
                "source_url": "http://music.taihe.com/song/7314542"
            },
            {
                "singer_id": "",
                "singer_name": "许茹芸",
                "song_id": "14431076",
                "song_name": "你的眼睛",
                "source_url": "http://music.taihe.com/song/14431076"
            },
            {
                "singer_id": "",
                "singer_name": "张学友,何惠如",
                "song_id": "939718",
                "song_name": "似曾相识",
                "source_url": "http://music.taihe.com/song/939718"
            },
            {
                "singer_id": "",
                "singer_name": "熊天平,阮丹青",
                "song_id": "7320154",
                "song_name": "我只能相信你",
                "source_url": "http://music.taihe.com/song/7320154"
            },
            {
                "singer_id": "",
                "singer_name": "童安格,金素梅",
                "song_id": "7322935",
                "song_name": "钻与石",
                "source_url": "http://music.taihe.com/song/7322935"
            },
            {
                "singer_id": "",
                "singer_name": "成龙,谭咏麟",
                "song_id": "7343542",
                "song_name": "我做得到",
                "source_url": "http://music.taihe.com/song/7343542"
            },
            {
                "singer_id": "",
                "singer_name": "草蜢,关淑怡,刘小慧,汤宝如",
                "song_id": "7316691",
                "song_name": "热力节拍 Wou Bom Ba",
                "source_url": "http://music.taihe.com/song/7316691"
            },
            {
                "singer_id": "",
                "singer_name": "陈晓东,陈慧琳,陈晓东,邱颖欣",
                "song_id": "258498947",
                "song_name": "打开天空",
                "source_url": "http://music.taihe.com/song/258498947"
            },
            {
                "singer_id": "",
                "singer_name": "区瑞强,卢冠廷,关正杰",
                "song_id": "2604874",
                "song_name": "蚌的启示",
                "source_url": "http://music.taihe.com/song/2604874"
            },
            {
                "singer_id": "",
                "singer_name": "蔡国权,谭咏麟",
                "song_id": "7345374",
                "song_name": "风中劲草",
                "source_url": "http://music.taihe.com/song/7345374"
            },
            {
                "singer_id": "",
                "singer_name": "张学友,谭咏麟,郑中基",
                "song_id": "7344336",
                "song_name": "感激",
                "source_url": "http://music.taihe.com/song/7344336"
            },
            {
                "singer_id": "",
                "singer_name": "陈秀雯,麦子杰",
                "song_id": "7345384",
                "song_name": "缘订今生",
                "source_url": "http://music.taihe.com/song/7345384"
            },
            {
                "singer_id": "",
                "singer_name": "刘小慧,黄凯芹",
                "song_id": "7316113",
                "song_name": "望星星",
                "source_url": "http://music.taihe.com/song/7316113"
            },
            {
                "singer_id": "",
                "singer_name": "苏永康,陈慧珊",
                "song_id": "7319583",
                "song_name": "暖流",
                "source_url": "http://music.taihe.com/song/7319583"
            },
            {
                "singer_id": "",
                "singer_name": "许志安,吴国敬",
                "song_id": "7317433",
                "song_name": "真心真意",
                "source_url": "http://music.taihe.com/song/7317433"
            }
        ],
        "source": "baidu",
        "subscribed_count": "677",
        "tags": "粤语,华语,怀旧",
        "total_songs": 36
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http: //localhost:8080/api/v1/song/tencent/000bdxcy3ta8Q3</summary>
```json
{
    "result": {
        "cover_url": "",
        "publish_time": "2017-03-02",
        "singer_id": "",
        "singer_name": "",
        "song_name": "飘向北方",
        "source": "tencent",
        "source_url": "https://y.qq.com/n/yqq/song/000bdxcy3ta8Q3.html"
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http://localhost:8080/api/v1/lyric/netease/418603133</summary>
```
{
    "result": {
        "lyric": "[by:鱼见]\n[00:26.59]Hello from the other side\n[00:56.26]Hello, it's me\n[01:01.76]I was wondering if after all these years you'd like to meet\n[01:07.85]To go over everything\n[01:13.55]They say that time's supposed to heal ya but I ain't done much healing\n[01:20.72]Hello, can you hear me\n[01:25.90]I'm in California dreaming about who we used to be\n[01:31.90]When we were younger and free\n[01:38.36]I've forgotten how it felt before the world fell at our feet\n[01:44.17]There's such a difference between us\n[01:50.51]And a million miles\n[01:57.06]Hello from the other side\n[02:02.86]I must've called a thousand times to tell you\n[02:09.75]I'm sorry, for everything that I've done\n[02:14.26]But when I call you never seem to be home\n[02:21.45]Hello from the other side\n[02:27.23]At least I can say that I've tried to tell you\n[02:33.84]I'm sorry, for breaking your heart\n[02:38.47]But it don't matter, it clearly doesn't tear you apart anymore\n[03:05.83]Hello from the other side\n[03:09.41]At least I can say that I've tried to tell you\n[03:13.58]I'm sorry, for breaking your heart\n[03:16.57]But it don't matter, it clearly doesn't tear you apart anymore\n[03:42.03]Hello from the other side\n[04:14.41]Hello from the other side\n[04:18.07]At least I can say that I've tried to tell you\n[04:22.15]I'm sorry, for breaking your heart\n[04:25.17]But it don't matter, it clearly doesn't tear you apart anymore\n",
        "tlyric": "[by:Armin6]\n[00:26.59]我还是想打给你 即使相隔天边\n[00:56.26]你好吗 是我\n[01:01.76]我犹豫着要不要给你来电 我不确定多年之后的今日你是否还愿意见我\n[01:07.85]是否愿意来闲聊寒暄 细数从前\n[01:13.55]人们都说时间能治愈一切 但似乎这说法不怎么适合我\n[01:20.72]嘿 你在听吗\n[01:25.90]我会梦到从前 美好的加州 美好的我们\n[01:31.90]当时那么年轻 向往自由的我们\n[01:38.36]我都快要忘了 但现实却让一切重现眼前\n[01:44.17]我们之间的差距愈见明显\n[01:50.51]有如天差地别\n[01:57.06]我还是想打给你 即使相隔天边\n[02:02.86]即使打上千遍万遍我也想给你来电\n[02:09.75]对我从前所有的一切 说声抱歉\n[02:14.26]但似乎我每次来电 都是忙音不断 没人接\n[02:21.45]我还是想打给你 即使相隔天边\n[02:27.23]至少能让我不留遗憾 告诉你我的想念\n[02:33.84]我想说我伤了你的心 真的很抱歉\n[02:38.47]但也许值得庆幸的是 不会再有人让你悲痛欲绝\n[03:05.83]我还是想打给你 即使相隔天边\n[03:09.41]至少能让我不留遗憾 告诉你我的想念\n[03:13.58]我想说我伤了你的心 真的很抱歉\n[03:16.57]但也许值得庆幸的是 不会再有人让你悲痛欲绝\n[03:42.03]我还是想打给你 即使相隔天边\n[04:14.41]我还是想打给你 即使相隔天边\n[04:18.07]至少能让我不留遗憾 告诉你我的想念\n[04:22.15]我想说我伤了你的心 真的很抱歉\n[04:25.17]但也许值得庆幸的是 不会再有人让你悲痛欲绝\n"
    },
    "status": "ok"
}
```
</details><br/>

<details>
<summary>http://localhost:8080/api/v1/download/kugou/09E8DE70A24C97B92A29F6A19F3528A2</summary>
```
{
    "result": {
        "url": "http:\/\/fs.ios.kugou.com\/201811241422\/75b9e41a18aac474d88dc1130ce9d7bf\/G066\/M06\/07\/05\/4oYBAFdWacmAKv-MADc2_4Ip1v0854.mp3",
    },
    "status": "ok"
}
```
</details><br/>

@MIT