package baidubce

import (
	"context"
	"errors"
	"io"

	"google.golang.org/grpc"

	commonv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/common/v1"
	ernievilgv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ernievilg/v1"
)

var _ ernievilgv1.ErnievilgServiceClient = new(Client)

func (cli Client) Txt2ImgV2(ctx context.Context, in *ernievilgv1.Txt2ImgV2Request, opts ...grpc.CallOption) (*ernievilgv1.Txt2ImgV2Response, error) {
	res := new(struct {
		commonv1.Error
		ernievilgv1.Txt2ImgV2Response
	})

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(in).
		SetSuccessResult(res).
		Post("/rpc/2.0/ernievilg/v1/txt2imgv2")
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

	return &res.Txt2ImgV2Response, nil
}

func (cli Client) GetImgV2(ctx context.Context, in *ernievilgv1.GetImgV2Request, opts ...grpc.CallOption) (*ernievilgv1.GetImgV2Response, error) {
	res := new(struct {
		commonv1.Error
		ernievilgv1.GetImgV2Response
	})

	resp, err := cli.client.R().
		SetQueryParam("access_token", cli.accessToken).
		SetBody(in).
		SetSuccessResult(res).
		Post("/rpc/2.0/ernievilg/v1/getImgv2")
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

	return &res.GetImgV2Response, err
}
