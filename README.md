# 说明

golang 爬取小说资源

## 小说来源

* [笔趣阁](http://www.biquge5200.com/)

## 启动项目

```
go run main.go
```

## 如何使用

以小说《圣墟》为例。

* 搜索小说资源`http://localhost:9000/novel/search/圣墟`

```
{
    "code": 200,
    "message": "OK",
    "data": [
        {
            "novelName": "圣墟",
            "novelUrl": "http://www.biquge5200.com/52_52542/",
            "lastChapterName": " 第六百四十一章 按捺冲动",
            "author": "辰东",
            "wordCount": "4K",
            "updateTime": "2017-09-17",
            "status": "连载",
            "source": 1
        },
        {
            "novelName": "太古圣墟",
            "novelUrl": "http://www.biquge5200.com/77_77328/",
            "lastChapterName": " 新书",
            "author": "春上三更",
            "wordCount": "12K",
            "updateTime": "2017-07-11",
            "status": "连载中",
            "source": 1
        },
        {
            "novelName": "魔法圣墟",
            "novelUrl": "http://www.biquge5200.com/77_77835/",
            "lastChapterName": " 第七章美女有何贵干",
            "author": "凌逐",
            "wordCount": "22K",
            "updateTime": "2017-07-25",
            "status": "连载中",
            "source": 1
        },
        {
            "novelName": "圣墟封神装逼系统",
            "novelUrl": "http://www.biquge5200.com/53_53409/",
            "lastChapterName": " 028 吃下未知的吞噬流法则道果",
            "author": "风十三郎A",
            "wordCount": "279K",
            "updateTime": "2017-08-15",
            "status": "连载中",
            "source": 1
        }
    ]
}
```

* 根据找到的`novelUrl`获取小说`http://localhost:9000/novel?url=http://www.biquge5200.com/52_52542/`

```
{
    "code": 200,
    "message": "OK",
    "data": {
        "name": "圣墟",
        "author": "辰东",
        "lastUpdateTime": "2017-09-17",
        "description": "在破败中崛起，在寂灭中复苏。沧海成尘，雷电枯竭，那一缕幽雾又一次临近大地，世间的枷锁被打开了，一个全新的世界就此揭开神秘的一角……",
        "chapters": [
            {
                "name": "第一章 沙漠中的彼岸花",
                "url": "http://www.biquge5200.com/52_52542/20380548.html",
                "content": ""
            },
            {
                "name": "第二章 后文明时代",
                "url": "http://www.biquge5200.com/52_52542/20380549.html",
                "content": ""
            },
            ...
            // 省略掉其余章节
        }
    }
```

* 根据章节`url`获取该章节内容`http://localhost:9000/novel/chapter?url=http://www.biquge5200.com/52_52542/20380549.html`

```
{
    "code": 200,
    "message": "OK",
    "data": {
        "name": " 第二章 后文明时代",
        "url": "",
        // 省略掉其余内容
        "content": "牛羊等牲畜疑似受惊，差点闯出栅栏，牧民阻拦，大声喝斥着，平日间几头很凶的藏獒此时低伏在地，嘶吼着，很不安。\n\n"
    }
}
```

## 不足

目前只支持对笔趣阁网站进行实时的页面抓取，并没有使用数据库存储信息。也没有web页面进行在线阅读小说。

## 后期完善

* 提供web界面进行在线浏览小说资源。
* 爬取其他小说网站的资源。
* 建立自己的数据库存储爬取得小说信息。（暂不考虑）
* 建立小说网站。（暂不考虑）