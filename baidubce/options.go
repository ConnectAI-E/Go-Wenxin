package baidubce

import (
	"context"
	"errors"

	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
)

type Option func(*Client) error

func WithBaseUrl(url string) Option {
	return func(cli *Client) error {
		if len(url) == 0 {
			return errors.New("base url can not empty")
		}
		cli.client.SetBaseURL(url)
		return nil
	}
}

func WithToken(token string) Option {
	return func(cli *Client) error {
		if len(token) == 0 {
			return errors.New("token can not empty")
		}
		cli.accessToken = token
		return nil
	}
}

func WithTokenRequest(in *baidubcev1.TokenRequest) Option {
	return func(cli *Client) error {
		if in.GrantType == "" {
			in.GrantType = "client_credentials"
		}
		res, err := cli.Token(context.Background(), in)
		if err == nil && res != nil && len(res.AccessToken) > 0 {
			cli.accessToken = res.AccessToken
		}
		return err
	}
}
