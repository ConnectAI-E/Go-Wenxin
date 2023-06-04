package baidubce

import (
	"context"
	"errors"
	"io"

	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	commonv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/common/v1"
	"github.com/ConnectAI-E/go-wenxin/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ ai_customv1.WenxinworkshopServiceClient = new(Client)

// ChatCompletions 文心一言云服务
func (cli Client) ChatCompletions(
	ctx context.Context, in *ai_customv1.ChatCompletionsRequest, opts ...grpc.CallOption,
) (*ai_customv1.ChatResponse, error) {
	res := new(struct {
		commonv1.Error
		ai_customv1.ChatResponse
	})

	in.Stream = false
	resp, err := cli.client.R().
		SetBody(in).
		SetSuccessResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions")
	if err != nil {
		return nil, err
	}
	if res.ErrorCode > 0 {
		return nil, errors.New(res.GetErrorMsg())
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return &res.ChatResponse, err
}

// ChatCompletions 文心一言云服务 流
func (cli Client) ChatCompletionsStream(
	ctx context.Context, in *ai_customv1.ChatCompletionsRequest, opts ...grpc.CallOption,
) (ai_customv1.WenxinworkshopService_ChatCompletionsStreamClient, error) {
	in.Stream = true
	resp, err := cli.client.R().
		DisableAutoReadResponse().
		SetBody(in).
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
	ctx context.Context, in *ai_customv1.ChatEbInstantRequest, opts ...grpc.CallOption,
) (*ai_customv1.ChatResponse, error) {
	res := new(struct {
		commonv1.Error
		ai_customv1.ChatResponse
	})

	resp, err := cli.client.R().
		SetBody(in).
		SetSuccessResult(res).
		Post("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant")
	if err != nil {
		return nil, err
	}
	if res.ErrorCode > 0 {
		return nil, errors.New(res.GetErrorMsg())
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	return &res.ChatResponse, err
}

// Balance
func (cli Client) Balance(
	ctx context.Context, in *ai_customv1.BalanceRequest, opts ...grpc.CallOption,
) (*ai_customv1.BalanceResponse, error) {
	return &ai_customv1.BalanceResponse{}, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
