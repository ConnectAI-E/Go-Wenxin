package baidubce

import (
	"context"
	"testing"

	ernievilgv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ernievilg/v1"
	"github.com/stretchr/testify/assert"
)

func TestTxt2ImgV2(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	res, err := cli.Txt2ImgV2(context.Background(), &ernievilgv1.Txt2ImgV2Request{
		Prompt: "一艘友谊的小船",
		Width:  800,
		Height: 600,
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetImgV2(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	res, err := cli.GetImgV2(context.Background(), &ernievilgv1.GetImgV2Request{
		TaskId: "",
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
