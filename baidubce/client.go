package baidubce

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"github.com/ConnectAI-E/go-wenxin/internal"
	"github.com/imroc/req/v3"
	"github.com/moul/http2curl"
	"google.golang.org/grpc"
)

const baidubceBaseUrl = "https://aip.baidubce.com"

type Client struct {
	client      *req.Client
	accessToken string
}

var _ ai_customv1.WenxinworkshopServiceClient = new(Client)

func New(token *baidubcev1.TokenRequest) (*Client, error) {
	cli := &Client{
		client: req.C().
			SetBaseURL(baidubceBaseUrl),
	}

	cli.client.OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		if os.Getenv("APP_ENV") == "debug" {
			curl, err := http2curl.GetCurlCommand(resp.Request.RawRequest)
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
func (cli *Client) Token(ctx context.Context, in *baidubcev1.TokenRequest) (*baidubcev1.TokenResponse, error) {
	res := new(baidubcev1.TokenResponse)

	resp, err := cli.client.R().
		SetQueryParam("grant_type", in.GrantType).
		SetQueryParam("client_id", in.ClientId).
		SetQueryParam("client_secret", in.ClientSecret).
		SetSuccessResult(res).
		Post("/oauth/2.0/token")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	if err == nil && res != nil && len(res.AccessToken) > 0 {
		cli.accessToken = res.AccessToken
	}

	return res, err
}

// ChatCompletions 文心一言云服务
func (cli Client) ChatCompletions(
	ctx context.Context, request *ai_customv1.ChatCompletionsRequest, opts ...grpc.CallOption,
) (*ai_customv1.ChatResponse, error) {
	res := new(ai_customv1.ChatResponse)

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(request).
		SetSuccessResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return res, err
}

// ChatCompletions 文心一言云服务 流
func (cli Client) ChatCompletionsStream(
	ctx context.Context, request *ai_customv1.ChatCompletionsRequest, opts ...grpc.CallOption,
) (ai_customv1.WenxinworkshopService_ChatCompletionsStreamClient, error) {
	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(request).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return internal.NewStreamReader[*ai_customv1.ChatResponse](resp.Body), err
}

// ChatEbInstant ErnieBot-Lite
func (cli Client) ChatEbInstant(
	ctx context.Context, request *ai_customv1.ChatEbInstantRequest, opts ...grpc.CallOption,
) (*ai_customv1.ChatResponse, error) {
	res := new(ai_customv1.ChatResponse)

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(request).
		SetSuccessResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant")

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return res, err
}
