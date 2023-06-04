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

```go
package main

import (
	"context"
	"fmt"
	"github.com/ConnectAI-E/go-wenxin/baidubce"
	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
)

//init client

func main() {
	ctx := context.Background()
	var opts []baidubce.Option
	opts = append(opts, baidubce.WithTokenRequest(&baidubcev1.TokenRequest{
		GrantType:    "client_credentials",
		ClientId:     "YOUR BAIDU_API Key",
		ClientSecret: "YOUR BAIDU_SECRET Key",
	}))
	client, _ := baidubce.New(opts...)

	//chat
	req := &ai_customv1.ChatCompletionsRequest{
		User: "feishu-user",
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: "嗨"},
		},
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Result) // output: 嗨！有什么我可以帮助你的吗？
}

```
