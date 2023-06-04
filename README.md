
<p align='center'>
    <img src='https://github.com/ConnectAI-E/go-wenxin/assets/50035229/84af1dfe-43c2-456b-85ce-52e787e034cb' alt='' width='1300'/>
</p>


## Baidu Wenxin SDK

> 封装 百度文心一言 API

### 功能及特点

1. 全接口字段注释
2. ChatCompletions 文心一言云服务
3. ChatEbInstant ErnieBot-Lite
4. 无缝对接官方文档：单轮问答、历史记忆问答、流返回

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


### 快速上手 examples:

<details>
<summary>ErnieBot completion</summary>

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
</details>


<details>
<summary>ErnieBot stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ConnectAI-E/go-wenxin/baidubce"
	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"io"
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
		Stream: true,
		Messages: []*ai_customv1.Message{
			{Role: "user",
				Content: "用鲁迅的口气写一封道歉信，给领导说对不起，我不该在会议上睡觉。200字左右"},
		},
	}
	stream, _ := client.ChatCompletionsStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Result)
	}

}

```
</details>


<details>
<summary>ErnieBot history stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ConnectAI-E/go-wenxin/baidubce"
	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"io"
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
		Stream: true,
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: "推荐三本书"},
			{Role: "assistant", Content: "1. 《活着》、《许三观卖血记》和《被掩埋的巨人》\n2. 《小王子》、《浮生六记》和《人类简史》\n3. 《百年孤独》、《查泰莱夫人的情人》和《1984》"},
			{Role: "user",
				Content: "从中选一本适合小朋友看的"},
		},
	}
	stream, _ := client.ChatCompletionsStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Result) //建议选择《小王子》一书，这本书的故事简单易懂，讲述了一个外星王子来到地球上的经历，以及他与一朵玫瑰、一只狐狸之间的奇遇。这本书的插图也很精美，有助于小朋友理解故事情节。此外，《小王子》一书中的寓言和哲理也对小朋友有一定启示作用，帮助他们理解人性
	}

}


```
</details>
