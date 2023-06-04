package baidubce

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"

	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"github.com/ConnectAI-E/go-wenxin/internal"
	"github.com/stretchr/testify/assert"
)

func getTestClient() (*Client, error) {
	return New(&baidubcev1.TokenRequest{
		GrantType:    "client_credentials",
		ClientId:     os.Getenv("TEST_BAIDU_CLIENT_ID"),
		ClientSecret: os.Getenv("TEST_BAIDU_CLIENT_SECRET"),
	})
}

func TestNew(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
}

func TestChatCompletions(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	res, err := cli.ChatCompletions(context.Background(), &ai_customv1.ChatCompletionsRequest{
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: "hi"},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestChatCompletionsStream(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	res, err := cli.ChatCompletionsStream(context.Background(), &ai_customv1.ChatCompletionsRequest{
		Stream: true,
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: "hi"},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	hasEnd := false
	for {
		response, err := res.Recv()
		if errors.Is(err, io.EOF) || errors.Is(err, internal.ErrTooManyEmptyStreamMessages) {
			break
		}
		assert.Nil(t, err)
		assert.NotNil(t, response)
		if response.IsEnd {
			hasEnd = true
			break
		}
	}
	assert.True(t, hasEnd)
}

func TestChatEbInstant(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)

	res, err := cli.ChatEbInstant(context.Background(), &ai_customv1.ChatEbInstantRequest{
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: "hi"},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
