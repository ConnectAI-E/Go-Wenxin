package baidubce

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"github.com/go-resty/resty/v2"
	"github.com/moul/http2curl"
)

const baidubceBaseUrl = "https://aip.baidubce.com"

type Client struct {
	client      *resty.Client
	accessToken string

	baidubcev1.UnimplementedOauthServiceServer
	ai_customv1.UnimplementedWenxinworkshopServiceServer
}

func New(token *baidubcev1.TokenRequest) (*Client, error) {
	cli := &Client{
		client: resty.New().
			SetBaseURL(baidubceBaseUrl),
	}

	cli.client.SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
		if os.Getenv("APP_ENV") == "debug" {
			curl, err := http2curl.GetCurlCommand(r)
			if err != nil {
				return err
			}
			log.Println(curl)
		}
		return nil
	})

	_, err := cli.Token(context.Background(), token)

	return cli, err
}

// Token 获取AccessToken
func (cli *Client) Token(ctx context.Context, request *baidubcev1.TokenRequest) (*baidubcev1.TokenResponse, error) {
	res := new(baidubcev1.TokenResponse)

	resp, err := cli.client.R().
		SetQueryParam("grant_type", request.GrantType).
		SetQueryParam("client_id", request.ClientId).
		SetQueryParam("client_secret", request.ClientSecret).
		SetResult(res).
		Post("/oauth/2.0/token")

	if resp.StatusCode() != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	if err == nil && res != nil && len(res.AccessToken) > 0 {
		cli.accessToken = res.AccessToken
	}

	return res, err
}

// ChatCompletions 文心一言云服务
func (cli Client) ChatCompletions(ctx context.Context, request *ai_customv1.ChatCompletionsRequest) (*ai_customv1.ChatResponse, error) {
	res := new(ai_customv1.ChatResponse)

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(request).
		SetResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions")

	if resp.StatusCode() != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	return res, err
}

// ChatEbInstant ErnieBot-Lite
func (cli Client) ChatEbInstant(ctx context.Context, request *ai_customv1.ChatEbInstantRequest) (*ai_customv1.ChatResponse, error) {
	res := new(ai_customv1.ChatResponse)

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(request).
		SetResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant")

	if resp.StatusCode() != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	return res, err
}
