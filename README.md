## Baidu Wenxin SDK

> 封装 百度文心一言 API

### 功能及特点

1. 全接口字段注释
2. ChatCompletions 文心一言云服务
3. ChatEbInstant ErnieBot-Lite
4. 更多功能同步中。。。

### 使用方法

1. 访问 https://console.bce.baidu.com/ai/#/ai/wenxinworkshop/app/list 并创建应用。
2. 取 API Key 作为 ClientId，Secret Key 作为 ClientSecret。
3. 使用 New 方法并根据传参说明生成 client。

### 示例

> 更多使用方法参考 test 文件

```go
import "github.com/ConnectAI-E/go-wenxin/baidubce"

cli, err := New(&baidubcev1.TokenRequest{
  GrantType:    "client_credentials",
  ClientId:     os.Getenv("TEST_BAIDU_CLIENT_ID"),
  ClientSecret: os.Getenv("TEST_BAIDU_CLIENT_SECRET"),
})
if err != nil {
  panic(err)
}

res, err := cli.ChatEbInstant(context.Background(), &ai_customv1.ChatEbInstantRequest{
  Messages: []*ai_customv1.Message{
    {Role: "user", Content: "hi"},
  },
})
if err != nil {
  panic(err)
}
log.Println(res.Result) // output: 嗨！有什么我可以帮助你的吗？
```
