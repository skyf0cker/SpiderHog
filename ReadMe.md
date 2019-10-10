# SpiderHog 轻量级异步爬虫框架
`author：vophan`

编码不易，Star me！

![logo](https://github.com/skyf0cker/SpiderHog/blob/master/logo.jpeg?raw=true)

## Update!

让大家等了好久，迎来了第二次更新。

Gospider正式更名为**SpiderHog**

这次对代码进行了彻底重构，让使用更加简易，并且加入许多新的内容：

- 实现log系统，日志分割
- 包装proxy，与代理池完美结合，方便插拔
- 加入web api，监控任务量以及最新日志查看
- 增加稳定性与可用性，加入超时机制，自动检测代理存活，自动更换代理
- ..........

## Waiting for implement

这次重构，其实我认为还是有许多地方可以优化：

- Fetcher中可以加入goroutine pool将爬取操作并行化，发挥极致性能，但是由于架构原因还没有想到好的方法，将pool包装到Fetcher中
- 需要更多的web api
- 需要为监控api写一个前端
- .......



## How to use?

首先，如果不使用代理：

1. 根据自己的情况，修改settings.json文件，去掉`proxysrc`接口地址
2. 重写`service.go`中`LinkExtract`和`ContentExtrat`还有`Save`方法

如果使用代理：

1. 首先得配置代理池，我这里使用了proxy_pool的开源实现，结合docker简单方便，如果自己实现代理池，需要自己修改接口。
2. 然后就上面一样了



## Web api

使用了`gin`框架，高效，快速。

- 127.0.0.1:3000/api/v1/log 获取日志内容

- 127.0.01:3000/api/v1/chan 获取任务队列中数量

## Demo

这里，我们举个例子：

我们爬取法邦网的法律文书作为演示：

首先配置代理池：

```shell
docker pull jhao104/proxy_pool
docker pull redis
docker run -p 6379:6379 -d redis redis-server
docker run --env db_type=REDIS --env db_host=172.17.0.2 --env db_port=6379 -p 5010:5010 jhao104/proxy_pool
```

然后修改`settings.json`:

```json
{
  "Fetcher": {
    "header": {
      "User-Agent": ["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/76.0.3809.100 Chrome/76.0.3809.100 Safari/537.36"]
    },
    "cookies":[{
      "Name": "",
      "Value": ""
    }],
    "begin_url": "https://code.fabao365.com/level_4.html",
    "method": "GET",
    "raw_data": {},
    "proxy_source": "http://127.0.0.1:5010"
  },
  "Parser": {},
  "Saver": {}
}

```

然后，重写`service.go`中的方法：

```go
func (p ParserKernel) LinkExtract(response string) (urllist []url.URL) {

        dom, _ := goquery.NewDocumentFromReader(strings.NewReader(response))
        selector := dom.Find("div.bs2l").Find("a[class=ch12],a[class=next]")
        selector.Each(func(i int, selection *goquery.Selection) {
                href, _ := selection.Attr("href")
                href = strings.TrimSpace(href)
                if !strings.Contains(href, host) {
                        href = host + href
                }
                Href, _ := url.Parse(href)
                urllist = append(urllist, *Href)
        })
        return
}

func (p ParserKernel) ContentExtract(response string) (Contents []interface{}) {
        //html, _ := utils.Response2String(response)
        selector := utils.ParseBySelector(response, ".bnr_s7")
        selector.Each(func(i int, selection *goquery.Selection) {
                content := strings.TrimSpace(selection.Text())
                Contents = append(Contents, content)
        })
        return
}

func (s SaverKernel) Save(content interface{}) (error) {

        if str, ok := content.(string); ok {
                strs := utils.ParseByReg(str, "【法规名称】(.*?)\n")
                title := strs[1]
                content := str
                fmt.Printf("Title:%s\n", title)
                fmt.Printf("Content:%s\n", content)
                return nil
        } else {
                return errors.New("Type Assertion Error")
        }
}

```

然后运行

## To be Continue

还是希望有更多的gopher加入我，一起来完善这个项目。

也希望得到大家的star，算是对我的一种支持！

![](https://upload-images.jianshu.io/upload_images/15885453-5cf10fcf339a87fc.png?imageMogr2/auto-orient/strip|imageView2/2/w/1037/format/webp)



## what's the GoSpider?

GoSpider的初衷其实是一个Go语言的入门级项目，大佬说，Golang三大入门项目`爬虫`，`博客`，`电商`。因为我是python转go，所以写一个爬虫真的是很小儿科了，所以，我想既然要做，不如去写一个框架，虽然之前用过`scrapy`这样的框架，但是，从来没有机会去自己想想框架是如何实现的，所以就有了GoSpider。

**但是，这并不代表他只是一个入门学习项目**，虽然是抱着学习的心态去做这个框架，但是，并没有想要放弃这个项目，我会维护这个项目，并不断更新，也希望有更多的朋友，小白，大佬一起加入进来。

## Structure

```
--GoSpider----conf				配置文件解析
	   |--demo				框架使用实例
  	   |--spider				爬虫核心代码
	   |--test				测试代码
           |--userAgent				userAgent文件
	   |--log				日志文件
	   |--utils				工具
```

框架的设计架构借鉴了`PSpider`的架构：

![](https://raw.githubusercontent.com/xianhu/PSpider/master/procedure.png)

## How to use it

目前，如果想要使用它，只需要：

> 1. 在根目录创建`main.go`
> 2. 在`main.go`中重写`parser`的`MiddleLinkRule`以及`LinkExtract`函数，`Saver`的`saveRule`函数
> 3. 根据需要修改`settings.json`配置文件

运行！

## For Learner

如果，你也像我一样，想要通过这个项目，熟悉Go语言，那么，实现它，**你将**：

1. 熟悉Golang的面向对象，接口，所谓的继承与多态，实际上的组合与接口
2. 了解Golang中的gorooutine，channel，sync，select等有关并发与协程的操作
3. 了解Golang中的http的相关操作
4. 熟悉Golang的语法

............

## For Developer

欢迎每一个想要加入的朋友，一起交流，毕竟他还是第一版，所以肯定大大小小的不足，所以，**我需要你们**：

1. 爬虫框架的log模块尚未实现
2. 爬虫框架的proxier模块尚需完善

.........

总之，我很菜，需要大家的帮助

## Contact Me

`qq`:`809866729`

`wx`:

![](https://raw.githubusercontent.com/skyf0cker/NumPy100/master/qrcode.png)
