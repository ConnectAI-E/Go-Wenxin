package baidubce

import (
	"context"
	stdErrors "errors"
	"io"
	"testing"

	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	"github.com/ConnectAI-E/go-wenxin/pkg/errors"
	"github.com/stretchr/testify/assert"
)

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
		if stdErrors.Is(err, io.EOF) || stdErrors.Is(err, errors.ErrTooManyEmptyStreamMessages) {
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
