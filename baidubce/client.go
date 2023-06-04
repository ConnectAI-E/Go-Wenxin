package baidubce

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"github.com/imroc/req/v3"
	"github.com/moul/http2curl"
)

const baidubceBaseUrl = "https://aip.baidubce.com"

type Client struct {
	client      *req.Client
	accessToken string
}

func New(opts ...Option) (*Client, error) {
	cli := &Client{
		client: req.C().
			SetBaseURL(baidubceBaseUrl),
	}

	cli.client.
		OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			if len(cli.accessToken) > 0 {
				req.SetQueryParam("access_token", cli.accessToken)
			}
			return nil
		}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if os.Getenv("APP_ENV") == "debug" {
				curl, err := http2curl.GetCurlCommand(resp.Request.RawRequest)
				if err != nil {
					return err
				}
				log.Println(curl)
			}
			return nil
		})

	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return nil, err
		}
	}

	return cli, nil
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

	return res, err
}
